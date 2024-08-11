// Esta é a raiz da aplicação
package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Cria uma nova aplicação Fyne
	myApp := app.New()

	// Cria uma nova janela com o título "Hello World"
	myWindow := myApp.NewWindow("Hello World")

	// Cria um novo widget de texto com a mensagem "Hello, World!"
	hello := widget.NewLabel("Hello, World!")

	// Adiciona o widget à janela
	myWindow.SetContent(hello)

	// Exibe a janela
	myWindow.ShowAndRun()
}
