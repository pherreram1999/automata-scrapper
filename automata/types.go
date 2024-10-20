package automata

type (
	State uint

	WordPosition struct {
		XRow  int `json:"x_row"`
		YLine int `json:"y_line"`
	}

	WordStatus struct {
		Word      string          `json:"word"`
		Frequency uint            `json:"frequency"`
		Positions []*WordPosition `json:"positions"`
	}
)
