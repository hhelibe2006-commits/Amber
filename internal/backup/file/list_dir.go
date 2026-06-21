package file

import (
	"github.com/hhelibe2006-commits/Amber/internal/value/openfile"
)

func ListDir(path []string) (openfile.OpenFile, error) {
	openFile := openfile.NewOpenFile(len(path))
	if err := openFile.Adds(path); err != nil {
		return nil, err
	}
	return openFile, nil
}
