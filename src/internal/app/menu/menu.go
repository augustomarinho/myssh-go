package menu

import (
	"bufio"
	"fmt"
	"internal/app/console"
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
	fmt.Println("Environments:")
	for index, content := range printer.ini.Sections()[1:] {
		fmt.Printf("%d) %s\n", index, content)
	}
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
	printer.showRootMenu()
}

func (printer Printer) ShowOptions(option int) map[string]string {
	option++
	var sectionName = printer.ini.Sections()[option]
	var sectionValuesMap map[string]string = printer.ReadOptions(sectionName)

	var item int = 1
	for k, v := range sectionValuesMap {
		fmt.Printf("%d) %s - %s\n", item, k, v)
		item++
	}

	return sectionValuesMap
}

func (printer Printer) ReadOptions(sectionName string) map[string]string {
	return printer.ini.SectionValues(sectionName)
}

func (printer Printer) ReadInput() string {

	fmt.Println("Option:")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}

	input = strings.TrimSpace(input)
	return input
}
