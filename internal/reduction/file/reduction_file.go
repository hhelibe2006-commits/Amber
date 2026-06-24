package file

import (
	"archive/zip"
	"encoding/gob"
	"io"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/reduction/file/read"
	"github.com/hhelibe2006-commits/Amber/internal/value"
	"github.com/hhelibe2006-commits/Amber/internal/value/chunkindex"
)

func ReductionFile(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}
	zipReader, err := zip.NewReader(file, info.Size())
	if err != nil {
		return err
	}
	temp, err := value.NewTempFiles()
	if err != nil {
		return err
	}
	var chunkIndex chunkindex.ChunkIndex
	fr, err := value.NewTempFiles()
	if err != nil {
		return err
	}
	defer fr.Close()
	for _, f := range zipReader.File {
		if f.Name == filepath.Base(temp.TempDate.Name()) {
			file, err := f.Open()
			if err != nil {
				return err
			}
			_, err = io.Copy(fr.TempDate, file)
			if err != nil {
				return err
			}
		} else if f.Name == filepath.Base(temp.TempFile.Name()) {
			file, err := f.Open()
			if err != nil {
				return err
			}
			decoder := gob.NewDecoder(file)
			err = decoder.Decode(&chunkIndex.FileMap)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		} else if f.Name == filepath.Base(temp.TempHash.Name()) {
			file, err := f.Open()
			if err != nil {
				return err
			}
			decoder := gob.NewDecoder(file)
			err = decoder.Decode(&chunkIndex.HashToPlace)
			if err != nil {
				return err
			}
		}
	}
	for path, hash := range chunkIndex.FileMap {
		err := re(path, hash, fr, chunkIndex)
		if err != nil {
			return err
		}
		if err := os.Chmod(path, hash.Mode); err != nil {
			return err
		}
		if err := os.Chtimes(path, hash.ModTime, hash.ModTime); err != nil {
			return err
		}
	}
	return nil
}

func re(path string, hash *value.File, fr *value.TempFiles, chunkIndex chunkindex.ChunkIndex) error {
	create, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(create *os.File) {
		err = create.Close()
		if err != nil {
			return
		}
	}(create)
	for _, ha := range hash.HashList {
		a := chunkIndex.HashToPlace[ha]
		_, err = fr.TempDate.Seek(a.A, io.SeekStart)
		if err != nil {
			return err
		}
		s := make([]byte, a.B)
		_, err = io.ReadFull(fr.TempDate, s)
		if err != nil {
			return err
		}
		u, err := read.GzipDecompress(s)
		if err != nil {
			return err
		}
		_, err = create.Write(u)
	}
	return nil
}
