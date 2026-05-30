package backup

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file"
)

func Run(typ string, input []string, output string) error {
	switch typ {
	case "file":
		fmt.Println("校验路径中")
		if err := file.PathValidation(input, output); err != nil {
			fmt.Println("校验错误:", err)
			return err
		}
		fmt.Println("获取文件中")
		fileList, err := file.ListDir(input)
		if err != nil {
			fmt.Println("获取文件错误:", err)
			return err
		}
		err = file.BackupFile(fileList, output)
		if err != nil {
			fmt.Println("备份出错:", err)
			return err
		}
	}
	return nil
}
