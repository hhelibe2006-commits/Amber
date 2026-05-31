package writer

import (
	"compress/gzip"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Writer(ch chan value.Info, tempFiles *value.TempFiles, wg *sync.WaitGroup) {
	defer wg.Done()
	fileMa := make(map[string]*value.File)
	ma := make(map[string]struct{})
	fileDir := tempFiles.TempDate
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
		file, err := os.Create(filepath.Join(fileDir, hash))
		if err != nil {
			fmt.Println(err)
		}
		func() {
			gWriter := gzip.NewWriter(file)
			defer func(gWriter *gzip.Writer) {
				err := gWriter.Close()
				if err != nil {
					fmt.Println(err)
				}
			}(gWriter)
			_, err = gWriter.Write(info.Value)
			if err != nil {
				return
			}
		}()
		err = file.Close()
		if err != nil {
			return
		}
		ma[hash] = struct{}{}
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
