package file

import (
	"crypto/sha256"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/cdc"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func BackupFile(input []string, output string) error {
	tempFiles := value.NewTempFiles()
	defer tempFiles.Close()
	defer tempFiles.Remove()
	fileMa := make(map[string]*value.File, len(input))
	ma := make(map[[32]byte]*struct {
		A int64
		B int64
	})
	var mu sync.Mutex
	wg := new(sync.WaitGroup)
	dc := make(chan struct{}, 2)
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
			fu := cdc.NewCDC("fastcdc")
			CDC := fu(f)
			for {
				by, err := CDC.Next()
				if err != nil {
					break
				}
				hash := sha256.Sum256(by)
				if _, ok := fileMa[file]; ok {
					c, err := os.Stat(file)
					if err != nil {

					}
					fileMa[file] = &value.File{
						HashList: make([][32]byte, 0, 1),
						Path:     file,
						Mode:     c.Mode(),
						ModTime:  c.ModTime(),
					}
				}
				mu.Lock()
				if _, ok := ma[hash]; ok {
					fileMa[file].HashList = append(fileMa[file].HashList, hash)
					mu.Unlock()
					continue
				}
				a, err := tempFiles.TempDate.Seek(0, io.SeekCurrent)
				if err != nil {
				}
				ma[hash] = &struct {
					A int64
					B int64
				}{A: a, B: 0}
				mu.Unlock()
				fileMa[file].HashList = append(fileMa[file].HashList, hash)
				va, err := writer.GzipCompress(by)
				if err != nil {
				}
				b, err := tempFiles.TempDate.Write(va)
				if err != nil {
				}
				mu.Lock()
				ma[hash].B = int64(b)
				mu.Unlock()
			}
		}()
	}
	wg.Wait()
	writer.Compress(output, tempFiles)
	return nil
}
