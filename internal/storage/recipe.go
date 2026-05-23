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

func NewChunk() *Chunk {
	ch := new(Chunk)
	ch.Chunk = NewChunkStore()
	ch.FileList = make([]File, 0)
	return ch
}

type File struct {
	FilePath string      //文件路径
	Hash     []string    //文件分块的哈希值
	ModeTime time.Time   //文件最后修改时间
	Mode     os.FileMode //文件权限
}

func NewFile() *File {
	ch := new(File)
	ch.Hash = make([]string, 0)
	return ch
}
