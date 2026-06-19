package cdc

import (
	"io"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/cdc/hash"
)

func init() {
	Add("fastcdc", NewFastCdc)
}

type FastCdc struct {
	gearHash *hash.GearHash
	file     *os.File

	avgBlockSize uint64
	maxBlockSize uint64
	minBlockSize uint64

	chunk   []byte
	readBuf []byte
	rest    []byte
}

func NewFastCdc(file *os.File) CDC {
	cdc := new(FastCdc)
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
	var t int
	for i, b := range data {
		cdc.chunk = append(cdc.chunk, b)
		cdc.gearHash.Next(b)
		if uint64(len(cdc.chunk)) < cdc.minBlockSize {
			continue
		}
		if uint64(i) > cdc.avgBlockSize {
			t = 2
		}
		if cdc.gearHash.Check((cdc.avgBlockSize-1)>>t) || uint64(len(cdc.chunk)) == cdc.maxBlockSize {
			cdc.rest = append([]byte{}, data[i+1:]...)
			return true
		}
	}
	return false
}

func (cdc *FastCdc) reset() {
	cdc.gearHash.Reset()
	cdc.chunk = cdc.chunk[:0]
}
