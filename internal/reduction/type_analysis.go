package reduction

import (
	"archive/tar"
	"errors"
	"io"
	"os"
)

func TypeAnalysis(inFile *os.File) (string, error) {
	typeName := "file"
	reader := tar.NewReader(inFile)
	for {
		header, err := reader.Next()
		if err == io.EOF {
			return "", errors.New("这对吗？")
		}
		if err != nil {
			return "", err
		}
		switch header.Name {
		case "file":
			typeName = header.Name
			return typeName, nil
		default:
			break
		}
	}
}
