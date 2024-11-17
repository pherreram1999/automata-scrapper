package automata

import (
	"bytes"
	_ "embed"
	"fmt"
	"os/exec"
	"text/template"
)

//go:embed graphviz.dot
var graphizLayout string

type (
	State uint

	WordPosition struct {
		CharPosition int `json:"char_position"`
		Line         int `json:"line"`
	}

	WordAutomata struct {
		Word      string          `json:"word"`
		Frequency uint            `json:"frequency"`
		Positions []*WordPosition `json:"positions"`
	}
)

func CopyPosition(wp *WordPosition) *WordPosition {
	return &WordPosition{
		Line:         wp.Line,
		CharPosition: wp.CharPosition,
	}
}

func WordInspection(word, text string) *WordAutomata {
	/**
	Tomando la premisa de las expresiones regulares son un serie de simbolos concatenados
	se observa un avance lineal (cambio de estado) por cada simbolo al siguiente,
	por lo que se puede automatizar
	la entrada y la salida
	*/
	wordSymbols := []rune(word)
	currentState := 0
	finalState := len(wordSymbols)
	wa := &WordAutomata{Word: word}
	wp := &WordPosition{}

	for _, char := range text {
		if char == '\n' {
			wp.Line++
			wp.CharPosition = 0
		}
		if currentState == finalState {
			wa.Frequency++
			wa.Positions = append(wa.Positions, CopyPosition(wp))
			currentState = 0
		} else if wordSymbols[currentState] == char {
			currentState++
		} else {
			currentState = 0
		}

		wp.CharPosition++
	}

	return wa
}

func (wa *WordAutomata) PrintInfo() *WordAutomata {
	fmt.Println("word:\t", wa.Word)
	fmt.Println("frequency:\t", wa.Frequency)
	return wa
}

func (wa *WordAutomata) RenderGraph() error {

	g := &GraphData{Name: wa.Word}

	g.FinalState = len(wa.Word) + 1

	var i int
	var symbol string
	for i = 0; i < g.FinalState; i++ {
		if i+1 == g.FinalState {
			symbol = ""
		} else {
			symbol = string(wa.Word[i])
		}
		g.Edges = append(g.Edges, &GraphNodeTransition{
			From:   i,
			To:     i + 1,
			Symbol: symbol,
		})
	}

	// epislon moves
	for i = 1; i < g.FinalState; i++ {
		g.Edges = append(g.Edges, &GraphNodeTransition{
			From:   i,
			To:     0,
			Symbol: "ε",
		})
	}

	tmpl, err := template.New(g.Name).Parse(graphizLayout)

	if err != nil {
		return err
	}

	var dotLayout, svg bytes.Buffer

	if err = tmpl.Execute(&dotLayout, g); err != nil {
		return err
	}

	cmd := exec.Command("dot", "-Tsvg")

	cmd.Stdin = &dotLayout
	cmd.Stdout = &svg

	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}
