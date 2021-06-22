package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	ff := app.New()

	w := ff.NewWindow("Hello")
	w.SetContent(widget.NewLabel("Hello Fyne!"))

	w.ShowAndRun()
}
