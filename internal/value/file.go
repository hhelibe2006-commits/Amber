package value

import (
	"os"
	"time"
)

type File struct {
	Path     string
	HashList []string
	Mode     os.FileMode
	ModTime  time.Time
}
