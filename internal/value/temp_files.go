package value

import (
	"os"
	"path/filepath"
)

type TempFiles struct {
	tempDir  string
	TempDate *os.File //数据的存放位置
	TempHash *os.File //哈希与位置的对应关系
	TempFile *os.File //文件相关信息
}

func NewTempFiles() *TempFiles {
	tempFiles := new(TempFiles)
	var err error
	tempFiles.tempDir, err = os.MkdirTemp("", "*")
	if err != nil {
		panic(err)
	}

	datePath := filepath.Join(tempFiles.tempDir, "date")
	tempFiles.TempDate, err = os.Create(datePath)
	if err != nil {
		panic(err)
	}

	hashPath := filepath.Join(tempFiles.tempDir, "hash")
	tempFiles.TempHash, err = os.Create(hashPath)
	if err != nil {
		panic(err)
	}

	filePath := filepath.Join(tempFiles.tempDir, "file")
	tempFiles.TempFile, err = os.Create(filePath)
	if err != nil {
		panic(err)
	}
	return tempFiles
}

func (tempFiles *TempFiles) Close() {
	if err := tempFiles.TempDate.Close(); err != nil {
		return
	}
	if err := tempFiles.TempHash.Close(); err != nil {
		return
	}
	if err := tempFiles.TempFile.Close(); err != nil {
		return
	}
}

func (tempFiles *TempFiles) Remove() {
	if err := os.RemoveAll(tempFiles.tempDir); err != nil {
		return
	}
}
