package file

import (
	"compress/gzip"
	"io"
	"os"
)

func WriteFile(file *os.File, content []byte) (int64, error) {
	writer := gzip.NewWriter(file)
	defer func(writer *gzip.Writer) {
		err := writer.Close()
		if err != nil {
			return
		}
	}(writer)
	_, err := writer.Write(content)
	if err != nil {
		return 0, err
	}
	offset, err := file.Seek(0, io.SeekCurrent)
	return offset, nil
}
