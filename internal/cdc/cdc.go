package cdc

import "os"

var pr = map[string]func(file *os.File) Cdc{}

func Add(name string, c func(file *os.File) Cdc) {
	pr[name] = c
}

type Cdc interface {
	Next() ([]byte, error)
}

func NewCDC(Type string) func(file *os.File) Cdc {
	if cdc, ok := pr[Type]; ok {
		return cdc
	}

	return nil
}
