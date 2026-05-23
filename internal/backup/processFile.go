package backup

import (
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
		if info.IsDir() {
			return nil
		}
		err = fastcdc.FastCDC(chunk, fl, input)
		if err != nil {
			return err
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				return
			}
		}(file)
		encoder := json.NewEncoder(file)
		err = encoder.Encode(chunk)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}
