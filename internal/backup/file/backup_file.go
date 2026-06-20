package file

import (
	"crypto/sha256"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/pkg/cdc"
)

func BackupFile(clas *value.Clas) error {
	tempFiles, err := value.NewTempFiles()
	if err != nil {
		return err
	}
	defer tempFiles.Close()
	defer tempFiles.Remove()
	wg := new(sync.WaitGroup)
	CDC := cdc.NewCDC(clas.Cdc)
	cDC(tempFiles, wg, clas.Input, CDC)
	wg.Wait()
	writer.Compress(clas.Output, tempFiles)
	return nil
}

func cDC(tempFiles *value.TempFiles, wg *sync.WaitGroup, input []string, CDC func(file *os.File) cdc.CDC) {
	fileMa := make(map[string]*value.File, len(input))
	ma := make(map[[32]byte]*struct {
		A int64
		B int64
	})
	var mu *sync.Mutex
	dc := make(chan struct{}, 2)
	for _, file := range input {
		var err error
		if fileMa[file], err = value.NewFile(file); err != nil {
			return
		}
	}
	for _, file := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			dc <- struct{}{}
			defer func() { <-dc }()
			f, err := os.Open(file)
			if err != nil {
				return
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			d := CDC(f)
			for {
				by, err := d.Next()
				if err != nil {
					if err != io.EOF {
						return
					}
					return
				}
				hash := sha256.Sum256(by)
				mu.Lock()
				fileMa[file].HashList = append(fileMa[file].HashList, hash)
				if _, ok := ma[hash]; ok {
					mu.Unlock()
					continue
				}
				err = writerFile(&ma, tempFiles.TempDate, by, &hash)
				if err != nil {
					return
				}
				mu.Unlock()
			}
		}()
	}
}

func writerFile(ma *map[[32]byte]*struct {
	A int64
	B int64
}, tempDate *os.File, by []byte, hash *[32]byte) error {
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
	(*ma)[*hash] = &struct {
		A int64
		B int64
	}{A: a, B: int64(b)}
	return nil
}
