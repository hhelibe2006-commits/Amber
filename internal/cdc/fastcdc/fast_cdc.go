package fastcdc

import (
	"fmt"
	"io"
	"os"
)

type FastCdc struct {
	gearHash *GearHash
	file     *os.File
	avg      uint64
	maxByte  uint64
	minByte  uint64
}

func (g *FastCdc) Next() ([]byte, error) {
	u := make([]byte, 0, g.maxByte)
	by := make([]byte, g.avg)
	for {
		n, err := g.file.Read(by)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取出错:", err)
			return nil, err
		}
		if n == 0 {
			break
		}
		for _, b := range by[:n] {
			g.gearHash.next(b)
			u = append(u, b)
			if uint64(len(u)) < g.minByte {
				continue
			}
			var t int
			if uint64(len(u)) > g.avg {
				t = 2
			}
			if g.gearHash.value&((g.avg-1)>>t) == 0 || uint64(len(u)) > g.maxByte {
				g.gearHash.value = 0
			}
		}
	}
	return nil, nil
}

func NewFastCdc(file *os.File) *FastCdc {
	g := new(FastCdc)
	g.avg = uint64(1 << 13)
	g.maxByte = g.avg * 4
	g.minByte = g.avg / 4
	g.gearHash = new(GearHash)
	g.file = file
	return g
}
