package file

import (
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/internal/value/openfile"
)

type ChunkIndex struct {
	FileMap     map[string]*value.File
	hashToPlace map[[32]byte]*value.FilePlace
	Mu          *sync.Mutex
}

func NewChunkIndex(fileList openfile.OpenFile) (*ChunkIndex, error) {
	fileMap := new(ChunkIndex)
	fileMap.FileMap = make(map[string]*value.File, len(fileList))
	for _, file := range fileList {
		var err error
		if fileMap.FileMap[file.Name()], err = value.NewFile(file.Name()); err != nil {
			return nil, err
		}
	}
	fileMap.hashToPlace = make(map[[32]byte]*value.FilePlace)
	fileMap.Mu = new(sync.Mutex)
	return fileMap, nil
}
