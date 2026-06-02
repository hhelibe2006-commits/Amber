package writer

import (
	"bytes"
	"compress/gzip"
	"fmt"
)

func GzipCompress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer func(w *gzip.Writer) {
		err := w.Close()
		if err != nil {
			fmt.Printf("字节压缩关闭错误: %s\n", err)
		}
	}(w)
	if _, err := w.Write(data); err != nil {
		return nil, err
	}
	if err := w.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
