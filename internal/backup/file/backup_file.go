package file

import (
	"fmt"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func BackupFile(input []string, output string) error {
	tempFiles := value.NewTempFiles()
	defer tempFiles.Close()
	defer tempFiles.Remove()
	uv := new(sync.WaitGroup)
	wg := new(sync.WaitGroup)
	ch := make(chan value.Info, 1000)
	wg.Add(1)
	go writer.Writer(ch, tempFiles, wg)
	dc := make(chan struct{}, 3)
	for _, file := range input {
		uv.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			dc <- struct{}{}
			defer func() { <-dc }()
			err := fastcdc.FastCDC(file, ch)
			if err != nil {
				fmt.Println("切割出错:", err)
			}
		}(uv)
	}
	uv.Wait()
	close(ch)
	wg.Wait()
	writer.Compress(output, tempFiles)
	return nil
}
