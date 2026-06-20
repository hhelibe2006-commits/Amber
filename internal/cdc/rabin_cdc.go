package cdc

import (
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/cdc/hash"
)

func init() {
	Add("rabincdc", NewRabinCDC)
}

type RabinCDC struct {
	*HashCDC
	rabinHash *hash.RabinHash
}

func NewRabinCDC(file *os.File) CDC {
	rabinCDC := new(RabinCDC)
	rabinCDC.HashCDC = new(HashCDC)
	rabinCDC.proc = rabinCDC
	rabinCDC.minBlockSize = 1 << 10
	rabinCDC.maxBlockSize = 1 << 18
	rabinCDC.avgBlockSize = 1 << 12
	rabinCDC.file = file
	rabinCDC.chunk = make([]byte, 0, rabinCDC.maxBlockSize)
	rabinCDC.readBuf = make([]byte, rabinCDC.avgBlockSize)
	rabinCDC.rest = make([]byte, 0)
	rabinCDC.rabinHash = hash.NewRabinHash()
	return rabinCDC
}

func (cdc *RabinCDC) processBytes(data []byte) bool {
	for i, b := range data {
		cdc.chunk = append(cdc.chunk, b)
		cdc.rabinHash.Next(b)
		if uint64(len(cdc.chunk)) < cdc.minBlockSize {
			continue
		}
		if cdc.rabinHash.Check(cdc.avgBlockSize-1) || uint64(len(cdc.chunk)) == cdc.maxBlockSize {
			cdc.rest = append([]byte{}, data[i+1:]...)
			return true
		}
	}
	return false
}

func (cdc *RabinCDC) reset() {
	cdc.rabinHash.Reset()
	cdc.chunk = cdc.chunk[:0]
}
