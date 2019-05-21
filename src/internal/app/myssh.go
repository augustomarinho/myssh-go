package main

import (
	"fmt"
	"internal/app/file"
)

func main() {
	bannerFile := file.NewBanner()
	bannerFile.Read()
	fmt.Println(bannerFile.Content())

	iniFile := file.NewINI()
	iniFile.Read()
	fmt.Println(iniFile.Sections())
}
