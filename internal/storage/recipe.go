package storage

import (
	"os"
	"time"
)

type Chunk struct {
	Name     string     `json:"name"`
	Chunk    ChunkStore `json:"chunk"`
	FileList []File     `json:"file list"`
}

type File struct {
	FilePath string      //文件路径
	Hash     [][32]byte  //文件分块的哈希值
	ModeTime time.Time   //文件最后修改时间
	Mode     os.FileMode //文件权限
}
