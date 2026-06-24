package file

import (
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/openfile"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/internal/value/chunkindex"
	"github.com/hhelibe2006-commits/Amber/pkg/cdc"
)

func BackupFile(clas *value.BackupClas, openFile openfile.OpenFile) error {
	defer func(openFile *openfile.OpenFile) {
		if err := openFile.Close(); err != nil {

		}
	}(&openFile)
	tempFiles, err := value.NewTempFiles()
	if err != nil {
		return err
	}
	defer tempFiles.Close()
	CDC := cdc.NewCDC(clas.Cdc)
	if err = chunkFiles(tempFiles, openFile, CDC); err != nil {
		return err
	}
	en := gob.NewEncoder(tempFiles.TempInfo)
	err = en.Encode(clas.Typ)
	if err != nil {
		return err
	}
	err = writer.Compress(clas.Output, tempFiles)
	if err != nil {
		return err
	}
	return nil
}

func chunkFiles(tempFiles *value.TempFiles, input openfile.OpenFile, CDC func(file *os.File) cdc.CDC) error {
	fileMap, err := chunkindex.NewChunkIndex(input)
	if err != nil {
		return err
	}
	wg := new(sync.WaitGroup)
	sem := make(chan struct{}, 1)
	errChan := make(chan error)
	for _, file := range input {
		wg.Add(1)
		file, err := input.To(file)
		if err != nil {
			errChan <- err
			continue
		}
		go func() {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			defer func(f *os.File) {
				if err = f.Close(); err != nil {
				}
			}(file)
			d := CDC(file)
			for {
				if err = handleChunk(d, fileMap, file, tempFiles); err != nil {
					if err != io.EOF {
						errChan <- err
					}
					return
				}
			}
		}()
	}
	errList := make([]error, 0)
	go func() {
		for err := range errChan {
			errList = append(errList, err)
		}
	}()
	wg.Wait()
	close(errChan)
	Writer(fileMap, tempFiles)
	return errors.Join(errList...)
}

func handleChunk(d cdc.CDC, fileMap *chunkindex.ChunkIndex, file *os.File, tempFiles *value.TempFiles) error {
	by, err := d.Next()
	var e error
	if err != nil && err != io.EOF {
		return err
	} else if err == io.EOF {
		e = err
	}
	hash := sha256.Sum256(by)
	fileMap.Mu.Lock()
	fileMap.FileMap[file.Name()].HashList = append(fileMap.FileMap[file.Name()].HashList, hash)
	if _, ok := fileMap.HashToPlace[hash]; ok {
		fileMap.Mu.Unlock()
		return nil
	}
	err = writerFile(&fileMap.HashToPlace, tempFiles.TempDate, by, &hash)
	fileMap.Mu.Unlock()
	if err != nil {
		return err
	}
	return e
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
