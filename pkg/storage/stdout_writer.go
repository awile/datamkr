package storage

import (
	"encoding/json"
	"log"

	"github.com/awile/datamkr/pkg/config"
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

func (std *stdoutStorageServiceWriter) Write(data map[string]interface{}) error {
	formattedJson, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	std.logger.Printf("%s\n", string(formattedJson))
	return nil
}

func (std *stdoutStorageServiceWriter) WriteAll(data []map[string]interface{}) error {
	return nil
}

func (std *stdoutStorageServiceWriter) Close() error {
	return nil
}
