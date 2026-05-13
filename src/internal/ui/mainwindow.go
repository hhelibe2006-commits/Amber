package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func App(app fyne.App) fyne.Window {
	w := app.NewWindow("Amber")
	tabs := container.NewAppTabs(TaskList(), TaskedItor(), ReStore())
	w.SetContent(tabs)
	w.Resize(fyne.NewSize(600, 400))
	w.CenterOnScreen()
	return w
}
