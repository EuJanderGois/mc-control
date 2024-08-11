package app

import (
	"fyne.io/fyne/v2/app"
	"github.com/eujandergois/mc-control/cmd/internal/ui"
)

func Run() {
	a := app.New()
	w := ui.NewMainWindow(a)
	w.ShowAndRun()
}
