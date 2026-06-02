package writer

import (
	"archive/zip"
	"fmt"
	"os"

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

	tarWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(tarWriter)
	for _, file := range fileList {
		TarPack(file, tarWriter)
	}
}
