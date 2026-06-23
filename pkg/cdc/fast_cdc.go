package cdc

import (
	"fmt"
	"os"

	"github.com/hhelibe2006-commits/Amber/pkg/cdc/hash"
)

func init() {
	if err := Add("fastcdc", NewFastCdc); err != nil {
		fmt.Println(err)
	}
}

type FastCDC struct {
	*HashCDC
	gearHash *hash.GearHash
}

func NewFastCdc(file *os.File) CDC {
	cdc := new(FastCDC)
	cdc.HashCDC = new(HashCDC)
	cdc.proc = cdc
	cdc.gearHash = hash.NewGearHash()
	cdc.file = file

	cdc.avgBlockSize = uint64(1 << 13)
	cdc.maxBlockSize = cdc.avgBlockSize * 4
	cdc.minBlockSize = cdc.avgBlockSize / 4

	cdc.chunk = make([]byte, 0, cdc.maxBlockSize)
	cdc.readBuf = make([]byte, cdc.avgBlockSize)
	cdc.rest = make([]byte, 0)
	return cdc
}

func (cdc *FastCDC) processBytes(data []byte) bool {
	var t int
	for i, b := range data {
		cdc.chunk = append(cdc.chunk, b)
		cdc.gearHash.Next(b)
		if uint64(len(cdc.chunk)) < cdc.minBlockSize {
			continue
		}
		if uint64(len(cdc.chunk)) > cdc.avgBlockSize {
			t = 2
		}
		if cdc.gearHash.Check((cdc.avgBlockSize-1)>>t) || uint64(len(cdc.chunk)) == cdc.maxBlockSize {
			cdc.rest = append([]byte{}, data[i+1:]...)
			return true
		}
	}
	return false
}

func (cdc *FastCDC) reset() {
	cdc.gearHash.Reset()
	cdc.chunk = cdc.chunk[:0]
}
