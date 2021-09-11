package storage

import (
	"github.com/awile/datamkr/pkg/dataset"
)

type StorageServiceWriterInterface interface {
	Init() error
	Write(data map[string]interface{}) error
	WriteAll(data []map[string]interface{}) error
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
