package cdc

import (
	"io"
	"os"
)

func init() {
	Add("fastcdc", NewFastCdc)
}

type FastCdc struct {
	gearHash *GearHash
	file     *os.File

	avgBlockSize uint64
	maxBlockSize uint64
	minBlockSize uint64

	chunk   []byte
	readBuf []byte
	rest    []byte
}

func NewFastCdc(file *os.File) Cdc {
	cdc := new(FastCdc)
	cdc.gearHash = NewGearHash()
	cdc.file = file

	cdc.avgBlockSize = uint64(1 << 13)
	cdc.maxBlockSize = cdc.avgBlockSize * 4
	cdc.minBlockSize = cdc.avgBlockSize / 4

	cdc.chunk = make([]byte, 0, cdc.maxBlockSize)
	cdc.readBuf = make([]byte, cdc.avgBlockSize)
	cdc.rest = make([]byte, 0)
	return cdc
}

func (cdc *FastCdc) Next() ([]byte, error) {
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

func (cdc *FastCdc) processBytes(data []byte) bool {
	for i, b := range data {
		cdc.chunk = append(cdc.chunk, b)
		cdc.gearHash.next(b)
		if uint64(len(cdc.chunk)) < cdc.minBlockSize {
			continue
		}
		if cdc.gearHash.value&(cdc.avgBlockSize-1) == 0 || uint64(len(cdc.chunk)) == cdc.maxBlockSize {
			cdc.rest = append([]byte(nil), data[i+1:]...)
			return true
		}
	}
	return false
}

func (cdc *FastCdc) reset() {
	cdc.gearHash.Reset()
	cdc.chunk = cdc.chunk[:0]
}
