package main

import (
	"ev0CloudConfig/gui"
	"ev0CloudConfig/logging"
	"ev0CloudConfig/registryReader"
	"ev0CloudConfig/utils"

	"os"

	"github.com/andygrunwald/vdf"
	log "github.com/sirupsen/logrus"
)

func init() {
	logging.Init()
}

func main() {
	csgopathstring := `\steamapps\common\Counter-Strike Global Offensive`
	steampath := registryReader.GetStringFromLocalMachine(`SOFTWARE\Wow6432Node\Valve\Steam`, "InstallPath")
	steamlanguage := registryReader.GetStringFromLocalMachine(`SOFTWARE\Wow6432Node\Valve\Steam`, "Language")
	log.Infof("Steam is installed under: %s", steampath)
	log.Infof("Steam default language is: %s", steamlanguage)

	var csgopath string

	if !utils.FileExists(steampath + `\steamapps\appmanifest_730.acf`) {
		libraryfolderspath := steampath + `\steamapps\libraryfolders.vdf`
		log.Info(libraryfolderspath)
		libraryfoldersvdf, _ := os.Open(libraryfolderspath)
		p := vdf.NewParser(libraryfoldersvdf)
		m, err := p.Parse()
		if err != nil {
			log.Fatal(err)
		}
		csgopath = m["LibraryFolders"].(map[string]interface{})["1"].(string) + csgopathstring
	} else {
		csgopath = steampath + csgopathstring
	}
	log.Infof("CSGO is installed under: %s", csgopath)

	configs := utils.WalkPath(csgopath + `\ev0lve`)
	for _, file := range configs {
		log.Info(file)
	}
	gui.Init(configs)
}
