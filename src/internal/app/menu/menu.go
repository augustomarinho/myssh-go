package menu

import (
	"bufio"
	"errors"
	"fmt"
	"internal/app/console"
	"internal/app/datastructures"
	"internal/app/file"
	"os"
	"strings"
)

const (
	USERNAME = "username"
)

type Printer struct {
	ini        file.INI
	cmd        *console.Command
	bannerFile *file.Banner
}

func (printer Printer) showRootMenu() {
	fmt.Println("[Environments]")
	for index, content := range printer.ini.Sections()[1:] {
		index++
		fmt.Printf("%d) %s\n", index, content)
	}
}

func (printer Printer) showTypeNumberOrReturn() {
	fmt.Println("Choose number or type 'r' to return")
}

func (printer Printer) BreakLine() {
	fmt.Println("")
}

func (printer Printer) ShowMessage(message ...string) {
	fmt.Println(message)
}

func (printer Printer) ShowError(message error) {
	fmt.Println(message)
}

func NewPrinter(iniFile file.INI) *Printer {
	printer := new(Printer)
	printer.ini = iniFile
	printer.cmd = console.NewCommand()
	printer.bannerFile = file.NewBanner()

	return printer
}

func (printer Printer) Show() {
	printer.cmd.Clear()
	fmt.Println(printer.bannerFile.Content())
	printer.showTypeNumberOrReturn()
	printer.BreakLine()
	printer.showRootMenu()
}

func (printer Printer) ShowSubMenus(menuName string) {
	printer.cmd.Clear()
	fmt.Println(printer.bannerFile.Content())

	var subMenuMap []datastructures.KV = printer.ReadSubMenuItems(menuName)

	fmt.Println("Environment:", menuName)

	if len(subMenuMap) > 0 {
		for index, kv := range subMenuMap {
			fmt.Printf("%d) %s - %s\n", index, kv.Key, kv.Value)
		}
	} else {
		fmt.Println("Invalid Option")
	}

	printer.BreakLine()
	printer.showTypeNumberOrReturn()
}

func (printer Printer) GetMenuName(menuPosition int) (string, error) {
	var secLen = len(printer.ini.Sections())
	if menuPosition < secLen {
		return printer.ini.Sections()[menuPosition], nil
	}
	return "", errors.New("Invalid Option")
}

func (printer Printer) GetSubmenu(menuName string, subMenuPosition int) (string, string, error) {
	return printer.ini.GetSubSection(menuName, subMenuPosition)
}

func (printer Printer) GetConfigByName(configName string) string {
	defaultValuesMap := printer.ini.DefaultSectionValues()

	return defaultValuesMap[configName]
}

func (printer Printer) ReadSubMenuItems(menuName string) []datastructures.KV {
	return printer.ini.SectionValues(menuName)
}

func (printer Printer) ReadInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	return input
}
