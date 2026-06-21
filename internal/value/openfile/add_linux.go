package openfile

import (
	"os"
	"syscall"
)

func (o *OpenFile) Add(path string) error {
	f, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		return err
	}
	*o = append(*o, f)
	return nil
}
