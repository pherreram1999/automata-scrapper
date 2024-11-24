package main

import "fyne.io/fyne/v2/dialog"

func ShowError(err error) {
	dialog.ShowError(err, w)
}

func Alert(title, message string) {
	dialog.ShowInformation(title, message, w)
}
