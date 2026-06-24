package file

import (
	"os"
)

func CreateFile(path string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	err = os.Remove(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
