package main

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"os"
	"text/tabwriter"
)

type WordStatus struct {
	Word           string
	frequency      uint
	FrequeyBind    binding.String
	FileWordPlaces *os.File
	TabWriter      *tabwriter.Writer
}

func (ws *WordStatus) Plus(line, charPos uint) {
	ws.frequency++
	ws.FrequeyBind.Set(fmt.Sprintf("%d", ws.frequency))
	_, _ = fmt.Fprintf(ws.TabWriter, "%d\t%d\n", line, charPos)
}

func (ws *WordStatus) Reset() {
	ws.frequency = 0
	ws.FrequeyBind.Set(fmt.Sprintf("%d", ws.frequency))
}
