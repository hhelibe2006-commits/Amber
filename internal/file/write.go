package file

import (
	"compress/gzip"
	"io"
	"os"
)

func WriteFile(f *os.File, content []byte) (int64, error) {
	writer := gzip.NewWriter(f)
	defer func(writer *gzip.Writer) {
		err := writer.Close()
		if err != nil {
			return
		}
	}(writer)
	offset, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}
	_, err = writer.Write(content)
	if err != nil {
		return 0, err
	}
	return offset, nil
}
