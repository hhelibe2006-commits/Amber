package openfile

import (
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

type OpenFile []string

func NewOpenFile(n int) OpenFile {
	return make(OpenFile, 0, n)
}

func (o *OpenFile) Adds(paths value.PathList) error {
	for _, path := range paths {
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			if err = o.Add(path); err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func (o *OpenFile) Close() error {
	return nil
}

func (o *OpenFile) Add(path string) error {
	*o = append(*o, path)
	return nil
}

func (o *OpenFile) To(path string) (*os.File, error) {
	file, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return file, err
	}
	return file, nil
}
