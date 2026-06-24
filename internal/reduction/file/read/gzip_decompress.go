package read

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
)

func GzipDecompress(src []byte) ([]byte, error) {
	w, err := gzip.NewReader(bytes.NewReader(src))
	if err != nil {
		return nil, err
	}
	defer func() {
		err := w.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	var out bytes.Buffer
	_, err = io.Copy(&out, w)
	if err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}
