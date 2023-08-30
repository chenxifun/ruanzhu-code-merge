package comm

import (
	"errors"
	"os"
	"path"
)

func WriteFile(data []byte, filePath string, cover bool) error {

	if !cover {
		if _, err1 := os.Stat(filePath); !os.IsNotExist(err1) {
			return errors.New("file is exist")
		}
	}

	err := os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0600)

}

func ReadFile(path string) ([]byte, error) {

	if _, err1 := os.Stat(path); os.IsNotExist(err1) {
		return nil, errors.New("file not found")
	}

	bytes, err := os.ReadFile(path) // nolint: gas
	if err != nil {
		return nil, err
	}
	if bytes == nil {
		return nil, errors.New("file not found")
	}
	return bytes, nil

}
