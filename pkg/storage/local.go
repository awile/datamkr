package storage

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/awile/datamkr/pkg/config"
)

type LocalStorage struct {
	fileDirectory string
}

func NewLocalStorage(c *config.DatamkrConfig) *LocalStorage {
	var ls LocalStorage

	ls.fileDirectory = c.DatasetsDir

	return &ls
}

func (ls *LocalStorage) Exists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (ls *LocalStorage) List() ([]string, error) {
	var files []string

	err := filepath.Walk(ls.fileDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			name := info.Name()
			files = append(files, strings.Split(name, ".")[0])
		}
		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return files, nil
}

func (ls *LocalStorage) Write(filePath string, data []byte) error {
	file, err := ls.getFileToWrite(filePath)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	return err
}

func (ls *LocalStorage) getFileToWrite(filePath string) (io.Writer, error) {
	fileExists, err := ls.Exists(filePath)
	if err != nil {
		return nil, err
	}

	if fileExists {
		return os.OpenFile(filePath, os.O_RDWR, 0644)
	}
	return os.Create(filePath)
}
