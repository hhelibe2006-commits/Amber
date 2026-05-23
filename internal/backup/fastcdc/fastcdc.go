package fastcdc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func FastCDC(chunk *storage.Chunk, fl *storage.File, input string) error {
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	l, r := (1<<13)-1, (1<<11)-1
	gearHash := NewGearHash()
	info, err := os.Stat(input)
	if err != nil {
		return err
	}
	fl.Mode = info.Mode()
	fl.ModeTime = info.ModTime()
	abs, err := filepath.Abs(input)
	if err != nil {
		return err
	}
	fl.FilePath = abs
	file, err := os.Open(input)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	buf := make([]byte, 1024)
	u := make([]byte, 0, maxByte)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				hash := sha256.Sum256(u)
				fl.Hash = append(fl.Hash, hash)
				if _, ok := chunk.Chunk[hash]; !ok {
					chunk.Chunk[hash] = u
				}
				break
			}
			return err
		}
		for i := 0; i < n; i++ {
			u = append(u, buf[i])
			gearHash.Next(buf[i])
			if len(u) < minByte {
				continue
			}
			var c uint64
			if len(u) < avg {
				c = uint64(l)
			} else {
				c = uint64(r)
			}
			if gearHash.l&c == 0 {
				hash := sha256.Sum256(u)
				fl.Hash = append(fl.Hash, hash)
				if _, ok := chunk.Chunk[hash]; !ok {
					chunk.Chunk[hash] = u
				}
				gearHash.l = 0
			}
			if len(u) > maxByte {
				hash := sha256.Sum256(u)
				fl.Hash = append(fl.Hash, hash)
				if _, ok := chunk.Chunk[hash]; !ok {
					chunk.Chunk[hash] = u
				}
				gearHash.l = 0
			}
		}
	}
	chunk.FileList = append(chunk.FileList, *fl)
	return nil
}
