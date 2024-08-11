package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewMainWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("My Fyne App")

	window.Resize(fyne.Size{
		Width:  540,
		Height: 320,
	})

	label := widget.NewLabel("Hello, Fyne!")

	vbox := container.NewCenter()

	vbox.Add(label)

	window.SetContent(vbox)

	return window
}
