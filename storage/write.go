package storage

import (
	"encoding/json"
	"os"

	"github.com/viktorkharts/projector/models"
)

func Write(db models.FileData, p models.Project) error {
	storage := os.Getenv("HOME") + "/projector-storage.json"

	db.SelectedProject = p.Name
	db.Projects = append(db.Projects, p)

	bd, err := json.Marshal(db)
	if err != nil {
		return &storageError{"failed to Marshal Project data before save", err.Error()}
	}

	if err := os.WriteFile(storage, bd, 0666); err != nil {
		return &storageError{"failed to write Project data into storage file", err.Error()}
	}

	return nil
}
