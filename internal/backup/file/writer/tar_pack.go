package writer

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func TarPack(file string, tarWriter *zip.Writer) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	info, err := f.Stat()
	if err != nil {
		fmt.Println(err)
	}
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		fmt.Println(err)
	}
	z, err := tarWriter.CreateHeader(header)
	if err != nil {
		return
	}
	_, err = io.Copy(z, f)
	if err != nil {
		fmt.Println(err)
	}
}
