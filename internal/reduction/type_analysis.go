package reduction

import (
	"archive/zip"
	"encoding/gob"
	"errors"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func TypeAnalysis(infile *os.File) (string, error) {
	info, err := infile.Stat()
	if err != nil {
		return "", err
	}
	zipReader, err := zip.NewReader(infile, info.Size())
	if err != nil {
		return "", err
	}
	file, err := value.NewTempFiles()
	if err != nil {
		return "", err
	}
	name := file.TempInfo.Name()
	name = filepath.Base(name)
	for _, f := range zipReader.File {
		if f.Name == name {
			re, err := f.Open()
			if err != nil {
				return "", err
			}
			decoder := gob.NewDecoder(re)
			var red string
			err = decoder.Decode(&red)
			if err != nil {
				return "", err
			}
			return red, nil
		}
	}
	return "", errors.New("没找到文件")
}
