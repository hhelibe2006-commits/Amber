package writer

import (
	"archive/zip"
	"fmt"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Compress(output string, tempFiles *value.TempFiles) error {
	fileList := tempFiles.FileList()
	zipFile, err := os.Create(output)
	if err != nil {
		return err
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
		err = ZipPack(file, zipWriter)
		if err != nil {
			return err
		}
	}
	return nil
}
