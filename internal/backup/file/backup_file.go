package file

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
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
	ch := make(chan fastcdc.Info, 100)
	go writer.Writer(ch, tempFiles)

	for _, file := range input {
		wg.Add(1)
		err := fastcdc.FastCDC(file, ch, wg)
		if err != nil {
			fmt.Println("切割出错:", err)
			return err
		}
	}
	fileList := make([]string, 0, 3)
	fileList = append(fileList, tempFiles.TempDate.Name())
	fileList = append(fileList, tempFiles.TempFile.Name())
	fileList = append(fileList, tempFiles.TempHash.Name())
	wg.Wait()
	close(ch)
	tempFiles.Close()
	zipFile, err := os.Create(output)
	if err != nil {
		return err
	}
	w := zip.NewWriter(zipFile)
	defer func(w *zip.Writer) {
		err := w.Close()
		if err != nil {
			return
		}
	}(w)
	for _, file := range fileList {
		f, err := filepath.Abs(file)
		if err != nil {
			return err
		}
		zipFile, err := os.Open(f)
		if err != nil {
			return err
		}
		defer func(zipFile *os.File) {
			err := zipFile.Close()
			if err != nil {
				return
			}
		}(zipFile)
		u, err := w.Create(zipFile.Name())
		if err != nil {
			return err
		}
		if _, err = io.Copy(u, zipFile); err != nil {
			return err
		}
		err = zipFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
