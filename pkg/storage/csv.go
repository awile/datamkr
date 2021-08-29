package storage

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/awile/datamkr/pkg/config"
)

type csvStorageService struct {
	filePath string
	writer   *csv.Writer

	fileHandler *os.File

	config *config.DatamkrConfig
}

type CsvStorageArgs struct {
	FileName string
	IsWriter bool
}

func newCsvStorageWithConfig(config *config.DatamkrConfig) StorageServiceInterface {
	var storageService csvStorageService

	storageService.config = config

	return &storageService
}

func (css *csvStorageService) Init(args interface{}) error {
	var csvArgs StorageArgs = args.(StorageArgs)

	if csvArgs.FileName == "" {
		return fmt.Errorf("Must provide csv storage service with a FileName\n")
	}
	css.filePath = fmt.Sprintf("./%s", csvArgs.FileName)

	if csvArgs.IsWriter {
		fileWriter, err := css.getWriter()
		if err != nil {
			return err
		}
		css.writer = csv.NewWriter(fileWriter)
	}

	return nil
}

func (css *csvStorageService) Write(data interface{}) error {
	if css.writer == nil {
		return fmt.Errorf("Must init csv writer first: csvStorageService.Init()\n")
	}
	var record []string = data.([]string)
	return css.writer.Write(record)
}

func (css *csvStorageService) WriteAll(data interface{}) error {
	if css.writer == nil {
		return fmt.Errorf("Must init csv writer first: csvStorageService.Init()\n")
	}
	var record [][]string = data.([][]string)
	return css.writer.WriteAll(record)
}

func (css *csvStorageService) Close() error {
	if css.writer == nil {
		return fmt.Errorf("No csv writer found")
	}
	css.writer.Flush()

	if css.fileHandler == nil {
		return fmt.Errorf("No file found")
	}
	return css.fileHandler.Close()
}

func (css *csvStorageService) getWriter() (io.Writer, error) {
	f, err := os.OpenFile(css.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	css.fileHandler = f
	return f, nil
}
