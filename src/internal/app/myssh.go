package main

import (
	"internal/app/file/banner"
	"internal/app/file/ini"
)

func main() {
	bannerFile := banner.NewBanner()
	iniFile := new(ini.File)
	iniFile.Read()
	bannerFile.Read()
}
