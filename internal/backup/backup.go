package backup

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/cli"
	"github.com/hhelibe2006-commits/Amber/internal/storage"
)

func Run(typ string, input cli.Put, output string) error {
	switch typ {
	case "file":
		fmt.Println("校验路径中")
		err := judge(input, output)
		if err != nil {
			return err
		}
		fmt.Println("路径转化文件")
		inputList := pathToFile(input)
		fmt.Println("备份中")
		if err = backupFile(inputList, output); err != nil {
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
				fmt.Println(path, "已完成")
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
	var size uint64
	for i := range input {
		info, err := os.Stat(input[i])
		if err != nil {
			return err
		}
		size += uint64(info.Size())
	}
	chunk := storage.NewChunk()
	err := processFile(input, output, chunk)
	if err != nil {
		return err
	}
	return nil
}
