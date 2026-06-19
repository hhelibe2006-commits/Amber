package cdc

import "os"

var pr = map[string]func(file *os.File) CDC{}

func Add(name string, c func(file *os.File) CDC) {
	pr[name] = c
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
