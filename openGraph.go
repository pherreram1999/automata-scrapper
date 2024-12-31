package main

import (
	"bytes"
	_ "embed"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

//go:embed resources/diagram.svg
var graph []byte

func OpenGraph() {
	win := a.NewWindow("Graph")
	win.Resize(fyne.NewSize(800, 1200))
	imageReader := bytes.NewReader(graph)
	image := canvas.NewImageFromReader(imageReader, "diagram.svg")
	win.SetContent(image)
}
