package config

import (
    "os"
    "github.com/go-yaml/yaml"
)

const ConfigFilePath = "./"
const ConfigFileName = "datamkr.yml"
const ConfigFileLocation = ConfigFilePath + ConfigFileName

func CreateDatamkrConfigFile() (bool, error) {
	hasConfigFile, err := currentDirHasConfigFile()
	if err != nil {
		return true, err
	}

	if hasConfigFile {
		return true, nil
	}

	err = initDatamkrConfigFile()
	if err != nil {
		return true, err
	}
	return false, nil
}

func currentDirHasConfigFile() (bool, error) {
    _, err := os.Stat(ConfigFileLocation)
    if (os.IsNotExist(err)) {
        return false, nil
    } else if (err != nil) {
        return true, err
    } else {
        return true, nil
    }
}

func createConfigFile() (*os.File, error) {
    return os.Create(ConfigFileLocation)
}

func initDatamkrConfigFile() error {
    config := DatamakrConfig{DatasetsDir: "test/datasets"}
    f, createErr := createConfigFile()
    if (createErr != nil) {
        return createErr
    }

    defer f.Close()

    configString, marshalErr := yaml.Marshal(&config)
    if (marshalErr != nil) {
        return marshalErr
    }

    _, writeErr := f.Write(append([]byte("---\n"), configString...))
    if (writeErr != nil) {
        return writeErr
    }
    return nil
}
