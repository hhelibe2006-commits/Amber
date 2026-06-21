package file

import (
	"crypto/sha256"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/internal/value/openfile"
	"github.com/hhelibe2006-commits/Amber/pkg/cdc"
)

func BackupFile(clas *value.Clas, openFile openfile.OpenFile) error {
	defer func(openFile *openfile.OpenFile) {
		err := openFile.Close()
		if err != nil {

		}
	}(&openFile)
	tempFiles, err := value.NewTempFiles()
	if err != nil {
		return err
	}
	defer tempFiles.Close()
	defer tempFiles.Remove()
	wg := new(sync.WaitGroup)
	CDC := cdc.NewCDC(clas.Cdc)
	chunkFiles(tempFiles, wg, openFile, CDC)
	wg.Wait()
	writer.Compress(clas.Output, tempFiles)
	return nil
}

func chunkFiles(tempFiles *value.TempFiles, wg *sync.WaitGroup, input openfile.OpenFile, CDC func(file *os.File) cdc.CDC) {
	fileMap, err := NewMap(input)
	if err != nil {

	}
	sem := make(chan struct{}, 2)
	for _, file := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			f := file
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			d := CDC(f)
			for {
				handleChunk(d, fileMap, file, tempFiles)
			}
		}()
	}
}

func handleChunk(d cdc.CDC, fileMap *ChunkIndex, file *os.File, tempFiles *value.TempFiles) {
	by, err := d.Next()
	if err != nil {
		if err != io.EOF {
			return
		}
		return
	}
	hash := sha256.Sum256(by)
	fileMap.Mu.Lock()
	fileMap.FileMap[file.Name()].HashList = append(fileMap.FileMap[file.Name()].HashList, hash)
	if _, ok := fileMap.hashToPlace[hash]; ok {
		fileMap.Mu.Unlock()
		return
	}
	err = writerFile(&fileMap.hashToPlace, tempFiles.TempDate, by, &hash)
	if err != nil {
		return
	}
	fileMap.Mu.Unlock()
}

func writerFile(ma *map[[32]byte]*value.FilePlace, tempDate *os.File, by []byte, hash *[32]byte) error {
	a, err := tempDate.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}
	va, err := writer.GzipCompress(by)
	if err != nil {
		return err
	}
	b, err := tempDate.Write(va)
	if err != nil {
		return err
	}
	(*ma)[*hash] = &value.FilePlace{A: a, B: int64(b)}
	return nil
}
