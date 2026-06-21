package openfile

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

type OpenFile []*os.File

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
	errList := make([]error, 0)
	for _, file := range *o {
		if err := file.Close(); err != nil {
			errList = append(errList, err)
		}
	}
	return errors.Join(errList...)
}
