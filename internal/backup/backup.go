package backup

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hhelibe2006-commits/Amber/internal/cli"
)

func Run(typ string, input cli.Put, output string) error {
	err := judge(input, output)
	if err != nil {
		return err
	}
	switch typ {
	case "file":
		backupFile(input, output)
	case "system":
		break
	case "disk":
		break
	default:
		return errors.New(fmt.Sprintf("不支持的备份类型%s", typ))
	}
	return nil
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

func backupFile(input cli.Put, output string) {
	for i := 0; i < len(input); i++ {
		inPath := filepath.Clean(input[i])
		if info, err := os.Stat(inPath); err != nil {
			return
		} else if info.IsDir() {

		} else {

		}
	}
}
