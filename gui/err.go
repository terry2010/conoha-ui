package Gui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"image/color"
)


func ErrorAlert(msg string) {

	myApp := app.New()
	myWindow := myApp.NewWindow("conoha-UI error")

	//img := canvas.NewImageFromResource(theme.CancelIcon())
	//img.FillMode = canvas.ImageFillOriginal
	//
	text := canvas.NewText("error:"+msg, color.White)
	content := fyne.NewContainerWithLayout(layout.NewCenterLayout(),
		text)
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(300, 200))
	myWindow.ShowAndRun()

}
