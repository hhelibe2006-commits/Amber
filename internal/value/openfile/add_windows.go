package openfile

import (
	"os"

	"golang.org/x/sys/windows"
)

func (o *OpenFile) Add(path string) error {
	filePath, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return err
	}
	handle, err := windows.CreateFile(
		filePath,
		windows.GENERIC_READ,
		windows.FILE_SHARE_READ,
		nil,
		windows.OPEN_ALWAYS,
		windows.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return err
	}
	file := os.NewFile(uintptr(handle), path)
	*o = append(*o, file)
	return nil
}
