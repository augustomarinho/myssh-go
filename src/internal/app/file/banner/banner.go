package banner

import (
	"fmt"
	"os"
)

type File struct {
	path     string
	fileName string
}

func NewBanner() *File {
	file := new(File)
	file.fileName = ".banner"
	file.path = os.Getenv("HOME")
	return file
}

func (banner *File) Read() {
	fmt.Println("Reading banner file from path", banner.Path())
}

func (banner File) Path() string {
	return banner.path
}
