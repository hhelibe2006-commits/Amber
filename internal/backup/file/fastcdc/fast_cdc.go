package fastcdc

import (
	"crypto/sha256"
	"io"
	"os"
)

func FastCDC(path string, ch chan Info) error {
	g := NewGearHash()
	g.value = 0
	u := make([]byte, 0)
	by := make([]byte, 1024)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	for {
		_, err := file.Read(by)
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			info := new(Info)
			info.Path = path
			hash := sha256.Sum256(u)
			info.Hash = string(hash[:])
			info.Value = u
			ch <- *info
			u = u[:0]
			break
		}
		for _, b := range by {
			g.Next(b)
			u = append(u, b)
			if g.value&3 == 0 {
				info := new(Info)
				info.Path = path
				hash := sha256.Sum256(u)
				info.Hash = string(hash[:])
				info.Value = u
				ch <- *info
				u = u[:0]
			}
		}
	}
	return nil
}
