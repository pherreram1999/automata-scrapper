package automata

import (
	"strings"
)

func AcosoAutomata(text string) *WordAutomata {
	text = strings.ToLower(text)
	// A C O S O
	var q0, q1, q2, q3, q4 State = 0, 1, 2, 3, 4
	currentState := q0

	var foundCounter uint

	var positions []*WordPosition

	wp := &WordPosition{}

	for _, char := range text {

		// por cada caracter se suma 1 en x
		if char == '\n' {
			wp.Line++           // es salto de linea
			wp.CharPosition = 0 // x empieza a contar de 0
		}

		if currentState == q0 && char == 'a' {
			currentState = q1
		} else if currentState == q1 && char == 'c' {
			currentState = q2
		} else if currentState == q2 && char == 'o' {
			currentState = q3
		} else if currentState == q3 && char == 's' {
			currentState = q4
		} else if currentState == q4 && (char == 'o') {
			// estado final
			foundCounter++
			currentState = q0 // para volver a empezar a contar
			// agregamos la posicion encontrada
			positions = append(positions, &WordPosition{ // con esto sacamos una copia y no pasar por memoria
				Line:         wp.Line,
				CharPosition: wp.CharPosition,
			})
		} else {
			// se vuelve al primer estado
			currentState = q0
		}

		wp.CharPosition++ // indicamos que se movio de caracter
	}

	return &WordAutomata{
		Word:      "acoso",
		Frequency: foundCounter,
		Positions: positions,
	}
}

func AcechoAutomata(text string) *WordAutomata {
	// A C E C H O
	var q0, q1, q2, q3, q4, q5 State = 0, 1, 2, 3, 4, 5
	currentState := q0

	ws := &WordAutomata{
		Word: "acecho",
	}

	wp := &WordPosition{}

	for _, char := range text {

		if char == '\n' {
			wp.Line++
			wp.CharPosition = 0
		}

		if currentState == q0 && char == 'a' {
			currentState = q1
		} else if currentState == q1 && char == 'c' {
			currentState = q2
		} else if currentState == q2 && char == 'e' {
			currentState = q3
		} else if currentState == q3 && char == 'c' {
			currentState = q4
		} else if currentState == q4 && char == 'h' {
			currentState = q5
		} else if currentState == q5 && char == 'o' {
			// estamos en estado final con el
			ws.Frequency++
			ws.Positions = append(ws.Positions, &WordPosition{
				Line:         wp.Line,
				CharPosition: wp.CharPosition,
			})
			currentState = q0 // empezamos de nuevo
		} else {
			// un carecter no valido volvemos a empezar
			currentState = q0
		}
		wp.CharPosition++
	}

	return ws
}

func SearchSet(text string) {
	var currentState State = 0 // estado inicial
	for _, char := range text {
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
			// estado final para acoso
			currentState = 5
		} else if currentState == 2 && char == 'e' { // se bifurca el automata en acecho
			currentState = 6
		} else if currentState == 6 && char == 'c' {
			currentState = 7
		} else if currentState == 7 && char == 'h' {
			currentState = 8
		} else {
			currentState = 0 // volvemos al estado incial en cualquier otra transicion no esperada
		}

	}
}
