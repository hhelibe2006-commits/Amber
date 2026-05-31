package writer

import (
	"bufio"
	"compress/gzip"
	"encoding/gob"
	"fmt"
	"io"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Writer(ch chan value.Info, tempFiles *value.TempFiles) {
	fileMa := make(map[string]*value.File)
	ma := make(map[[32]byte]struct {
		a int64
		b int64
	})
	dateWrite := gzip.NewWriter(tempFiles.TempDate)
	bufferedWriter := bufio.NewWriterSize(dateWrite, 16*1024*1024)
	defer func(dateWrite *gzip.Writer) {
		err := dateWrite.Close()
		if err != nil {
			return
		}
	}(dateWrite)
	for info := range ch {
		hash := info.Hash
		if _, ok := fileMa[info.Path]; !ok {
			c, err := os.Stat(info.Path)
			if err != nil {
				fmt.Println("文件信息获取错误:", err)
				return
			}
			fileMa[info.Path] = &value.File{
				HashList: make([][32]byte, 0, 1),
				Path:     info.Path,
				ModTime:  c.ModTime(),
				Mode:     c.Mode(),
			}
		}
		if _, ok := ma[hash]; ok {
			fileMa[info.Path].HashList = append(fileMa[info.Path].HashList, hash)
			continue
		}
		a, err := tempFiles.TempDate.Seek(0, io.SeekCurrent)
		if err != nil {
			fmt.Println("切片起始位置获取错误")
			return
		}
		if _, err := bufferedWriter.Write(info.Value); err != nil {
			return
		}
		b, err := tempFiles.TempDate.Seek(0, io.SeekCurrent)
		if err != nil {
			fmt.Println("切片末尾获取错误")
			return
		}
		ma[hash] = struct {
			a int64
			b int64
		}{a: a, b: b}
		fileMa[info.Path].HashList = append(fileMa[info.Path].HashList, hash)
	}
	write := gob.NewEncoder(tempFiles.TempHash)

	if err := write.Encode(ma); err != nil {
		return
	}
	write = gob.NewEncoder(tempFiles.TempFile)
	if err := write.Encode(fileMa); err != nil {
		return
	}
}
