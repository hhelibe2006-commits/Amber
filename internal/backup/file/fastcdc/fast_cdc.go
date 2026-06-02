package fastcdc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func FastCDC(path string, ch chan value.Info) error {
	fmt.Println("正在备份", path)
	avg := uint64(1 << 14)
	maxByte, minByte := avg*4, avg/4
	g := NewGearHash()
	g.value = 0
	u := make([]byte, 0, maxByte)
	by := make([]byte, avg)
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
		n, err := file.Read(by)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取出错:", err)
			return err
		}
		if n == 0 {
			info := new(value.Info)
			info.Set(path, &u, sha256.Sum256(u), ch)
			g.value = 0
			break
		}
		for _, b := range by[:n] {
			g.Next(b)
			u = append(u, b)
			if uint64(len(u)) < minByte {
				continue
			}
			var t int
			if uint64(len(u)) > avg {
				t = 2
			}
			if g.value&((avg-1)>>t) == 0 || uint64(len(u)) > maxByte {
				info := new(value.Info)
				info.Set(path, &u, sha256.Sum256(u), ch)
				g.value = 0
			}
		}
	}
	return nil
}
