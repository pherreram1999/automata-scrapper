package automata

import (
	"encoding/json"
	"fmt"
	"os"
)

const StatusPath = "status"

func SaveStatusInDisk(status *WordStatus) error {
	_ = os.Mkdir(StatusPath, os.ModePerm) // aseguramos que existe la carpeta
	jsonFile, err := os.Create(fmt.Sprintf("%s/%s_status.json", StatusPath, status.Word))
	if err != nil {
		return err
	}
	if err = json.NewEncoder(jsonFile).Encode(status); err != nil {
		return err
	}

	return nil
}
