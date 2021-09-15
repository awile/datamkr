package config

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"
)

func NewConfig() *DatamkrConfig {
	var config DatamkrConfig

	settings := viper.GetStringMap("datamkr")
	if len(settings) == 0 {
		return &config
	}

	config.DatasetsDir = settings["datasetsdir"].(string)

	if settings["storage"] != nil {
		storageAliases := make(map[string]StorageAlias)
		configStorageMap := settings["storage"].(map[string]interface{})
		for storageName := range configStorageMap {
			var storageAlias StorageAlias
			storageMap := configStorageMap[storageName].(map[string]interface{})
			storageAlias.ConnectionString = storageMap["connection"].(string)
			storageAlias.Type = storageMap["type"].(string)
			storageAliases[storageName] = storageAlias
		}
		config.StorageAliases = storageAliases
	}

	return &config
}

type ConfigFactory interface {
	GetConfig() (*DatamkrConfig, error)
	ConfigToByteString() ([]byte, error)
	HasConfigInDirectory() (bool, error)
	InitDatamkrConfigFile(configFile io.Writer) error
	CreateNewConfigFile() io.Writer
}

type DatamkrConfigFactory struct {
	config       DatamkrConfig
	FileLocation string
}

func NewDatamkrConfigFactory() (*DatamkrConfigFactory, error) {
	return &DatamkrConfigFactory{FileLocation: "./.datamkr.yml"}, nil
}

func (dcf *DatamkrConfigFactory) ConfigToByteString() ([]byte, error) {
	wrappedConfig := make(map[string]DatamkrConfig)
	wrappedConfig["datamkr"] = dcf.config
	return yaml.Marshal(&wrappedConfig)
}

func (dcf *DatamkrConfigFactory) HasConfigInDirectory() (bool, error) {
	_, err := os.Stat(dcf.FileLocation)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (dcf *DatamkrConfigFactory) InitDatamkrConfigFile(configFile io.Writer) error {
	dcf.config = DatamkrConfig{DatasetsDir: "./datasets", Version: 1}

	configByteString, err := dcf.ConfigToByteString()
	if err != nil {
		return err
	}
	err = viper.ReadConfig(bytes.NewBuffer(configByteString))
	if err != nil {
		return err
	}

	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

func (dcf *DatamkrConfigFactory) CreateNewConfigFile() io.Writer {
	configFile, err := os.Create(dcf.FileLocation)
	if err != nil {
		log.Fatal(err)
	}
	return configFile
}

func (dcf *DatamkrConfigFactory) GetConfig() (*DatamkrConfig, error) {
	return NewConfig(), nil
}
