package backup

import (
	"encoding/gob"
	"encoding/json"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/backup/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func processFile(input string, output string, chunk *storage.Chunk, fl *storage.File) error {
	file, err := os.Open(output)
	if err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	if file != nil {
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		return errors.New("输出路径已经存在")
	}
	chunkStore := storage.NewChunkStore()
	err = filepath.WalkDir(input, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		file, err = os.Create(output)
		if err != nil {
			return err
		}
		func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		info, err := os.Stat(path)
		if err != nil {
			return err
		}
		if !info.IsDir() {
			err = fastcdc.FastCDC(chunk, fl, path, chunkStore)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	file, err = os.Create(output)
	if err != nil {
		return err
	}
	ender := json.NewEncoder(file)
	err = ender.Encode(chunk)
	if err != nil {
		return err
	}
	filef := filepath.Dir(output)
	f, err := filepath.Abs(filef)
	if err != nil {
		return err
	}
	f = f + "/data.gob"
	file, err = os.Create(f)
	encoder := gob.NewEncoder(file)
	err = encoder.Encode(chunkStore)
	if err != nil {
		return err
	}
	return nil
}
