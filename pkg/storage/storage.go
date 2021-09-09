package storage

import (
	"github.com/awile/datamkr/pkg/config"
)

type StorageClientInterface interface {
	GetStorageService(storageType string, args interface{}) StorageServiceInterface
}

type StorageServiceInterface interface {
	Init() error
	Write(data map[string]interface{}) error
	WriteAll(data []map[string]interface{}) error
	Close() error
}

type storageClient struct {
	config *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) StorageClientInterface {
	var sc storageClient

	sc.config = config

	return &sc
}

func (sc *storageClient) GetStorageService(storageType string, args interface{}) StorageServiceInterface {
	switch storageType {
	case "csv":
		var csvStorageArgs = args.(CsvStorageArgs)
		return newCsvStorageWriter(sc.config, csvStorageArgs)
	case "local":
		var csvStorageArgs = args.(CsvStorageArgs)
		return newCsvStorageWriter(sc.config, csvStorageArgs)
	default:
		return nil
	}
}
