package automata

import "strings"

func AcosoCounter(text string) uint {
	text = strings.ToLower(text)
	// A C O S O
	var q0, q1, q2, q3, q4 State = 0, 1, 2, 3, 4
	currentState := q0

	var counter uint

	for _, char := range text {

		if currentState == q0 && char == 'a' {
			currentState = q1
		} else if currentState == q1 && char == 'c' {
			currentState = q2
		} else if currentState == q2 && char == 'o' {
			currentState = q3
		} else if currentState == q3 && char == 's' {
			currentState = q4
		} else if currentState == q4 && (char == 'o' || char == 'ó') {
			// estado final
			counter++
			currentState = q0 // para volver a empezar a contar
		} else {
			// se vuelve al primer estado
			currentState = q0
		}
	}

	return counter
}
