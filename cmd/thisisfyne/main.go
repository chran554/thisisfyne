package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
	a "thisisfyne/internal/app"
	"thisisfyne/internal/app/resources"
)

func main() {
	application := app.New()
	application.SetIcon(theme.ComputerIcon())

	window := application.NewWindow("This is fyne")
	window.SetIcon(resources.ThisIsFyneIconResource)

	mainLayout := a.ApplicationContent(window)

	window.Resize(fyne.NewSize(800, 600))
	window.SetContent(mainLayout)

	window.ShowAndRun()
}
