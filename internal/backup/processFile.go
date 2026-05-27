package backup

import (
	"compress/gzip"
	"encoding/gob"
	"os"
)

func processFile(input []string, output string) error {
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
	writer := gzip.NewWriter(file)
	encoder := gob.NewEncoder(writer)
	println(encoder.Encode(input))
	return nil
}
