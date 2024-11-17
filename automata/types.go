package automata

type (
	GraphNodeTransition struct {
		From   int
		To     int
		Symbol string
	}

	GraphData struct {
		Name       string
		Edges      []*GraphNodeTransition
		FinalState int
	}
)
