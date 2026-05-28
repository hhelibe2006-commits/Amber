package backup

import (
	"archive/zip"
	"encoding/gob"
	"io"
	"os"
	"path/filepath"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/backup/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/file"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func processFile(input []string, output string, progress chan *uint64) error {
	tmp, err := os.MkdirTemp("", "amber-*")
	if err != nil {
		return err
	}
	defer func() {
		if err := os.RemoveAll(tmp); err != nil {
			panic(err)
		}
	}()
	wg := new(sync.WaitGroup)
	dateFile := filepath.Join(tmp, "date")
	snapshotFile := filepath.Join(tmp, "snapshot")
	ch := make(chan *storage.ChunkStore)
	DFile, err := os.Create(dateFile)
	if err != nil {
		return err
	}
	defer func() {
		err := DFile.Close()
		if err != nil {
			return
		}
	}()
	SFile, err := os.Create(snapshotFile)
	if err != nil {
		return err
	}
	defer func() {
		err := SFile.Close()
		if err != nil {
			return
		}
	}()
	m := storage.NewManifest()
	wg.Add(1)
	go oWrite(ch, DFile, m, wg)
	for _, line := range input {
		err := fastcdc.FastCDC(line, ch)
		if err != nil {
			return err
		}
	}
	close(ch)
	w := gob.NewEncoder(SFile)
	err = w.Encode(m)
	if err != nil {
		return err
	}
	f, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			return
		}
	}(f)
	write := zip.NewWriter(f)
	defer func(write *zip.Writer) {
		err := write.Close()
		if err != nil {
			return
		}
	}(write)
	_, err = write.Create(dateFile)
	if err != nil {
		return err
	}
	_, err = write.Create(snapshotFile)
	if err != nil {
		return err
	}
	wg.Wait()
	return nil
}

func oWrite(ch chan *storage.ChunkStore, outfile *os.File, m *storage.Manifest, wg *sync.WaitGroup) {
	for c := range ch {
		if _, ok := m.FileNameList[c.Name]; !ok {
			s := storage.NewFileMeta()
			f, err := os.Stat(c.Name)
			if err != nil {
				return
			}
			s.FilePath = c.Name
			s.Mode = f.Mode()
			s.ModeTime = f.ModTime()
			m.FileNameList[c.Name] = s
		}
		writeFile, err := file.WriteFile(outfile, c.Buf)
		if err != nil {
			return
		}
		ret, err := outfile.Seek(0, io.SeekCurrent)
		if err != nil {
			return
		}
		ret = ret - writeFile
		m.FileNameList[c.Name].Lengths = append(m.FileNameList[c.Name].Lengths, writeFile)
		m.FileNameList[c.Name].Locations = append(m.FileNameList[c.Name].Locations, ret)
	}
	wg.Done()
}
