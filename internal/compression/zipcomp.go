package compression

import (
	"archive/zip"
	"io"
	"os"
)

func AddFileToZip(output string) error {
	zipFile, err := os.Create(output + ".zip")
	if err != nil {
		return err
	}
	defer func(zipFile *os.File) {
		err := zipFile.Close()
		if err != nil {

		}
	}(zipFile)
	zipWriter := zip.NewWriter(zipFile)
	defer func(zipWriter *zip.Writer) {
		err := zipWriter.Close()
		if err != nil {

		}
	}(zipWriter)
	fileToZip, err := os.Open(output)
	if err != nil {
		return err
	}
	defer func(fileToZip *os.File) {
		err := fileToZip.Close()
		if err != nil {
			return
		}
	}(fileToZip)
	writer, err := zipWriter.Create(output + ".zip")
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	if err != nil {
		return err
	}
	return nil
}
