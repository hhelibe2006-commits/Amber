package writer

import (
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Writer(ch chan value.Info, tempFiles *value.TempFiles, wg *sync.WaitGroup) {
	defer wg.Done()
	fileMa := make(map[string]*value.File)
	ma := make(map[string]struct {
		a int64
		b int64
	})
	dateFile := tempFiles.TempDate
	for info := range ch {
		hash := hex.EncodeToString(info.Hash[:])
		if _, ok := fileMa[info.Path]; !ok {
			c, err := os.Stat(info.Path)
			if err != nil {
				fmt.Println("文件信息获取错误:", err)
				return
			}
			fileMa[info.Path] = &value.File{
				HashList: make([]string, 0, 1),
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
		b, err := dateFile.Write(info.Value)
		if err != nil {
			return
		}
		ma[hash] = struct {
			a int64
			b int64
		}{a: a, b: int64(b)}
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
