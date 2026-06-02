package value

import (
	"os"
	"time"
)

type File struct {
	Path     string
	HashList [][32]byte
	Mode     os.FileMode
	ModTime  time.Time
}
