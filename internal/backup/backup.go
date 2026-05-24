package backup

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sync"

	"github.com/hhelibe2006-commits/Amber/internal/cli"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func Run(typ string, input cli.Put, output string) error {
	switch typ {
	case "file":
		err := judge(input, output)
		if err != nil {
			return err
		}
		inputList := pathToFile(input)
		err = backupFile(inputList, output)
		if err != nil {
			return err
		}
	case "system":
		break
	case "disk":
		break
	default:
		return errors.New(fmt.Sprintf("不支持的备份类型%s", typ))
	}
	return nil
}

func pathToFile(path cli.Put) cli.Put {
	fileList := make(cli.Put, 0, len(path))
	for _, file := range path {
		err := filepath.WalkDir(file, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !d.IsDir() {
				fileList = append(fileList, path)
			}
			return nil
		})
		if err != nil {
			panic(err)
		}
	}
	return fileList
}

func judge(input cli.Put, output string) error {
	for i := 0; i < len(input); i++ {
		inPath := filepath.Clean(input[i])
		if _, err := os.Stat(inPath); os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("不存在该文件或目录%s", input[i]))
		}
	}
	outPath := filepath.Dir(filepath.Clean(output))
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		return errors.New(fmt.Sprintf("不存在该文件或目录%s", output))
	}
	return nil
}

func backupFile(input cli.Put, output string) error {
	var wg sync.WaitGroup
	var size uint64
	for i := range input {
		info, err := os.Stat(input[i])
		if err != nil {
			return err
		}
		size += uint64(info.Size())
	}
	errCh := make(chan error, len(input))
	sem := make(chan struct{}, 8)
	chunk := storage.NewChunk()
	for _, file := range input {
		wg.Add(1)
		sem <- struct{}{}
		go func(p string) {
			defer wg.Done()
			defer func() { <-sem }()
			fl := storage.NewFile()
			err := processFile(file, output, chunk, fl)
			if err != nil {
				errCh <- err
			}
		}(file)
	}
	if len(errCh) > 0 {
		err := <-errCh
		return err
	}
	wg.Wait()
	return nil
}
