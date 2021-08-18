package storage

type StorageInterface interface {
	List() ([]string, error)
}
