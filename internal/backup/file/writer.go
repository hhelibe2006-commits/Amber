package file

import (
	"encoding/gob"
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/internal/value/chunkindex"
)

func Writer(fileMap *chunkindex.ChunkIndex, tempFiles *value.TempFiles) {
	encoder := gob.NewEncoder(tempFiles.TempFile)
	err := encoder.Encode(fileMap.FileMap)
	if err != nil {
		fmt.Println(err)
	}
	encoder = gob.NewEncoder(tempFiles.TempHash)
	err = encoder.Encode(fileMap.HashToPlace)
	if err != nil {
		fmt.Println(err)
	}
}
