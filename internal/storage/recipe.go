package storage

import (
	"os"
	"path/filepath"
	"time"
)

type Chunk struct {
	Name     string `json:"name"`
	FileList []File `json:"file list"`
}

func NewChunk() *Chunk {
	ch := new(Chunk)
	ch.FileList = make([]File, 0)
	return ch
}

type File struct {
	FilePath string            //文件路径和文件名
	Hash     map[uint64]string //文件分块的哈希值
	ModeTime time.Time         //文件最后修改时间
	Mode     os.FileMode       //文件权限
}

func (file *File) Set(path string) error {
	if info, err := os.Stat(path); err != nil {
		return err
	} else {
		file.Mode = info.Mode()
		file.ModeTime = info.ModTime()
	}
	if abs, err := filepath.Abs(path); err != nil {
		return err
	} else {
		file.FilePath = abs
	}
	return nil
}

func NewFile() *File {
	ch := new(File)
	ch.Hash = make(map[uint64]string)
	return ch
}
