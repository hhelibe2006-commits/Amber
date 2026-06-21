package file

import (
	"encoding/gob"
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Writer(fileMap *ChunkIndex, tempFiles *value.TempFiles) {
	encoder := gob.NewEncoder(tempFiles.TempFile)
	err := encoder.Encode(fileMap.FileMap)
	if err != nil {
		fmt.Println(err)
	}
	encoder = gob.NewEncoder(tempFiles.TempHash)
	err = encoder.Encode(fileMap.hashToPlace)
	if err != nil {
		fmt.Println(err)
	}
}
