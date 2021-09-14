package storage

import (
	"encoding/json"
	"log"

	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/maker/providers"
)

type stdoutStorageServiceWriter struct {
	logger *log.Logger
}

func newStdoutStorageServiceWriter(config *config.DatamkrConfig) StorageServiceWriterInterface {
	var storageService stdoutStorageServiceWriter

	return &storageService
}

func (std *stdoutStorageServiceWriter) Init() error {
	log.SetFlags(0)
	std.logger = log.Default()
	return nil
}

func (std *stdoutStorageServiceWriter) Write(data map[string]providers.ProviderField) error {
	stringMap := make(map[string]string, len(data))
	for column := range data {
		stringMap[column] = data[column].String()
	}

	formattedJson, err := json.MarshalIndent(stringMap, "", "  ")
	if err != nil {
		return err
	}
	std.logger.Printf("%s\n", string(formattedJson))
	return nil
}

func (std *stdoutStorageServiceWriter) WriteAll(data []map[string]providers.ProviderField) error {
	return nil
}

func (std *stdoutStorageServiceWriter) Close() error {
	return nil
}
