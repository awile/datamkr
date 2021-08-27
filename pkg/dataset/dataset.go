package dataset

import (
	"fmt"

	"github.com/awile/datamkr/pkg/config"
	"github.com/awile/datamkr/pkg/storage"
	"github.com/go-yaml/yaml"
)

type DatasetClientInterface interface {
	Add(name string) error
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

func (dc *DatasetClient) Add(name string) error {
	filePath := fmt.Sprintf("%s/%s.yml", dc.config.DatasetsDir, name)

	fileExists, err := dc.storageService.Exists(filePath)
	if err != nil {
		return err
	}
	if fileExists {
		return fmt.Errorf("Dataset already exists at: %s\n", filePath)
	}

	fileContents, err := dc.getDatasetFileContent(name)
	if err != nil {
		return err
	}

	return dc.storageService.Write(filePath, fileContents)
}

func (dc *DatasetClient) getDatasetFileContent(name string) ([]byte, error) {
	fieldMap := make(map[string]DatasetDefinitionField)
	fieldMap["id"] = DatasetDefinitionField{Type: "uuid"}
	fieldMap["full_name"] = DatasetDefinitionField{Type: "string"}
	fieldMap["age"] = DatasetDefinitionField{Type: "number"}

	datasetDefinition := DatasetDefinition{Fields: fieldMap}

	definitionDescription := make(map[string]DatasetDefinition)
	definitionDescription[name] = datasetDefinition
	return yaml.Marshal(&definitionDescription)
}
