package automata

import "os"

type GraphImage struct {
	data []byte
}

func (gi *GraphImage) Save(path string) error {
	return os.WriteFile(path, gi.data, os.ModePerm)
}
