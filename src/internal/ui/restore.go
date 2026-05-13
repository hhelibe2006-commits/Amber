package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ReStore() *container.TabItem {
	box := container.NewVBox()
	box.Add(hbox())
	tab := container.NewTabItem("恢复", box)
	return tab
}

func hbox() *fyne.Container {
	box := container.NewHBox()
	btn := widget.NewButton("恢复", func() {})
	entry := widget.NewEntry()
	box.Add(entry)
	box.Add(btn)
	return box
}
