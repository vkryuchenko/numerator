package filesystem

import (
	"encoding/json"
	"os"
)

func (fs *FSStorage) LoadData() error {
	f, err := os.Open(fs.DumpPath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&fs.data)
	if err != nil {
		return err
	}
	return nil
}
