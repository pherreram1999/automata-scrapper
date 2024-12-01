package main

import (
	"automata-scrapper/automata"
)

func SearchSet(text string, words SetWords) {
	var currentState automata.State = 0 // estado inicial
	words.Reset()
	var linePos, charPos uint = 1, 0

	for _, char := range text {
		if char == '\n' {
			linePos++
			charPos = 0
		}
		charPos++
		// condicionales para acoso
		if currentState == 0 && char == 'a' {
			currentState = 1
		} else if currentState == 1 && char == 'c' {
			currentState = 2
		} else if currentState == 2 && char == 'o' {
			currentState = 3
		} else if currentState == 3 && char == 's' {
			currentState = 4
		} else if currentState == 4 && char == 'o' {
			// currentState = 5 // ESTADO final "acoso"
			words["acoso"].Plus(linePos, charPos)
			currentState = 0
		} else if currentState == 2 && char == 'e' { // se bifurca el automata en acecho
			currentState = 6
		} else if currentState == 6 && char == 'c' {
			currentState = 7
		} else if currentState == 7 && char == 'h' {
			currentState = 8
		} else if currentState == 8 && char == 'o' {
			// currentState = 9 ESTADO FINAL "acecho"
			words["acecho"].Plus(linePos, charPos)
		} else if currentState == 1 && char == 'g' {
			currentState = 10
		} else if currentState == 10 && char == 'r' {
			currentState = 11
		} else if currentState == 11 && char == 'e' {
			currentState = 12
		} else if currentState == 12 && char == 's' {
			currentState = 13
		} else if currentState == 13 && char == 'i' {
			currentState = 14
		} else if currentState == 14 && (char == 'ó' || char == 'o') {
			currentState = 15
		} else if currentState == 15 && char == 'n' {
			// currentState = 16 ESTADO FINAL "agresion"
			words["víctima"].Plus(linePos, charPos)
		} else if currentState == 0 && char == 'v' { // inicia victima
			currentState = 17
		} else if currentState == 17 && (char == 'i' || char == 'í') {
			currentState = 18
		} else if currentState == 18 && char == 'c' {
			currentState = 19
		} else if currentState == 19 && char == 't' {
			currentState = 20
		} else if currentState == 20 && char == 'i' {
			currentState = 21
		} else if currentState == 21 && char == 'm' {
			currentState = 22
		} else if currentState == 22 && char == 'a' {
			// currentState = 23 ESTADO FINAL "victima"
			words["víctima"].Plus(linePos, charPos)
			currentState = 0
		} else if currentState == 18 && char == 'o' {
			currentState = 24
		} else if currentState == 24 && char == 'l' {
			currentState = 25
		} else if currentState == 25 && char == 'a' {
			currentState = 26
		} else if currentState == 26 && char == 'c' {
			currentState = 27
		} else if currentState == 27 && char == 'i' {
			currentState = 28
		} else if currentState == 28 && (char == 'ó' || char == 'o') {
			currentState = 29
		} else if currentState == 29 && char == 'n' {
			// currentState = 30 ESTADO FINAL de "violación"
			words["violación"].Plus(linePos, charPos)
			currentState = 0
		} else if currentState == 0 && char == 'm' {
			currentState = 31
		} else if currentState == 31 && char == 'a' {
			currentState = 32
		} else if currentState == 32 && char == 'c' {
			currentState = 33
		} else if currentState == 33 && char == 'h' {
			currentState = 34
		} else if currentState == 34 && char == 'i' {
			currentState = 35
		} else if currentState == 35 && char == 's' {
			currentState = 36
		} else if currentState == 36 && char == 't' {
			currentState = 37
		} else if currentState == 37 && char == 'a' {
			// currentState = 38 Estado final de "machista"
			words["machista"].Plus(linePos, charPos)
			currentState = 0
		} else {
			currentState = 0 // volvemos al estado incial en cualquier otra transicion no esperada
		}

	}
}
