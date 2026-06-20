package cdc

import (
	"io"
	"os"
)

type hashCDC interface {
	processBytes(data []byte) bool
	reset()
}

// HashCDC 为HashCDC算法的基类，请勿实例化
type HashCDC struct {
	file *os.File

	avgBlockSize uint64
	maxBlockSize uint64
	minBlockSize uint64

	chunk   []byte
	readBuf []byte
	rest    []byte

	proc hashCDC
}

func (cdc *HashCDC) Next() ([]byte, error) {
	var err error
	var n int
	for {
		if cdc.proc.processBytes(cdc.rest) {
			break
		}
		cdc.rest = make([]byte, 0)
		n, err = cdc.file.Read(cdc.readBuf)
		if err != nil && err != io.EOF {
			break
		}
		if cdc.proc.processBytes(cdc.readBuf[:n]) {
			break
		}
		if err == io.EOF {
			break
		}
	}
	c := make([]byte, len(cdc.chunk))
	copy(c, cdc.chunk[:len(cdc.chunk)])
	cdc.proc.reset()
	if err != nil && err != io.EOF {
		return nil, err
	}
	if len(cdc.rest) == 0 && err == io.EOF {
		return c, err
	}
	return c, nil
}
