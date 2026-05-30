package fastcdc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"sync"
)

func FastCDC(path string, ch chan Info, wg *sync.WaitGroup) error {
	fmt.Println("正在备份", path)
	defer wg.Done()
	avg := uint64(65536)
	g := NewGearHash()
	g.value = 0
	u := make([]byte, 0)
	by := make([]byte, 1024)
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("文件出错:", err)
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("关闭出错", err)
			return
		}
	}()
	for {
		_, err := file.Read(by)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取出错:", err)
			return err
		}
		if err == io.EOF {
			info := new(Info)
			info.Path = path
			hash := sha256.Sum256(u)
			info.Hash = string(hash[:])
			info.Value = u
			ch <- *info
			u = u[:0]
			break
		}
		for _, b := range by {
			g.Next(b)
			u = append(u, b)
			if g.value&avg == 0 {
				info := new(Info)
				info.Path = path
				hash := sha256.Sum256(u)
				info.Hash = string(hash[:])
				info.Value = u
				ch <- *info
				u = u[:0]
			}
		}
	}
	return nil
}
