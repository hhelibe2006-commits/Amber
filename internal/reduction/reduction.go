package reduction

import (
	"errors"
	"fmt"
	"os"

	"github.com/hhelibe2006-commits/Amber/internal/reduction/file"
)

func Run(input string) error {
	f, err := os.Open(input)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(f)
	if IsDir(f) {
		return errors.New(fmt.Sprint(input + "不是文件"))
	}
	typ, err := TypeAnalysis(f)
	if err != nil {
		return err
	}
	switch typ {
	case "file":
		err = file.ReductionFile(f)
	default:
		fmt.Println("这不对吧")
	}
	return nil
}
