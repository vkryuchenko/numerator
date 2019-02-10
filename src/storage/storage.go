package storage

// Storage represent data workaround interface
type Storage interface {
	ShowData() interface{}
	Get(key string) (uint, error)
	Set(key string, value uint) error
	Delete(key string) error
}
