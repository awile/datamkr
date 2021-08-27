package storage

type StorageInterface interface {
	Exists() (bool, error)
	List() ([]string, error)
	Write() error
}
