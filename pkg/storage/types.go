package storage

import (
	"github.com/awile/datamkr/pkg/dataset"
	"github.com/awile/datamkr/pkg/maker/providers"
)

type StorageServiceWriterInterface interface {
	Init() error
	Write(data map[string]providers.ProviderField) error
	WriteAll(data []map[string]providers.ProviderField) error
	Close() error
}

type StorageServiceReaderInterface interface {
	Init() error
	GetDatasetDefinition() (dataset.DatasetDefinition, error)
	Close() error
}

type WriterOptions struct {
	Id                string
	SecondaryId       string
	FieldKeys         []string
	DatasetDefinition dataset.DatasetDefinition
}

func CreateWriterOptions() WriterOptions {
	return WriterOptions{}
}

type ReaderOptions struct {
	Id          string
	SecondaryId string
}

func CreateReaderOptions() ReaderOptions {
	return ReaderOptions{}
}
