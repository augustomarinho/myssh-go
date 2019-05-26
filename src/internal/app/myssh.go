package main

import (
	"fmt"
	"internal/app/file"
	"internal/app/menu"
	"os"
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
	for true {
		menu.Show()
		var menuInput = menu.ReadInput()

		if len(menuInput) == 0 || menuInput == "r" {
			fmt.Println("good bye")
			os.Exit(1)
		}
		menuOption, err := strconv.Atoi(menuInput)

		if err != nil {
			fmt.Println(err)
		}

		menu.ShowSubMenus(menuOption)
		var menuName = menu.GetMenuName(menuOption)
		var subMenuInput = menu.ReadInput()
		if len(subMenuInput) == 0 || subMenuInput == "r" {
			continue
		}

		optionSubMenu, errSubMenu := strconv.Atoi(subMenuInput)
		if errSubMenu != nil {
			fmt.Println(errSubMenu)
		}

		var key, value = menu.GetSubmenu(menuName, optionSubMenu)
		fmt.Println(key, value)
	}
}
