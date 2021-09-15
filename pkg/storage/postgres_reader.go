package storage

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
)

type postgresStorageServiceReader struct {
	ConnectionString string
	Table            string

	db *sql.DB
}

func newPostgresStorageReader(config *config.DatamkrConfig, opt ReaderOptions) StorageServiceReaderInterface {
	var storageService postgresStorageServiceReader

	storageService.ConnectionString = opt.Id
	storageService.Table = opt.SecondaryId

	return &storageService
}

func (pss *postgresStorageServiceReader) Init() error {
	db, err := sql.Open("postgres", pss.ConnectionString)
	if err != nil {
		return err
	}
	pss.db = db

	return nil
}

func (pss *postgresStorageServiceReader) GetDatasetDefinition() (dataset.DatasetDefinition, error) {
	var datasetDefinition dataset.DatasetDefinition
	datasetDefinition.Fields = make(map[string]dataset.DatasetDefinitionField)
	datasetDefinition.Table = pss.Table

	column_stmt := fmt.Sprintf("SELECT column_name, data_type, character_maximum_length FROM information_schema.columns where table_name = '%s'", pss.Table)
	res, err := pss.db.Query(column_stmt)
	if err != nil {
		return datasetDefinition, err
	}

	var column_name sql.NullString
	var column_type sql.NullString
	var char_limit sql.NullInt32
	for res.Next() {
		err = res.Scan(&column_name, &column_type, &char_limit)
		if err != nil {
			return datasetDefinition, err
		}
		if !(column_name.Valid && column_type.Valid) {
			continue
		}

		fieldDefinition := pss.getColumnType(column_name.String, column_type.String)
		if fieldDefinition.Type == "" {
			continue
		}
		datasetDefinition.Fields[column_name.String] = fieldDefinition
	}

	return datasetDefinition, nil
}

func (pss *postgresStorageServiceReader) Close() error {
	return pss.db.Close()
}

func (pss *postgresStorageServiceReader) getColumnType(column_name string, column_type string) dataset.DatasetDefinitionField {
	var d dataset.DatasetDefinitionField
	if column_type == "character varying" {
		if column_name == "email" {
			d.Type = "email"
		} else if column_name == "name" {
			d.Type = "name"
		} else {
			d.Type = "string"
		}
	} else if column_type == "uuid" {
		d.Type = "uuid"
	} else if column_type == "boolean" {
		d.Type = "boolean"
	} else if column_type == "bigint" {
		d.Type = "int"
	} else if strings.Contains(column_type, "timestamp") {
		d.Type = "datetime"
	}

	return d
}
