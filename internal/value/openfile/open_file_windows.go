package openfile

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/value"
	"golang.org/x/sys/windows"
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

func (o *OpenFile) Add(path string) error {
	filePath, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	handle, err := windows.CreateFile(
		filePath,
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ,
		nil,
		windows.OPEN_ALWAYS,
		windows.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return err
	}
	file := os.NewFile(uintptr(handle), path)
	*o = append(*o, file)
	return nil
}

func (o *OpenFile) To(file *os.File) (*os.File, error) {
	return file, nil
}
