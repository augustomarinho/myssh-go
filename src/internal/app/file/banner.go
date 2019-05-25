package file

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Banner struct {
	path     string
	fileName string
	content  string
}

func NewBanner() *Banner {
	file := new(Banner)
	file.fileName = ".banner"
	file.path = strings.Split(os.Getenv("GOPATH"), ":")[0]
	file.path += "/assets/banner"
	file.Read()
	return file
}

func (banner *Banner) Read() {
	file, err := os.Open(banner.path)

	if err != nil {
		fmt.Println("failed opening file: %s", err)
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	var str strings.Builder

	for scanner.Scan() {
		str.Write(scanner.Bytes())
	}

	file.Close()

	banner.content = str.String()
}

func (banner Banner) Path() string {
	return banner.path
}

func (banner Banner) Content() string {
	return banner.content
}
