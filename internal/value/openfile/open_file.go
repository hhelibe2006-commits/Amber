package openfile

import (
	"errors"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

type OpenFile []*os.File

func NewOpenFile(n int) OpenFile {
	return make(OpenFile, 0, n)
}

func (o *OpenFile) Adds(paths value.PathList) error {
	for _, path := range paths {
		if err := o.Add(path); err != nil {
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
