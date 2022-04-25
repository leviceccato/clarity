package game

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"os"
	"path/filepath"
)

type saveData struct {
	InputBindings map[command][]any
}

func (g Game) save(profileName string) error {
	// Create saveData that will be serialised into a file
	d := saveData{}
	d.InputBindings = g.InputBindings

	// Encode data as .gob
	var saveBuffer bytes.Buffer
	encoder := gob.NewEncoder(&saveBuffer)
	err := encoder.Encode(d)
	if err != nil {
		return fmt.Errorf("encoding save data: %w", err)
	}

	// Calculate save file location
	absPath, err := filepath.Abs(g.savePath + profileName + ".gob")
	if err != nil {
		return fmt.Errorf("creating save file path: %w", err)
	}

	// Write save file
	err = os.WriteFile(absPath, saveBuffer.Bytes(), 0644)
	if err != nil {
		return fmt.Errorf("writing save file: %w", err)
	}

	return nil
}
