package menu

import (
	"bufio"
	"fmt"
	"internal/app/console"
	"internal/app/datastructures"
	"internal/app/file"
	"os"
	"strings"
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

func (printer Printer) breakLine() {
	fmt.Println("")
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
	printer.breakLine()
	printer.showRootMenu()
}

func (printer Printer) ShowSubMenus(menuPosition int) {
	printer.cmd.Clear()
	fmt.Println(printer.bannerFile.Content())
	var menuName = printer.GetMenuName(menuPosition)
	var subMenuMap []datastructures.KV = printer.ReadSubMenuItems(menuName)

	fmt.Println("Environment:", menuName)
	for index, kv := range subMenuMap {
		fmt.Printf("%d) %s - %s\n", index, kv.Key, kv.Value)
	}

	printer.breakLine()
	printer.showTypeNumberOrReturn()
}

func (printer Printer) GetMenuName(menuPosition int) string {
	return printer.ini.Sections()[menuPosition]
}

func (printer Printer) GetSubmenu(menuName string, subMenuPosition int) (string, string) {
	return printer.ini.GetSubSection(menuName, subMenuPosition)
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
