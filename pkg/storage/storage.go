package storage

import (
	"fmt"

	"github.com/awile/datamkr/pkg/config"
)

type StorageClientInterface interface {
	GetStorageService(storageType string) (StorageServiceInterface, error)
}

type StorageServiceInterface interface {
	Init(args interface{}) error
	Write(data interface{}) error
	WriteAll(data interface{}) error
	Close() error
}

type StorageArgs struct {
	FileName string
	IsWriter bool
}

type storageClient struct {
	config *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) StorageClientInterface {
	var sc storageClient

	sc.config = config

	return &sc
}

func (sc *storageClient) GetStorageService(storageType string) (StorageServiceInterface, error) {
	switch storageType {
	case "csv":
		return newCsvStorageWithConfig(sc.config), nil
	default:
		return nil, fmt.Errorf("No storage server %s found.", storageType)
	}
}
