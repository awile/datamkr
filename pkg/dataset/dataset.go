package dataset

import (
	"github.com/awile/datamkr/pkg/config"
)

type DatasetClientInterface interface {
	List() []string
}

type DatasetClient struct {
	config *config.DatamkrConfig
}

func NewWithConfig(config *config.DatamkrConfig) *DatasetClient {
	return &DatasetClient{config: config}
}

func (dc *DatasetClient) List() []string {
	return []string{"users", "organizations"}
}
