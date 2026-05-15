package backup

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/src/internal/args"
	"github.com/hhelibe2006-commits/Amber/src/internal/chunker"
)

func Backup(cfg *args.Config) {
	for i := 0; i < len(cfg.InputPath); i++ {
		fileInfo, err := os.Stat(cfg.InputPath[i])
		if err != nil {
			if os.IsNotExist(err) {
				fmt.Printf("%s不存在\n", cfg.InputPath[i])
			} else {
				fmt.Println("发生错误：", err)
			}
			return
		}
		if fileInfo.IsDir() {
			err := filepath.WalkDir(
				cfg.InputPath[i],
				func(path string, d fs.DirEntry, err error) error {
					if err != nil {
						return err
					}
					if !d.IsDir() {
						data, err := os.ReadFile(path)
						if err != nil {
							fmt.Println("读取文件失败", err)
							return err
						}
						chunker.Chunker(data)
					}
					return nil
				})
			if err != nil {
				fmt.Printf("便利出错%v\n", err)
			}
		} else {
			data, err := os.ReadFile(cfg.InputPath[i])
			if err != nil {
				fmt.Println("读取文件失败", err)
			}
			chunker.Chunker(data)
		}
	}
}
