package reduction

import (
	"fmt"
	"os"
)

func Run(input string) error {
	file, err := os.Open(input)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)
	if err != nil {
		return err
	}
	typ, err := TypeAnalysis(file)
	if err != nil {
		return err
	}
	println(typ)
	return nil
}
