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
	cfg          *ini.File
}

func NewINI() *INI {
	var sep = string(os.PathSeparator)

	file := new(INI)
	file.fileName = ".ssh_machines"
	file.path = os.Getenv("HOME")

	//if !Exists(file.path + sep + file.fileName) {
	file.path = strings.Split(os.Getenv("GOPATH"), ":")[0]
	file.fullFilePath = file.path + sep + "assets" + sep + file.fileName
	//}

	//file.fullFilePath = file.path + sep + file.fileName
	return file
}

func (iniFile *INI) Read() {
	//fmt.Println("Reading file INI from", iniFile.fullFilePath)

	// iniLoadOptions := ini.LoadOptions{
	// 	AllowBooleanKeys: true,
	// }

	cfg, err := ini.Load(iniFile.fullFilePath)
	iniFile.cfg = cfg
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		panic(err)
	}

	iniFile.sections = cfg.SectionStrings()
}

func (iniFile INI) Sections() []string {
	return iniFile.sections
}

func (iniFile INI) SectionValues(sectionName string) map[string]string {
	mapValues := make(map[string]string)

	var keys []*ini.Key
	keys = iniFile.cfg.Section(sectionName).Keys()

	for _, key := range keys {
		mapValues[key.Name()] = key.Value()
	}

	return mapValues
}
