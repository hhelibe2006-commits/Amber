package backup

import (
	"encoding/json"
	"errors"
	"fmt"
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
		if !info.IsDir() {
			err = fastcdc.FastCDC(chunk, fl, path)
			if err != nil {
				return err
			}
		}
		return nil
	})

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	jsonData, err := json.MarshalIndent(chunk, "", "  ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(output, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}
