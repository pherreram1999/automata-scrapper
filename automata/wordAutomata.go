package automata

import (
	"fmt"
)

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

	return nil
}
