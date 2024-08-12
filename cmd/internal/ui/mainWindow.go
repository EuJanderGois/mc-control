package ui

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/eujandergois/mc-control/cmd/lang"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func NewMainWindow(app fyne.App) fyne.Window {
	systemLang := lang.GetSystemLanguage()
	fmt.Println("Detected Sytem Language:", display.English.Tags().Name(systemLang))
	fmt.Println(systemLang)

	// Verificação de linguagem. No futuro adionar verificação: Se o usuário houver
	// modificado a linguagem nas configurações da aplicação ignorar checagem e exibir linguagem
	// selecionada pelo usuário. Isso deve ser implementado em conjunto com as configurações.
	var defaultLang lang.Lang
	if systemLang == language.BrazilianPortuguese {
		defaultLang = lang.PT_BR
	}
	if systemLang == language.AmericanEnglish {
		defaultLang = lang.EN_US
	}
	if systemLang != language.AmericanEnglish || systemLang != language.BrazilianPortuguese {
		defaultLang = lang.EN_US
	}

	var appLang, err = lang.GetLang(defaultLang)
	if err != nil {
		log.Fatalf("Erro ao carregar a linguagem: %v", err)
	}

	window := app.NewWindow(appLang.AppName)

	window.Resize(fyne.Size{
		Width:  840,
		Height: 520,
	})

	label := widget.NewLabel("Hello, Fyne!")
	windowName := widget.NewLabel(appLang.AppWindows.Home.Name)

	main := container.NewVBox()
	topbox := container.NewHBox()
	vbox := container.NewCenter()

	topbox.Add(windowName)
	vbox.Add(label)

	main.Add(topbox)
	main.Add(vbox)

	window.SetContent(main)

	return window
}
