package fastcdc

import (
	"fmt"
	"io"
	"math/rand"
	"os"
	"sync"
)

type GearHash struct {
	table [256]uint64
	l     uint64
}

func (g *GearHash) Next(b byte) {
	g.l = (g.l << 1) + g.table[b]
}

func NewGearHash(i int64) *GearHash {
	g := new(GearHash)
	rng := rand.New(rand.NewSource(i))
	for i := 0; i < 256; i++ {
		g.table[i] = rng.Uint64()
	}
	return g
}

func FastCDC(i int64, input string, wg *sync.WaitGroup) {
	defer wg.Done()
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	gearHash := NewGearHash(i)
	info, _ := os.Stat(input)
	li := make([]byte, 0, info.Size())
	file, _ := os.Open(input)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	buf := make([]byte, 0, 1024)
	u := make([]byte, 0, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println(err)
		}
		for i := 0; i < n; i++ {
			u = append(u, buf[i])
			gearHash.Next(buf[i])
			if len(u) < minByte {
				continue
			}
			if gearHash.l&1 == 0 {
				li = append(li, u...)
				u = u[:0]
			}
			if len(u) > maxByte {
				li = append(li, u...)
				u = u[:0]
			}
		}
	}
}
