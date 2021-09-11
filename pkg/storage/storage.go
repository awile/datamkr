package storage

import (
	"github.com/awile/datamkr/pkg/config"
)

type StorageClientInterface interface {
	GetStorageWriterService(storageType string, opts WriterOptions) StorageServiceWriterInterface
}

type StorageServiceWriterInterface interface {
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

func (sc *storageClient) GetStorageWriterService(storageType string, opts WriterOptions) StorageServiceWriterInterface {
	switch storageType {
	case "csv":
		return newCsvStorageWriter(sc.config, opts)
	case "postgres":
		return newPostgresStorageWriter(sc.config, opts)
	default:
		return nil
	}
}
