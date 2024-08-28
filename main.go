package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	"github.com/tadeaspetak/fyne-notifications/res"
)

func main() {
	a := app.NewWithID("com.tadeaspetak.fyne-notifications")
	a.SetIcon(res.ResourceIconPng)

	w := a.NewWindow("Notification Example")

	button := widget.NewButton("click me", func() {
		a.SendNotification(fyne.NewNotification("Title", "Button clicked."))
	})
	w.SetContent(button)

	w.Resize(fyne.NewSize(200, 200))
	w.ShowAndRun()
}
