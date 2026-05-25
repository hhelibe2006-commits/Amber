package storage

import (
	"os"
	"path/filepath"
	"time"
)

type Manifest struct {
	Name     string     `json:"name"`
	FileList []FileMeta `json:"file list"`
}

func NewManifest() *Manifest {
	ch := new(Manifest)
	ch.FileList = make([]FileMeta, 0)
	return ch
}

type FileMeta struct {
	FilePath string            //文件路径和文件名
	Hash     map[uint64]string //文件分块的哈希值
	ModeTime time.Time         //文件最后修改时间
	Mode     os.FileMode       //文件权限
}

func (file *FileMeta) Set(path string) error {
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

func NewFileMeta() *FileMeta {
	ch := new(FileMeta)
	ch.Hash = make(map[uint64]string)
	return ch
}
