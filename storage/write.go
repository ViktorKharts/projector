package storage

import (
	"encoding/json"
	"os"

	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/ui"
)

type storageData struct {
	SelectedProject string
	Projects        []models.Project
}

func Write(m ui.Main) error {
	storage := os.Getenv("HOME") + "/projector-storage.json"

	data := storageData{
		SelectedProject: m.SelectedProject,
		Projects:        m.Projects,
	}

	bd, err := json.Marshal(data)
	if err != nil {
		return &storageError{"failed to Marshal Project data before save", err.Error()}
	}

	if err := os.WriteFile(storage, bd, 0666); err != nil {
		return &storageError{"failed to write Project data into storage file", err.Error()}
	}

	return nil
}
