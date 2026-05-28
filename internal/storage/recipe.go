package storage

import (
	"os"
	"time"
)

type Manifest struct {
	Name         string
	FileNameList map[string]*FileMeta
}

func NewManifest() *Manifest {
	m := new(Manifest)
	m.FileNameList = make(map[string]*FileMeta)
	return m
}

type FileMeta struct {
	FilePath string //文件路径和文件名
	//Hash     map[uint64]string //文件分块的哈希值
	ModeTime  time.Time   //文件最后修改时间
	Mode      os.FileMode //文件权限
	Locations []int64
	Lengths   []int64
}

func NewFileMeta() *FileMeta {
	fileMeta := new(FileMeta)
	fileMeta.Lengths = make([]int64, 0)
	fileMeta.Locations = make([]int64, 0)
	return fileMeta
}
