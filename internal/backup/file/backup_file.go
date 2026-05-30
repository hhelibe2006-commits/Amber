package file

import (
	"github.com/hhelibe2006-commits/Amber/internal/backup/file/fastcdc"
	"github.com/hhelibe2006-commits/Amber/internal/backup/file/writer"
	"github.com/hhelibe2006-commits/Amber/internal/value"
)

func BackupFile(input []string, output string) error {
	tempFiles := value.NewTempFiles()
	defer tempFiles.Close()
	defer tempFiles.Remove()
	ch := make(chan fastcdc.Info, 100)
	go writer.Writer()

	for _, file := range input {
		err := fastcdc.FastCDC(file, ch)
		if err != nil {
			return err
		}
	}
	return nil
}
