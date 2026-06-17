package file

import (
	"crypto/sha256"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/cdc/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func BackupFile(input []string, output string) error {
	tempFiles := value.NewTempFiles()
	defer tempFiles.Close()
	defer tempFiles.Remove()
	uv := new(sync.WaitGroup)
	wg := new(sync.WaitGroup)
	ch := make(chan value.Info, 100)
	wg.Add(1)
	//go writer.Writer(ch, tempFiles, wg, len(input))
	dc := make(chan struct{}, 2)
	for _, file := range input {
		uv.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			dc <- struct{}{}
			f, err := os.Open(file)
			if err != nil {

			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {

				}
			}(f)
			fastCDC := fastcdc.NewFastCdc(f)
			for {
				by, err := fastCDC.Next()
				if err != nil {
					break
				}
				_ = sha256.Sum256(by)
			}
			defer func() { <-dc }()
			//err := fastcdc.FastCDC(file, ch)
			//if err != nil {
			//	fmt.Println("切割出错:", err)
			//}
		}(uv)
	}
	uv.Wait()
	close(ch)
	wg.Wait()
	writer.Compress(output, tempFiles)
	return nil
}
