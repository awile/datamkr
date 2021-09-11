package storage

import (
	"github.com/awile/datamkr/pkg/config"
)

type StorageClientInterface interface {
	GetStorageServiceWriter(storageType string, opts WriterOptions) StorageServiceWriterInterface
	GetStorageServiceReader(storageType string, opts ReaderOptions) StorageServiceReaderInterface
}

type storageClient struct {
	config *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) StorageClientInterface {
	var sc storageClient

	sc.config = config

	return &sc
}

func (sc *storageClient) GetStorageServiceWriter(storageType string, opts WriterOptions) StorageServiceWriterInterface {
	switch storageType {
	case "csv":
		return newCsvStorageWriter(sc.config, opts)
	case "postgres":
		return newPostgresStorageWriter(sc.config, opts)
	default:
		return nil
	}
}

func (sc *storageClient) GetStorageServiceReader(storageType string, opts ReaderOptions) StorageServiceReaderInterface {
	switch storageType {
	case "postgres":
		return newPostgresStorageReader(sc.config, opts)
	default:
		return nil
	}
}
