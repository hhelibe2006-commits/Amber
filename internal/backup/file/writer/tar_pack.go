package writer

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

func TarPack(file string, tarWriter *tar.Writer) {
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
	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		fmt.Println(err)
	}
	err = tarWriter.WriteHeader(header)
	if err != nil {
		return
	}
	_, err = io.Copy(tarWriter, f)
	if err != nil {
		fmt.Println(err)
	}
}
