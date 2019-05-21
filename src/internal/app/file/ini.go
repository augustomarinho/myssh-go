package file

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-ini/ini"
)

type INI struct {
	path         string
	fileName     string
	sections     []string
	fullFilePath string
}

func NewINI() *INI {
	var sep = string(os.PathSeparator)

	file := new(INI)
	file.fileName = ".ssh_machines"
	file.path = os.Getenv("HOME")

	if !Exists(file.path + sep + file.fileName) {
		file.path = strings.Split(os.Getenv("GOPATH"), ":")[0]
	}

	file.fullFilePath = file.path + sep + file.fileName
	return file
}

func (iniFile *INI) Read() {
	fmt.Println("Reading file INI")

	cfg, err := ini.Load(iniFile.fullFilePath)
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		panic(err)
	}

	iniFile.sections = cfg.SectionStrings()
}

func (iniFile INI) Sections() []string {
	return iniFile.sections
}
