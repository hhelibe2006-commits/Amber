package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TaskedItor() *container.TabItem {
	box := container.NewVBox()
	btn := widget.NewButton("新建", func() {})
	box.Add(btn)
	tab := container.NewTabItem("新建", box)
	return tab
}
