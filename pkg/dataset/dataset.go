package dataset

import (
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/storage"
)

type DatasetClientInterface interface {
	List() ([]string, error)
}

type DatasetClient struct {
	config         *config.DatamkrConfig
	storageService storage.StorageInterface
}

func NewWithConfig(config *config.DatamkrConfig) *DatasetClient {
	var dc DatasetClient

	dc.config = config
	dc.storageService = storage.NewLocalStorage(config)

	return &dc
}

func (dc *DatasetClient) List() ([]string, error) {
	datasets, err := dc.storageService.List()
	if err != nil {
		return nil, err
	}
	return datasets, nil
}
