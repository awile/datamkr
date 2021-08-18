package storage

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/awile/datamkr/pkg/config"
)

type LocalStorageInterface interface {
	List() ([]string, error)
}

type LocalStorage struct {
	fileDirectory string
}

func NewLocalStorage(c *config.DatamkrConfig) LocalStorageInterface {
	var ls LocalStorage

	ls.fileDirectory = c.DatasetsDir

	return &ls
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
