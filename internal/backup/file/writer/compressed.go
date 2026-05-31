package writer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Compress(output string, tempFiles *value.TempFiles) {
	fileList := tempFiles.FileList()
	tempFiles.Close()
	zipFile, err := os.Create(output)
	if err != nil {
		fmt.Println(err)
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(zipFile)

	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(zipWriter)

	for _, file := range fileList {
		func() {
			f, err := os.Open(file)
			if err != nil {
				fmt.Println(err)
			}
			info, err := f.Stat()
			header, err := zip.FileInfoHeader(info)
			header.Method = zip.Store
			name, err := filepath.Abs(file)
			if err != nil {
				fmt.Println(err)
			}
			header.Name = filepath.Base(name)
			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				fmt.Println(err)
			}
			_, err = io.Copy(writer, f)
			if err != nil {
				fmt.Println(err)
			}
			defer func(f *os.File) {
				err := f.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(f)
		}()
	}
}
