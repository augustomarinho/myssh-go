package main

import (
	"fmt"
	"internal/app/file"
	"internal/app/menu"
	menuPkg "internal/app/menu"
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
	for {

		menu.Show()
		var menuInput = menu.ReadInput()

		for {
			if len(menuInput) <= 0 || menuInput == "r" {
				fmt.Println("See you soon ...")
				os.Exit(1)
			}
			menuOption, err := strconv.Atoi(menuInput)

			if err != nil {
				menu.ShowError(err)
			}

			var menuName, errMenuName = menu.GetMenuName(menuOption)
			if errMenuName != nil {
				menu.BreakLine()
				menu.ShowError(errMenuName)
				break
			}

			menu.ShowSubMenus(menuName)

			var subMenuInput = menu.ReadInput()
			if len(subMenuInput) == 0 || subMenuInput == "r" {
				break
			}

			optionSubMenu, errSubMenu := strconv.Atoi(subMenuInput)
			if errSubMenu != nil {
				menu.ShowError(errSubMenu)
			}

			var key, value, errSubMenuOption = menu.GetSubmenu(menuName, optionSubMenu)
			if errSubMenuOption != nil {
				menu.BreakLine()
				menu.ShowError(errSubMenuOption)
				break
			}

			menu.ShowMessage(key, value)
			menu.ShowMessage(menu.GetConfigByName(menuPkg.USERNAME))
			menu.ReadInput()
		}

		menu.BreakLine()
		menu.ShowMessage("Press enter to continue...")
		menu.ReadInput()
	}
}
