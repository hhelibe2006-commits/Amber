package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func App() {
	var a fyne.App = app.New()
	var w = a.NewWindow("App")
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
