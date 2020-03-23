package gui

import (
	"ev0CloudConfig/gui/screens"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	log "github.com/sirupsen/logrus"
)

// Init init the gui
func Init(configs []string) {
	log.Info("GUI Init")
	os.Setenv("FYNE_SCALE", "1")
	a := app.New()
	w := a.NewWindow("ev0-cloud-config")

	w.SetContent(screens.ConfigListingScreen(w, configs))
	w.Resize(fyne.NewSize(300, 200))
	w.ShowAndRun()
}
