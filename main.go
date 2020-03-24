package main

import (
	"context"
	"ev0CloudConfig/config"
	"ev0CloudConfig/logging"
	"ev0CloudConfig/registryreader"
	"ev0CloudConfig/ui"
	"ev0CloudConfig/utils"

	"os"

	"github.com/andygrunwald/vdf"
	"github.com/google/go-github/github"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func init() {
	logging.Init()
	config.Init()
}

func main() {
	csgopathstring := `\steamapps\common\Counter-Strike Global Offensive`
	steampath := registryreader.GetStringFromLocalMachine(`SOFTWARE\Wow6432Node\Valve\Steam`, "InstallPath")
	steamlanguage := registryreader.GetStringFromLocalMachine(`SOFTWARE\Wow6432Node\Valve\Steam`, "Language")
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

	//github
	ctx := context.Background()
	accessToken := viper.GetString("accessToken")
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	ui.UI(configs, csgopath, client)
}
