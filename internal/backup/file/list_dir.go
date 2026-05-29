package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func ListDir(path []string) ([]string, error) {
	fileList := make([]string, len(path))
	for _, file := range path {
		err := filepath.Walk(file, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return nil
			}
			fmt.Println(path)
			fileList = append(fileList, path)
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	return fileList, nil
}
