package cdc

import (
	"io"
	"os"
)

func init() {
	Add("rabincdc", NewRabinCDC)
}

type RabinCDC struct {
	Window      []byte
	P           uint64
	M           uint64
	Avg         uint64
	MaxSize     uint64
	MinSize     uint64
	FingerPrint uint64
	BlockLen    uint64
	File        *os.File
	u           []byte
	by          []byte
}

func NewRabinCDC(file *os.File) Cdc {
	rabinCDC := new(RabinCDC)
	rabinCDC.Window = make([]byte, 64)
	rabinCDC.P = 1007
	rabinCDC.M = 1<<64 - 59
	rabinCDC.Avg = 12
	rabinCDC.File = file
	rabinCDC.u = make([]byte, 0, rabinCDC.MaxSize)
	rabinCDC.by = make([]byte, rabinCDC.Avg)
	return rabinCDC
}

func (r *RabinCDC) Next() ([]byte, error) {
	for {
		n, err := r.File.Read(r.by)
		if err != nil && err != io.EOF {
			return nil, err
		} else if n == 0 {
			break
		} else if err == io.EOF {
			return append(r.u, r.by...), nil
		}

	}
	u := r.u[:len(r.u)]
	r.u = r.u[:0]
	return u, nil
}
