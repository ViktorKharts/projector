package storage

import (
	"encoding/json"
	"os"

	"github.com/viktorkharts/projector/models"
)

type storageData struct {
	SelectedProject string
	Projects        []models.Project
}

func Write(db models.Storage) error {
	storage := os.Getenv("HOME") + "/projector-storage.json"

	data := storageData{
		SelectedProject: db.SelectedProject,
		Projects:        db.Projects,
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
