package storage

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/awile/datamkr/pkg/config"
)

type csvStorageService struct {
	FilePath   string
	HeaderKeys []string
	Writer     *csv.Writer

	fileHandler *os.File

	fileName string
	config   *config.DatamkrConfig
}

func newCsvStorageWriter(config *config.DatamkrConfig, opts WriterOptions) StorageServiceInterface {
	var storageService csvStorageService

	storageService.config = config
	storageService.fileName = opts.Id

	if opts.FieldKeys != nil {
		storageService.HeaderKeys = opts.FieldKeys
	} else {
		headerKeys := make([]string, len(opts.DatasetDefinition.Fields))
		var i int = 0
		for fieldKey := range opts.DatasetDefinition.Fields {
			headerKeys[i] = fieldKey
			i++
		}
		storageService.HeaderKeys = headerKeys
	}

	return &storageService
}

func (css *csvStorageService) Init() error {
	if css.fileName == "" {
		return fmt.Errorf("Must provide csv storage service with a FileName\n")
	}
	css.FilePath = fmt.Sprintf("./%s", css.fileName)

	fileWriter, err := css.getWriter()
	if err != nil {
		return err
	}
	css.Writer = csv.NewWriter(fileWriter)

	err = css.Writer.Write(css.HeaderKeys)
	if err != nil {
		return err
	}

	return nil
}

func (css *csvStorageService) Write(data map[string]interface{}) error {
	if css.Writer == nil {
		return fmt.Errorf("Must init csv writer first: csvStorageService.Init()\n")
	}
	record := make([]string, len(css.HeaderKeys))
	for i, headerKey := range css.HeaderKeys {
		value := data[headerKey]
		record[i] = value.(string)
	}
	return css.Writer.Write(record)
}

func (css *csvStorageService) WriteAll(data []map[string]interface{}) error {
	if css.Writer == nil {
		return fmt.Errorf("Must init csv writer first: csvStorageService.Init()\n")
	}
	records := make([][]string, len(data))
	for i, row := range data {
		record := make([]string, len(css.HeaderKeys))
		for j, headerKey := range css.HeaderKeys {
			value := row[headerKey]
			record[j] = value.(string)
		}
		records[i] = record
	}
	return css.Writer.WriteAll(records)
}

func (css *csvStorageService) Close() error {
	if css.Writer == nil {
		return fmt.Errorf("No csv writer found")
	}
	css.Writer.Flush()

	if css.fileHandler == nil {
		return fmt.Errorf("No file found")
	}
	return css.fileHandler.Close()
}

func (css *csvStorageService) getWriter() (io.Writer, error) {
	f, err := os.OpenFile(css.FilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	css.fileHandler = f
	return f, nil
}
