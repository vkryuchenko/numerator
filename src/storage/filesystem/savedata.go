package filesystem

import (
	"encoding/json"
	"os"
)

// SaveData store data into local file system
func (fs *FSStorage) SaveData() error {
	f, err := os.OpenFile(fs.DumpPath, os.O_CREATE|os.O_WRONLY, 0775)
	if err != nil {
		return err
	}
	encoder := json.NewEncoder(f)
	err = encoder.Encode(fs.data)
	if err != nil {
		return err
	}
	return nil
}
