package main

import (
	"automata-scrapper/automata"
	"embed"
	_ "embed"
	"errors"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"io"
	"log"
	"net/http"
	"strings"
)

//go:embed assets
var cheemsWeb embed.FS

var loadContainer *fyne.Container

var urlSite string
var windowParent fyne.Window

func visitSite() {
	loadContainer.Show()
	defer loadContainer.Hide()
	res, err := http.Get(urlSite)
	if err != nil {
		dialog.ShowError(err, windowParent)
	}

	if res.StatusCode != 200 {
		dialog.ShowError(errors.New(res.Status), windowParent)
	}

	htmlRes, err := io.ReadAll(res.Body)

	if err != nil {
		dialog.ShowError(err, windowParent)
	}

	htmlText := strings.ToLower(string(htmlRes))

	acecho := automata.WordInspection("acecho", htmlText)
	acoso := automata.WordInspection("acoso", htmlText)
	if err = acecho.RenderGraph(); err != nil {
		dialog.ShowError(err, windowParent)
	}
	acecho.PrintInfo()
	acoso.PrintInfo()
	/*status := []*automata.WordAutomata{
		automata.AcosoAutomata(htmlText),
		automata.AcechoAutomata(htmlText),
	}*/

	/*automata.WordInspection("acecho", htmlText)
	automata.WordInspection("acoso", htmlText)*/

	// empezamos a contar

	/*for _, s := range status {
		if err = automata.SaveStatusInDisk(s); err != nil {
			dialog.ShowError(err, windowParent)
			return
		}
	}*/

	dialog.ShowInformation("Counter", "status guarda", windowParent)
}

func main() {
	a := app.New()
	windowParent = a.NewWindow("Automata Scrapper")
	windowParent.Resize(fyne.NewSize(600, 400))

	// header
	urlBind := binding.BindString(&urlSite)
	inputSite := widget.NewEntry()
	inputSite.Bind(urlBind)
	bntSubmit := widget.NewButton("Cargar", visitSite)
	gridHeader := container.New(layout.NewGridLayout(2), inputSite, bntSubmit)

	// waiting load
	cheemsImgFile, err := cheemsWeb.Open("assets/cheems.jpeg")
	if err != nil {
		log.Fatal(err)
	}
	loadText := canvas.NewText("Cargando sitio,por favor espere", color.Black)
	loadText.TextSize = 16
	loadText.TextStyle.Bold = true
	loadImage := canvas.NewImageFromReader(cheemsImgFile, "cheems.jpeg")
	loadImage.Resize(fyne.NewSize(150, 200))
	loadImage.Move(fyne.NewPos(0, 23))
	// loading container
	loadContainer = container.NewWithoutLayout(loadText, loadImage)
	loadContainer.Hidden = true
	// container app
	allContainer := container.New(layout.NewVBoxLayout(), gridHeader, loadContainer)

	windowParent.SetContent(allContainer)
	windowParent.ShowAndRun()
}
