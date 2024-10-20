package automata

import "strings"

func AcosoStatus(text string) *WordStatus {
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
			wp.YLine++  // es salto de linea
			wp.XRow = 0 // x empieza a contar de 0
		}

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
			foundCounter++
			currentState = q0 // para volver a empezar a contar
			// agregamos la posicion encontrada
			positions = append(positions, wp)
			wp = &WordPosition{} // nuevo asiganacion en memoria
		} else {
			// se vuelve al primer estado
			currentState = q0
		}

		wp.XRow++ // iteramos el caracter
	}

	return &WordStatus{
		Word:      "acoso",
		Frequency: foundCounter,
		Positions: positions,
	}
}
