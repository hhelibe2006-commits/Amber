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
	wg := new(sync.WaitGroup)
	ch := make(chan value.Info, 1000)
	go writer.Writer(ch, tempFiles)

	for _, file := range input {
		wg.Add(1)
		func() {
			err := fastcdc.FastCDC(file, ch, wg)
			if err != nil {
				fmt.Println("切割出错:", err)
			}
		}()
	}
	wg.Wait()
	close(ch)
	writer.Compress(output, tempFiles)
	return nil
}
