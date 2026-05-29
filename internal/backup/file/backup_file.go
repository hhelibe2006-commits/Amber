package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file/fastcdc"
)

func BackupFile(input []string, output string) error {
	//用于临时存放备份数据的目录
	tempDir, err := os.MkdirTemp("", "amber-*")
	if err != nil {
		return err
	}
	defer func(path string) {
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println(err)
		}
	}(tempDir)
	//用于存放数据切片的临时文件
	datePath := filepath.Join(tempDir, "date")
	tempDate, err := os.Create(datePath)
	if err != nil {
		return err
	}
	defer func() {
		err := tempDate.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//用于存放哈希对应位置的临时文件
	hashPath := filepath.Join(tempDir, "hash")
	tempHash, err := os.Create(hashPath)
	if err != nil {
		return err
	}
	defer func() {
		err := tempHash.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	//用于存放文件信息的临时文件
	filePath := filepath.Join(tempDir, "file")
	tempFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		err := tempFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	ch := make(chan fastcdc.Info, 100)
	for _, file := range input {
		fastcdc.FastCDC(file, ch)
	}
	return nil
}
