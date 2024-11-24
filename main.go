package main

import (
	"automata-scrapper/automata"
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

//go:embed assets
var cheemsWeb embed.FS

var w fyne.Window

func main() {
	a := app.NewWithID("automata-search")
	w = a.NewWindow("Automata Search")
	w.Resize(fyne.NewSize(800, 600))

	cont := container.New(
		layout.NewVBoxLayout(),
		urlSearchWidget(),
	)

	w.SetContent(cont)
	w.ShowAndRun()
}

func urlSearchWidget() *fyne.Container {

	urlBind := binding.NewString()
	statusBind := binding.NewString()

	label := widget.NewLabel("URL Site")

	statusLbl := widget.NewLabel("(Status)")
	statusLbl.Bind(statusBind)

	input := widget.NewEntry()
	input.Bind(urlBind)

	hasher := fnv.New32a()

	hashFunc := func(txt string) string {
		hasher.Reset()
		_, _ = hasher.Write([]byte(txt))
		return fmt.Sprintf("%x", hasher.Sum32())
	}

	downloadSiteFunc := func(filename, url string) error {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			return err
		}
		_ = os.Mkdir("sites", os.ModePerm) // nos aseguramos que exista la carpeta
		fileSite, err := os.Create(filename)

		if err != nil {
			return err
		}

		defer fileSite.Close()

		_, err = io.Copy(fileSite, res.Body)

		if err != nil {
			return err
		}
		return nil
	}

	visitSiteFunc := func(url string) {
		filename := fmt.Sprintf("%s.html", hashFunc(url))
		path := filepath.Join("sites", filename)
		_, err := os.Stat(path)

		if errors.Is(err, os.ErrNotExist) {
			if err = downloadSiteFunc(path, url); err != nil {
				ShowError(err)
				return
			}
			statusBind.Set("Sitio descargado localmente")
		} else {
			statusBind.Set("(Sitio cargado desde local)")
		}

		html, err := os.ReadFile(path)

		if err != nil {
			ShowError(err)
			return
		}

		automata.SearchSet(string(html))

		Alert("Sitio Descargado", "El sitio se ha guardado localmente")
	}

	searchFunc := func() {
		url, _ := urlBind.Get()
		url = strings.TrimSpace(url)
		if len(url) == 0 {
			ShowError(errors.New("URL vacia"))
			return
		}

		visitSiteFunc(url)

	}

	searchBtn := widget.NewButtonWithIcon("Search", theme.SearchIcon(), searchFunc)
	buttonContainer := container.New(layout.NewHBoxLayout(), searchBtn, statusLbl)

	return container.New(
		layout.NewVBoxLayout(),
		label,
		input,
		buttonContainer,
	)
}
