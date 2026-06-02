package writer

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Writer(ch chan value.Info, tempFiles *value.TempFiles, wg *sync.WaitGroup, size int) {
	defer wg.Done()
	mc := make(map[int64]struct{})
	fileMa := make(map[string]*value.File, size)
	ma := make(map[[32]byte]struct {
		A int64
		B int64
	}, 1600*size)
	dateFile := tempFiles.TempDate
	for info := range ch {
		hash := info.Hash
		if _, ok := fileMa[info.Path]; !ok {
			c, err := os.Stat(info.Path)
			if err != nil {
				fmt.Println("文件信息获取错误:", err)
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
		a, err := dateFile.Seek(0, io.SeekCurrent)
		if err != nil {
			fmt.Println(err)
		}
		if _, ok := mc[a]; ok {
			fmt.Println(a)
		} else {
			mc[a] = struct{}{}
		}
		va, err := GzipCompress(info.Value)
		if err != nil {
			fmt.Println(err)
		}
		b, err := dateFile.Write(va)
		if err != nil {
			fmt.Println(err)
		}
		ma[hash] = struct {
			A int64
			B int64
		}{A: a, B: int64(b)}
		fileMa[info.Path].HashList = append(fileMa[info.Path].HashList, hash)
	}
	write := gob.NewEncoder(tempFiles.TempHash)

	if err := write.Encode(ma); err != nil {
		fmt.Println("哈希编码错误", err)
	}
	write = gob.NewEncoder(tempFiles.TempFile)
	if err := write.Encode(fileMa); err != nil {
		fmt.Println("文件编码错误:", err)
	}
}
