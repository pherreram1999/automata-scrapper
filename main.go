package main

import (
	"embed"
	_ "embed"
	"errors"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"hash/fnv"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

//go:embed assets
var cheemsWeb embed.FS

var w fyne.Window

func main() {
	a := app.NewWithID("automata-search")
	w = a.NewWindow("Automata Scrapper")
	w.Resize(fyne.NewSize(800, 600))

	wordsSet := InitSetWords()

	pieWidget := container.New(layout.NewVBoxLayout())

	totalBind := binding.NewString()

	gridInfo := container.New(
		layout.NewHBoxLayout(),
		setWordsWidget(wordsSet, totalBind),
		pieWidget,
	)

	cont := container.New(
		layout.NewVBoxLayout(),
		// contiene la barra de busqueda
		urlSearchWidget(wordsSet, pieWidget, totalBind),
		// son los resultados
		gridInfo,
	)

	w.SetContent(cont)
	w.ShowAndRun()
}

func setWordsWidget(setWords SetWords, totalBind binding.String) *fyne.Container {

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

	totalLbl := widget.NewLabel("total")
	resultLbl := widget.NewLabel("0")
	resultLbl.Bind(totalBind)

	cont.Add(container.New(layout.NewHBoxLayout(), totalLbl, resultLbl))

	return cont
}

func urlSearchWidget(setWords SetWords, pieWidget *fyne.Container, totalBind binding.String) *fyne.Container {

	urlBind := binding.NewString()
	statusBind := binding.NewString()

	label := widget.NewLabel("URL Site")

	statusLbl := widget.NewLabel("(Status)")
	statusLbl.Bind(statusBind)

	filePathBind := binding.NewString()
	filePathLabel := widget.NewLabel("")
	filePathLabel.Bind(filePathBind)

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

	searchProcess := func(htmlStr, filename string) {
		// normalizamos el html
		htmlStr = strings.ToLower(strings.TrimSpace(string(htmlStr)))
		// abrimos los archivos de cada palabra para guardar las concurrencias
		if err := setWords.OpenFileOcurrences(filename); err != nil {
			ShowError(err)
			return
		}
		// esta funcion es la contiene el automata de busqueda
		SearchSet(htmlStr, setWords)
		defer setWords.CloseFiles()

		// generamos la grafica
		pieChartData, err := setWords.RenderPie()
		if err != nil {
			ShowError(err)
			return
		}
		pieImg := canvas.NewImageFromReader(pieChartData, "pie.png")
		pieImg.FillMode = canvas.ImageFillOriginal
		pieWidget.RemoveAll() // quitamos la ultima generada
		pieWidget.Add(pieImg)

		totalBind.Set(fmt.Sprintf("%d", setWords.TotalFrequency()))

	}

	/**
	Se encarga de visitar el sitio y guardar una copia y realizar el conteo de palabras
	*/
	visitSiteFunc := func(url string) {
		filename := fmt.Sprintf("%s.html", hashFunc(url))
		_ = os.Remove(filename)
		path := filepath.Join("sites", filename)

		if err := downloadSiteFunc(path, url); err != nil {
			ShowError(err)
			return
		}

		filePathBind.Set(path)

		html, err := os.ReadFile(path)

		if err != nil {
			ShowError(err)
			return
		}

		searchProcess(string(html), filename)
	}

	searchFunc := func() {
		url, _ := urlBind.Get()
		url = strings.TrimSpace(url)
		if len(url) == 0 {
			ShowError(errors.New("URL vacia"))
			return
		}
		// valiadamos que empieza con http el domincio para que sea un url valida

		mustHttp := regexp.MustCompile(`http*`)

		if !mustHttp.MatchString(url) {
			ShowError(errors.New("No es una URL valida,debe emepzar con http"))
		}

		visitSiteFunc(url)
	}

	openFileFunc := func() {
		dialog.ShowFileOpen(
			func(reader fyne.URIReadCloser, err error) {
				if err != nil {
					ShowError(err)
					return
				}

				path := reader.URI().Path()
				file, err := os.Open(path)
				if err != nil {
					ShowError(err)
					return
				}

				defer file.Close()

				htmlStr, err := io.ReadAll(file)

				if err != nil {
					ShowError(err)
					return
				}

				filePathBind.Set(path)
				searchProcess(string(htmlStr), filepath.Base(path))
			},
			w,
		)
	}

	searchBtn := widget.NewButtonWithIcon("Search In Website", theme.SearchIcon(), searchFunc)

	pickFileBtn := widget.NewButtonWithIcon("Open File", theme.FileIcon(), openFileFunc)

	buttonContainer := container.New(layout.NewHBoxLayout(), pickFileBtn, searchBtn, statusLbl, filePathLabel)

	return container.New(
		layout.NewVBoxLayout(),
		label,
		input,
		buttonContainer,
	)
}
