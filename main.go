package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func TwoDimVecWindow(appl fyne.App, home fyne.Window) {

}

func main() {
	myApp := app.New()
	home := myApp.NewWindow("Vector Caluclator")
	home.SetMaster()

	home.SetContent(container.NewVBox(
		widget.NewButton("2D Vectors", func() {
			TwoDimVecWindow(myApp, home)
		}),
	))
	home.Resize(fyne.NewSize(100, 100))
	home.ShowAndRun()
}
