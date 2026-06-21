package backup

import (
	"fmt"

	"github.com/hhelibe2006-commits/Amber/internal/backup/file"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func Run(clas *value.Clas) error {
	switch clas.Typ {
	case "file":
		fmt.Println("校验路径中")
		if err := file.PathValidation(clas.Input, clas.Output); err != nil {
			fmt.Println("校验错误:", err)
			return err
		}
		fmt.Println("获取文件中")
		openFile, err := file.ListDir(clas.Input)
		if err != nil {
			fmt.Println("获取文件错误:", err)
			return err
		}
		err = file.BackupFile(clas, openFile)
		if err != nil {
			fmt.Println("备份出错:", err)
			return err
		}
	default:
		return fmt.Errorf("无此备份类型")
	}
	return nil
}
