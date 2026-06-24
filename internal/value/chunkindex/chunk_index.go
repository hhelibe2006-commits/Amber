package chunkindex

import (
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/openfile"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

type ChunkIndex struct {
	FileMap     map[string]*value.File
	HashToPlace map[[32]byte]*value.FilePlace
	Mu          *sync.Mutex
}

func NewChunkIndex(fileList openfile.OpenFile) (*ChunkIndex, error) {
	fileMap := new(ChunkIndex)
	fileMap.FileMap = make(map[string]*value.File, len(fileList))
	for _, file := range fileList {
		var err error
		f, err := fileList.To(file)
		if err != nil {
			return nil, err
		}
		if fileMap.FileMap[f.Name()], err = value.NewFile(f.Name()); err != nil {
			return nil, err
		}
	}
	fileMap.HashToPlace = make(map[[32]byte]*value.FilePlace)
	fileMap.Mu = new(sync.Mutex)
	return fileMap, nil
}
