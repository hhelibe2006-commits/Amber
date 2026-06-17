package reduction

import "os"

func IsDir(file *os.File) bool {
	info, err := file.Stat()
	if err != nil {
		return false
	}
	return info.IsDir()
}
