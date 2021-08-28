package client

import (
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/awile/datamkr/pkg/maker"
	"github.com/awile/datamkr/pkg/storage"
)

type Interface interface {
	Datasets() dataset.DatasetClientInterface
	Maker() maker.MakerClientInterface
	Storage() storage.StorageClientInterface
}

type Client struct {
	datasets *dataset.DatasetClient
	maker    *maker.MakerClient
	storage  *storage.StorageClient
}

func NewWithConfig(config *config.DatamkrConfig) *Client {
	var client Client

	client.datasets = dataset.NewWithConfig(config)
	client.maker = maker.NewWithConfig(config)
	client.storage = storage.NewWithConfig(config)

	return &client
}

func (c *Client) Datasets() dataset.DatasetClientInterface {
	return c.datasets
}

func (c *Client) Maker() maker.MakerClientInterface {
	return c.maker
}

func (c *Client) Storage() storage.StorageClientInterface {
	return c.storage
}
