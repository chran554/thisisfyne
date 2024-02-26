package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	a "thisisfyne/internal/app"
)

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("This is fyne")
	mainLayout := a.ApplicationContent(window)

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(mainLayout)

	window.ShowAndRun()
}
