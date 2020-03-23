package screens

import (
	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	log "github.com/sirupsen/logrus"
)

func ConfigListingScreen(win fyne.Window, configs []string) fyne.CanvasObject {
	buttons := widget.NewVBox()
	action := func(name string) {
		log.Printf("button pressed, service: %s", name)
	}
	for _, name := range configs {
		buttons.Append(
			widget.NewButton(name, func() {
				action(name)
			}),
		)
	}
	container := widget.NewGroupWithScroller("Configs", buttons)
	container.Resize(fyne.NewSize(300, 200))
	return container
}
