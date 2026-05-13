package main

import (
	"fyne.io/fyne/v2/app"
	"github.com/hhelibe2006-commits/Amber/src/internal/ui"
)

func main() {
	a := app.New()
	w := ui.App(a)
	w.ShowAndRun()
}
