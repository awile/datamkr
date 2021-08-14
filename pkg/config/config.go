package config

import (
	"io"
	"log"
	"os"

	"github.com/go-yaml/yaml"
)

type DatamkrConfig struct {
	DatasetsDir string `yaml:"datasetsDir"`
}

type ConfigFactory interface {
	ConfigToByteString() ([]byte, error)
	HasConfigInDirectory() (bool, error)
	InitDatamkrConfigFile(configFile io.Writer) error
	CreateNewConfigFile() io.Writer
}

type DatamkrConfigFactory struct {
	config       DatamkrConfig
	fileLocation string
}

func NewDatamkrConfigFactory() (*DatamkrConfigFactory, error) {
	return &DatamkrConfigFactory{fileLocation: "./datamkr.yml"}, nil
}

func (dcf *DatamkrConfigFactory) ConfigToByteString() ([]byte, error) {
	return yaml.Marshal(&dcf.config)
}

func (dcf *DatamkrConfigFactory) HasConfigInDirectory() (bool, error) {
	_, err := os.Stat(dcf.fileLocation)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (dcf *DatamkrConfigFactory) InitDatamkrConfigFile(configFile io.Writer) error {
	dcf.config = DatamkrConfig{DatasetsDir: "test/datasets"}

	configByteString, err := dcf.ConfigToByteString()
	if err != nil {
		return err
	}

	_, err = configFile.Write([]byte("---\n"))
	if err != nil {
		return err
	}
	_, err = configFile.Write(configByteString)
	if err != nil {
		return err
	}
	return nil
}

func (dcf *DatamkrConfigFactory) CreateNewConfigFile() io.Writer {
	configFile, err := os.Create(dcf.fileLocation)
	if err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
	return configFile
}
