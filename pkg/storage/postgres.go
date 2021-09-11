package storage

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/awile/datamkr/pkg/config"
	_ "github.com/lib/pq"
)

type postgresStorageService struct {
	ConnectionString string
	Columns          []string
	Table            string

	db  *sql.DB
	ctx context.Context
}

func newPostgresStorageWriter(config *config.DatamkrConfig, opts WriterOptions) StorageServiceInterface {
	var storageService postgresStorageService

	storageService.ConnectionString = opts.Id
	storageService.Table = opts.SecondaryId

	return &storageService
}

func (pss *postgresStorageService) Init() error {
	db, err := sql.Open("postgres", pss.ConnectionString)
	if err != nil {
		return err
	}
	pss.db = db

	columns_query := fmt.Sprintf("SELECT column_name FROM information_schema.columns WHERE table_name = '%s'", pss.Table)
	res, err := db.Query(columns_query)
	if err != nil {
		return err
	}

	var column string
	var columns []string

	for res.Next() {
		err = res.Scan(&column)
		if err != nil {
			return err
		}
		columns = append(columns, column)
	}
	pss.Columns = columns

	return nil
}

func (pss *postgresStorageService) Write(data map[string]interface{}) error {
	values := make([]string, len(data))
	for i, column := range pss.Columns {
		values[i] = fmt.Sprintf("'%s'", data[column].(string))
	}
	insert_query := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		pss.Table,
		strings.Join(pss.Columns, ","),
		strings.Join(values, ","),
	)

	_, err := pss.db.Exec(insert_query)
	if err != nil {
		return err
	}
	return nil
}

func (pss *postgresStorageService) WriteAll(data []map[string]interface{}) error {
	tx, err := pss.db.BeginTx(pss.ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		return err
	}

	column_commas := strings.Join(pss.Columns, ",")
	for _, row := range data {

		values := make([]string, len(data))
		for i, column := range pss.Columns {
			values[i] = row[column].(string)
		}

		insert_query := fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s)",
			pss.Table,
			column_commas,
			strings.Join(values, ","),
		)
		_, err = tx.Exec(insert_query)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (pss *postgresStorageService) Close() error {
	return pss.db.Close()
}
