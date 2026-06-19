package cdc

import (
	"io"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/cdc/hash"
)

func init() {
	Add("rabincdc", NewRabinCDC)
}

type RabinCDC struct {
	rabinHash *hash.RabinHash
	file      *os.File

	avgBlockSize uint64
	maxBlockSize uint64
	minBlockSize uint64

	chunk   []byte
	readBuf []byte
	rest    []byte
}

func NewRabinCDC(file *os.File) CDC {
	rabinCDC := new(RabinCDC)
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

func (cdc *RabinCDC) Next() ([]byte, error) {
	var err error
	var n int
	for {
		if cdc.processBytes(cdc.rest) {
			break
		}
		cdc.rest = make([]byte, 0)
		n, err = cdc.file.Read(cdc.readBuf)
		if err != nil && err != io.EOF {
			break
		}
		if cdc.processBytes(cdc.readBuf[:n]) {
			break
		}
		if err == io.EOF {
			break
		}
	}
	c := make([]byte, len(cdc.chunk))
	copy(c, cdc.chunk[:len(cdc.chunk)])
	cdc.reset()
	if err != nil && err != io.EOF {
		return nil, err
	}
	if len(cdc.rest) == 0 && err == io.EOF {
		return c, err
	}
	return c, nil
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
