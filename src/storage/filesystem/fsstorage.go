package filesystem

import "sync"

type FSStorage struct {
	DumpPath string
	sync.Mutex
	data map[string]uint
}

func New(dumpPath string) (*FSStorage, error) {
	fs := FSStorage{
		DumpPath: dumpPath,
		data:     make(map[string]uint),
	}
	err := fs.LoadData()
	if err != nil {
		return &fs, err
	}
	return &fs, nil
}

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

func (fs FSStorage) Set(key string, value uint) error {
	fs.Lock()
	defer fs.Unlock()
	fs.data[key] = value
	return nil
}

func (fs FSStorage) Delete(key string) error {
	fs.Lock()
	defer fs.Unlock()
	delete(fs.data, key)
	return nil
}

func (fs FSStorage) ShowData() interface{} {
	fs.Lock()
	defer fs.Unlock()
	return fs.data
}
