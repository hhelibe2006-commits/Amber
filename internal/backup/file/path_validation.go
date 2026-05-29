package file

import (
	"os"
	"path/filepath"
)

func PathValidation(input []string, output string) error {
	if _, err := os.Stat(filepath.Dir(output)); err != nil {
		return err
	}
	for _, v := range input {
		if _, err := os.Stat(v); err != nil {
			return err
		}
	}
	return nil
}
