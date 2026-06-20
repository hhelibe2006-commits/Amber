package value

import (
	"os"
	"time"
)

type File struct {
	Path     string
	HashList [][32]byte
	Mode     os.FileMode
	ModTime  time.Time
}

func NewFile(path string) (*File, error) {
	file := new(File)
	file.Path = path
	info, err := os.Stat(file.Path)
	if err != nil {
		return nil, err
	}
	file.ModTime = info.ModTime()
	file.Mode = info.Mode()
	file.HashList = make([][32]byte, 0, 1)
	return file, nil
}
