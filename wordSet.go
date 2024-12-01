package main

import (
	"bytes"
	"errors"
	"fmt"
	"fyne.io/fyne/v2/data/binding"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

type SetWords map[string]*WordStatus

func (words SetWords) RenderPie() (*bytes.Buffer, error) {
	labels := []string{}
	data := []string{}

	total := words.TotalFrequency()

	for _, word := range words {
		labels = append(labels, word.Word)
		data = append(data, fmt.Sprintf("%d", (word.frequency*100)/total))
	}

	labelsStr := strings.Join(labels, ",")
	dataStr := strings.Join(data, ",")

	cmd := exec.Command("python3", "pie.py", `--labels=`+labelsStr, `--data=`+dataStr)

	var buffer, errBuff bytes.Buffer

	cmd.Stdout = &buffer
	cmd.Stderr = &errBuff

	if err := cmd.Run(); err != nil {
		return nil, errors.Join(
			errors.New("No fue posible renderizar la grafica"),
			err,
			errors.New(errBuff.String()),
		)
	}

	return &buffer, nil
}

func (words SetWords) TotalFrequency() uint {
	var total uint
	for _, word := range words {
		total += word.frequency
	}
	return total
}

func (words SetWords) CloseFiles() {
	for _, word := range words {
		word.TabWriter.Flush()
		word.FileWordPlaces.Close()
	}
}

func (words SetWords) OpenFileOcurrences(filename string) error {
	var err error
	for _, ws := range words {
		ws.FileWordPlaces, err = os.Create("word-ocurrences/" + ws.Word + ".txt")

		if err != nil {
			return err
		}

		fmt.Fprintln(ws.FileWordPlaces, "File: "+filename+"\n")

		ws.TabWriter = tabwriter.NewWriter(ws.FileWordPlaces, 0, 4, 4, '\t', 0)
		_, _ = fmt.Fprint(ws.TabWriter, "Line\tchar\n")
	}
	return nil
}

func (words SetWords) Reset() {
	for _, word := range words {
		word.Reset()
	}
}

func InitSetWords() SetWords {
	words := []string{"acoso", "acecho", "víctima", "violación", "machista", "agresión"}
	_ = os.RemoveAll("word-ocurrences")
	_ = os.Mkdir("word-ocurrences", os.ModePerm)
	wordStatus := make(SetWords, len(words))
	for _, word := range words {
		ws := &WordStatus{
			Word:        word,
			FrequeyBind: binding.NewString(),
		}
		wordStatus[word] = ws
	}
	return wordStatus
}
