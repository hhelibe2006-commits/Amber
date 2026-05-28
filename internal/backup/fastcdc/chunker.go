package fastcdc

import (
	"crypto/sha256"
	"io"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

// BlockMeta 用于文件切片的哈希值、字节切片以及在文件的位置
type BlockMeta struct {
	Hash  string
	bytes []byte
	Index uint64
}

// FastCDC 用于将单个文件切分成多个块，并将每个块的哈希值存储在chunkStore中，同时将文件的元数据存储在Manifest中
func FastCDC(input string, ch chan *storage.ChunkStore) error {
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	var l, r uint64 = (1 << 13) - 1, (1 << 11) - 1
	u := make([]byte, 0, maxByte)
	d := make([]byte, 1024)
	gHash := NewGearHash()
	var p uint64
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer func() {
		err = f.Close()
		if err != nil {
			return
		}
	}()
	for {
		_, err := f.Read(d)
		if err != nil && err != io.EOF {
			return err
		} else if err == io.EOF {
			if len(u) != 0 {
				hash := sha256.Sum256(u)
				ma := storage.NewChunkStore()
				ma.Name = input
				ma.Buf = u
				ma.Hash = string(hash[:])
				ch <- ma
				u = u[:0]
			}
			break
		}
		for _, i := range d {
			gHash.Next(i)
			u = append(u, i)
			if len(u) < minByte {
				continue
			}
			if len(u) > avg {
				p = r
			} else {
				p = l
			}
			if p&gHash.hashValue == 0 || len(u) > maxByte {
				hash := sha256.Sum256(u)
				ma := storage.NewChunkStore()
				ma.Buf = u
				ma.Hash = string(hash[:])
				ma.Name = input
				ch <- ma
				u = u[:0]
			}
		}
	}
	return nil
}
