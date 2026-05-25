package backup

import (
	"encoding/gob"
	"errors"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/backup/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func processFile(input []string, output string, chunk *storage.Manifest) error {
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
	for _, fileName := range input {
		err = filepath.WalkDir(fileName, func(path string, d fs.DirEntry, err error) error {
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
				err = fastcdc.FastCDC(chunk, path, &chunkStore)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return err
		}
	}
	if file, err = os.Create(output); err != nil {
		return err
	}
	encoder := gob.NewEncoder(file)
	if err = encoder.Encode(chunk); err != nil {
		return err
	}
	for key, value := range chunkStore {
		if err := encoder.Encode([2]interface{}{key, value}); err != nil {
			return err
		}
	}
	return nil
}
