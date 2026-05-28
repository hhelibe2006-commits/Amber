package backup

import (
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/backup/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/file"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func processFile(input []string, output string, progress chan *uint64) error {
	dateFile := filepath.Join(output, "date")
	snapshotFile := filepath.Join(output, "snapshot")
	ch := make(chan *storage.ChunkStore)
	go oWrite(ch, dateFile)

	if err := os.Mkdir(output, 0755); err != nil {
		return err
	}
	DFile, err := os.Create(dateFile)
	if err != nil {
		return err
	}
	SFile, err := os.Create(snapshotFile)
	defer func() {
		err := DFile.Close()
		if err != nil {
			return
		}
	}()
	defer func() {
		err := SFile.Close()
		if err != nil {
			return
		}
	}()
	for _, line := range input {
		err := fastcdc.FastCDC(line, ch)
		if err != nil {
			return err
		}
	}
	return nil
}

func oWrite(ch chan *storage.ChunkStore, outfile string) {
	f, err := os.Create(outfile)
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	for c := range ch {
		writeFile, err := file.WriteFile(f, c.Buf)
		if err != nil {
			return
		}
		println(writeFile)
	}
}
