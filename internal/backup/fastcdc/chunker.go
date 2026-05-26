package fastcdc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

// BlockMeta 用于文件切片的哈希值、字节切片以及在文件的位置
type BlockMeta struct {
	Hash  string
	bytes []byte
	Index uint64
}

// FastCDC 用于将单个文件切分成多个块，并将每个块的哈希值存储在chunkStore中，同时将文件的元数据存储在Manifest中
func FastCDC(chunk *storage.Manifest, input string, chunk2 *storage.ChunkStore) error {
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	l, r := (1<<13)-1, (1<<11)-1
	fl := storage.NewFileMeta()
	gearHash := NewGearHash()
	ch := make(chan *BlockMeta, 100)
	wg := &sync.WaitGroup{}
	var seq uint64
	go consumeChunks(fl, chunk2, ch)
	if err := fl.Set(input); err != nil {
		return err
	}
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
	buffer := make([]byte, 0, maxByte)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				seq += 1
				wg.Add(1)
				go addChunk(ch, buffer, seq, wg)
				buffer = buffer[:0]
				break
			}
			return err
		}
		for i := 0; i < n; i++ {
			buffer = append(buffer, buf[i])
			gearHash.Next(buf[i])
			if len(buffer) < minByte {
				continue
			}
			var c uint64
			if len(buffer) < avg {
				c = uint64(l)
			} else {
				c = uint64(r)
			}
			if gearHash.hashValue&c == 0 {
				seq += 1
				wg.Add(1)
				go addChunk(ch, buffer, seq, wg)
				gearHash.hashValue = 0
				buffer = buffer[:0]
			}
			if len(buffer) > maxByte {
				seq += 1
				wg.Add(1)
				go addChunk(ch, buffer, seq, wg)
				gearHash.hashValue = 0
				buffer = buffer[:0]
			}
		}
	}
	wg.Wait()
	close(ch)
	chunk.FileList = append(chunk.FileList, *fl)
	return nil
}

func addChunk(ch chan *BlockMeta, u []byte, c uint64, wg *sync.WaitGroup) {
	defer wg.Done()
	hash := sha256.Sum256(u)
	str := string(hash[:])
	fi := BlockMeta{Hash: str, bytes: u, Index: c}
	ch <- &fi
}

func consumeChunks(fl *storage.FileMeta, chunkStore *storage.ChunkStore, ch chan *BlockMeta) {
	for input := range ch {
		if _, ok := (*chunkStore)[input.Hash]; !ok {
			(*chunkStore)[input.Hash] = input.bytes
		}
		fl.Hash[input.Index] = input.Hash
	}
}
