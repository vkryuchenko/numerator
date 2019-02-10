package filesystem

import "sync"

// FSStorage local file system implementation of Storage interface
type FSStorage struct {
	DumpPath string
	sync.Mutex
	data map[string]uint
}

// New implement Storage.New()
func New(dumpPath string) (*FSStorage, error) {
	fs := FSStorage{
		DumpPath: dumpPath,
		data:     make(map[string]uint),
	}
	err := fs.loadData()
	if err != nil {
		return &fs, err
	}
	return &fs, nil
}

// Get implement Storage.Get()
func (fs FSStorage) Get(key string) (uint, error) {
	fs.Lock()
	defer fs.Unlock()
	var value uint = 0
	_, ok := fs.data[key]
	if ok {
		value = fs.data[key]
	}
	result := value + 1
	fs.data[key] = result
	return result, nil
}

// Set implement Storage.Set()
func (fs FSStorage) Set(key string, value uint) error {
	fs.Lock()
	defer fs.Unlock()
	fs.data[key] = value
	return nil
}

// Delete implement Storage.Delete()
func (fs FSStorage) Delete(key string) error {
	fs.Lock()
	defer fs.Unlock()
	delete(fs.data, key)
	return nil
}

// ShowData implement Storage.ShowData()
func (fs FSStorage) ShowData() interface{} {
	fs.Lock()
	defer fs.Unlock()
	return fs.data
}
