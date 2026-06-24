package file

import (
	"github.com/hhelibe2006-commits/Amber/internal/openfile"
)

func ListDir(path []string) (openfile.OpenFile, error) {
	openFile := openfile.NewOpenFile(len(path))
	if err := openFile.Adds(path); err != nil {
		return openFile, err
	}
	return openFile, nil
}
