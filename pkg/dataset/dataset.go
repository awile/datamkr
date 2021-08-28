package dataset

import (
	"fmt"

	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/storage"
	"github.com/go-yaml/yaml"
)

type DatasetClientInterface interface {
	Add(name string, definition DatasetDefinition) error
	Get(name string) (DatasetDefinition, error)
	List() ([]string, error)
}

type DatasetClient struct {
	config         *config.DatamkrConfig
	storageService *storage.LocalStorage
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

func (dc *DatasetClient) Add(name string, definition DatasetDefinition) error {
	filePath := dc.getDatasetPath(name)

	fileExists, err := dc.storageService.Exists(filePath)
	if err != nil {
		return err
	}
	if fileExists {
		return fmt.Errorf("Dataset already exists at: %s\n", filePath)
	}

	fileContents, err := dc.getDatasetFileContent(name, definition)
	if err != nil {
		return err
	}

	return dc.storageService.Write(filePath, fileContents)
}

func (dc *DatasetClient) getDatasetFileContent(name string, definition DatasetDefinition) ([]byte, error) {
	definitionDescription := make(map[string]DatasetDefinition)
	definitionDescription[name] = definition
	return yaml.Marshal(&definitionDescription)
}

func (dc *DatasetClient) getDatasetPath(name string) string {
	return fmt.Sprintf("%s/%s.yml", dc.config.DatasetsDir, name)
}

func (dc *DatasetClient) Get(name string) (DatasetDefinition, error) {
	var datasetDefinition DatasetDefinition
	filePath := dc.getDatasetPath(name)

	fileExists, err := dc.storageService.Exists(filePath)
	if err != nil {
		return datasetDefinition, err
	}
	if !fileExists {
		return datasetDefinition, fmt.Errorf("Dataset does not already exists at %s\ncreate with: datamkr add %s\n", filePath, name)
	}

	definitionBytes, err := dc.storageService.Read(filePath)
	if err != nil {
		return datasetDefinition, err
	}

	var definitionFile map[string]DatasetDefinition
	err = yaml.Unmarshal(definitionBytes, &definitionFile)
	if err != nil {
		return datasetDefinition, err
	}
	datasetDefinition = definitionFile[name]
	return datasetDefinition, nil
}
