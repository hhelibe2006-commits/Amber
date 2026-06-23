package cdc

import (
	"errors"
	"os"
)

var pr = map[string]func(file *os.File) CDC{}

func Add(name string, c func(file *os.File) CDC) error {
	if _, ok := pr[name]; !ok {
		pr[name] = c
	} else {
		return errors.New("重名了" + name)
	}
	return nil
}

type CDC interface {
	Next() ([]byte, error)
}

func NewCDC(Type string) func(file *os.File) CDC {
	if cdc, ok := pr[Type]; ok {
		return cdc
	}

	return nil
}
