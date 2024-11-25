package main

import (
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

	wordsSet := InitSetWords()

	gridInfo := container.New(
		layout.NewGridLayout(2),
		setWordsWidget(wordsSet),
	)

	cont := container.New(
		layout.NewVBoxLayout(),
		urlSearchWidget(wordsSet),
		gridInfo,
	)

	w.SetContent(cont)
	w.ShowAndRun()
}

type SetWords map[string]*WordStatus

func (words SetWords) Reset() {
	for _, word := range words {
		word.Reset()
	}
}

func InitSetWords() SetWords {
	words := []string{"acoso", "acecho", "victima", "violacion", "machista"}
	wordStatus := make(SetWords, len(words))
	for _, word := range words {
		wordStatus[word] = &WordStatus{
			Word:        word,
			FrequeyBind: binding.NewString(),
		}
	}
	return wordStatus
}

func setWordsWidget(setWords SetWords) *fyne.Container {

	cont := container.New(layout.NewVBoxLayout())

	for _, ws := range setWords {
		frequencyLbl := widget.NewLabel("0")
		frequencyLbl.Bind(ws.FrequeyBind)
		cont.Add(
			container.New(
				layout.NewGridLayout(2),
				widget.NewLabel(ws.Word),
				frequencyLbl,
			),
		)
	}

	return cont
}

func urlSearchWidget(setWords SetWords) *fyne.Container {

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
			return errors.New(res.Status)
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

		htmlStr := strings.ToLower(strings.TrimSpace(string(html)))
		SearchSet(htmlStr, setWords)

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
