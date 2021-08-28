package storage

import (
	"strings"

	"github.com/awile/datamkr/pkg/config"
)

type StorageInterface interface {
	Exists() (bool, error)
	List() ([]string, error)
	Write() error
}

type StorageClientInterface interface {
	Init(filePath string) error
	Insert(filePath string, row []string) error
}

type StorageClient struct {
	config         *config.DatamkrConfig
	storageService *LocalStorage
}

func NewWithConfig(config *config.DatamkrConfig) *StorageClient {
	var sc StorageClient

	sc.config = config
	sc.storageService = NewLocalStorage(config)

	return &sc
}

func (sc *StorageClient) Init(filePath string) error {
	fileExists, err := sc.storageService.Exists(filePath)
	if err != nil {
		return err
	}
	if !fileExists {
		sc.storageService.Create(filePath)
	}
	return nil
}

func (sc *StorageClient) Insert(filePath string, row []string) error {
	stringArray := strings.Join(row, ",") + "\n"
	return sc.storageService.Write(filePath, []byte(stringArray))
}
