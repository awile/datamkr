package client

import (
	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/dataset"
)

type Interface interface {
	Datasets() dataset.DatasetClientInterface
}

type Client struct {
	datasets *dataset.DatasetClient
}

func NewWithConfig(config *config.DatamkrConfig) *Client {
	var client Client

	client.datasets = dataset.NewWithConfig(config)

	return &client
}

func (c *Client) Datasets() dataset.DatasetClientInterface {
	return c.datasets
}
