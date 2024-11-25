package main

import (
	"fmt"
	"fyne.io/fyne/v2/data/binding"
)

type WordStatus struct {
	Word        string
	frequency   uint
	FrequeyBind binding.String
}

func (ws *WordStatus) Plus() {
	ws.frequency++
	ws.FrequeyBind.Set(fmt.Sprintf("%d", ws.frequency))
}
