package fastcdc

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

type File struct {
	Hash string
	by   []byte
	i    uint64
}

func FastCDC(chunk *storage.Chunk, fl *storage.File, input string, chunk2 *storage.ChunkStore) error {
	avg := 8 * 1024
	minByte, maxByte := avg/4, avg*4
	l, r := (1<<13)-1, (1<<11)-1
	gearHash := NewGearHash()
	ch := make(chan File, 100)
	wg := &sync.WaitGroup{}
	var sd uint64
	go ForChen(fl, chunk2, &ch)
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
	u := make([]byte, 0, maxByte)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				sd += 1
				go AddChunk(&ch, u, sd, wg)
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
				sd += 1
				go AddChunk(&ch, u, sd, wg)
				gearHash.l = 0
			}
			if len(u) > maxByte {
				sd += 1
				go AddChunk(&ch, u, sd, wg)
				gearHash.l = 0
			}
		}
	}
	wg.Wait()
	close(ch)
	chunk.FileList = append(chunk.FileList, *fl)
	return nil
}

func AddChunk(ch *chan File, u []byte, c uint64, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	hash := sha256.Sum256(u)
	str := string(hash[:])
	fi := File{Hash: str, by: u, i: c}
	*ch <- fi
}

func ForChen(fl *storage.File, chunkStore *storage.ChunkStore, ch *chan File) {
	for input := range *ch {
		if _, ok := (*chunkStore)[input.Hash]; !ok {
			(*chunkStore)[input.Hash] = input.by
		}
		fl.Hash[input.i] = input.Hash
	}
}
