package value

import (
	"fmt"
	"os"
	"path/filepath"
)

type TempFiles struct {
	tempDir  string
	TempInfo *os.File //备份信息
	TempDate *os.File //数据的存放位置
	TempHash *os.File //哈希与位置的对应关系
	TempFile *os.File //文件相关信息
}

func NewTempFiles() (*TempFiles, error) {
	tempFiles := new(TempFiles)
	var err error
	tempFiles.tempDir, err = os.MkdirTemp("", "*")
	if err != nil {
		panic(err)
	}

	infoPath := filepath.Join(tempFiles.tempDir, "info")
	tempFiles.TempInfo, err = os.Create(infoPath)
	if err != nil {
		return nil, err
	}

	datePath := filepath.Join(tempFiles.tempDir, "date")
	tempFiles.TempDate, err = os.Create(datePath)
	if err != nil {
		return nil, err
	}

	hashPath := filepath.Join(tempFiles.tempDir, "hash")
	tempFiles.TempHash, err = os.Create(hashPath)
	if err != nil {
		return nil, err
	}

	filePath := filepath.Join(tempFiles.tempDir, "file")
	tempFiles.TempFile, err = os.Create(filePath)
	if err != nil {
		return nil, err
	}
	return tempFiles, nil
}

func (tempFiles *TempFiles) Close() {
	if err := tempFiles.TempDate.Close(); err != nil {
		fmt.Print("date", err)
	}
	if err := tempFiles.TempHash.Close(); err != nil {
		fmt.Print("hash", err)
	}
	if err := tempFiles.TempFile.Close(); err != nil {
		fmt.Print("file", err)
	}
}

func (tempFiles *TempFiles) Remove() {
	if err := os.RemoveAll(tempFiles.tempDir); err != nil {
		fmt.Print("remove", err)
	}
}

func (tempFiles *TempFiles) FileList() []string {
	fileList := make([]string, 0, 4)
	fileList = append(fileList, tempFiles.TempInfo.Name())
	fileList = append(fileList, tempFiles.TempFile.Name())
	fileList = append(fileList, tempFiles.TempDate.Name())
	fileList = append(fileList, tempFiles.TempHash.Name())
	return fileList
}
