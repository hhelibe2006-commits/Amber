package cdc

import (
	"fmt"
	"io"
	"os"
)

func init() {
	Add("fastcdc", NewFastCdc)
}

func NewFastCdc(file *os.File) Cdc {
	g := new(FastCdc)
	g.u = make([]byte, 0, g.maxByte)
	g.avg = uint64(1 << 13)
	g.maxByte = g.avg * 4
	g.minByte = g.avg / 4
	g.gearHash = NewGearHash()
	g.file = file
	g.by = make([]byte, g.avg)
	return g
}

type FastCdc struct {
	gearHash *GearHash
	file     *os.File
	avg      uint64
	maxByte  uint64
	minByte  uint64
	u        []byte
	by       []byte
}

func (g *FastCdc) Next() ([]byte, error) {
	for {
		n, err := g.file.Read(g.by)
		if err != nil && err != io.EOF {
			fmt.Println("文件读取出错:", err)
			return nil, err
		} else if n == 0 {
			break
		} else if err == io.EOF {
			return append(g.u, g.by[:n]...), err
		}
		var t int
		for _, b := range g.by[:n] {
			g.gearHash.next(b)
			g.u = append(g.u, b)
			if uint64(len(g.u)) < g.minByte {
				continue
			}
			if uint64(len(g.u)) > g.avg {
				t = 2
			}
			if g.gearHash.value&((g.avg-1)>>t) == 0 || uint64(len(g.u)) > g.maxByte {
				g.gearHash.value = 0
				break
			}
		}
		if g.gearHash.value == 0 {
			break
		}
	}
	u := g.u[:len(g.u)]
	g.u = g.u[:0]
	return u, nil
}
