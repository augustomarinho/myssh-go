package file

import (
	"errors"
	"fmt"
	"internal/app/datastructures"
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

func (iniFile INI) SectionValues(sectionName string) []datastructures.KV {

	var keys []*ini.Key
	keys = iniFile.cfg.Section(sectionName).Keys()

	var array = make([]datastructures.KV, len(keys))

	for index, key := range keys {
		var kv = datastructures.NewKV()
		kv.Key = key.Name()
		kv.Value = key.Value()
		array[index] = *kv
	}

	return array
}

func (iniFile INI) GetSubSection(sectionName string, position int) (string, string, error) {
	var keySize = len(iniFile.cfg.Section(sectionName).Keys())

	if position < keySize {
		key := iniFile.cfg.Section(sectionName).Keys()[position]
		return key.Name(), key.Value(), nil
	}

	return "", "", errors.New("Invalid Option")
}

func (iniFile INI) DefaultSectionValues() map[string]string {

	var sec, err = iniFile.cfg.GetSection(ini.DEFAULT_SECTION)

	if err != nil {
		panic(err)
	}

	var mapDefaultValues map[string]string = make(map[string]string)

	for _, key := range sec.Keys() {
		mapDefaultValues[key.Name()] = key.Value()
	}

	return mapDefaultValues
}
