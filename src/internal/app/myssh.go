package main

import (
	"fmt"
	"internal/app/file"
	"internal/app/menu"
	"strconv"
)

func main() {
	myssh := NewMyssh()
	myssh.Run()
}

type Myssh struct {
}

func NewMyssh() *Myssh {
	myssh := new(Myssh)
	return myssh
}

func (myssh *Myssh) Run() {
	iniFile := file.NewINI()
	iniFile.Read()

	menu := menu.NewPrinter(*iniFile)
	menu.Show()

	var input = menu.ReadInput()
	option, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println(err)
	}

	menu.ShowOptions(option)
}
