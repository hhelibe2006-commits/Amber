package backup

import (
	"compress/gzip"
	"encoding/gob"
	"os"
	"path/filepath"
)

func processFile(input []string, output string, progress chan *uint64) error {
	dateFile := filepath.Join(output, "date")
	snapshotFile := filepath.Join(output, "snapshot")
	println(dateFile, snapshotFile)

	if err := os.Mkdir(output, 0755); err != nil {
		return err
	}
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			return
		}
	}()
	for _, line := range input {
		println(line)
	}
	writer := gzip.NewWriter(file)
	encoder := gob.NewEncoder(writer)
	println(encoder.Encode(input))
	return nil
}
