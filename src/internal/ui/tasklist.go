package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TaskList() *container.TabItem {
	label := widget.NewLabel("Task List")
	tab := container.NewTabItem("任务列表", label)
	return tab
}
