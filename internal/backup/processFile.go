package backup

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

func processFile(input []string, output string, progress chan uint64) error {
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
