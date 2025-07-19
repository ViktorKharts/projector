package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/viktorkharts/projector/models"
)

func Read() (models.Storage, error) {
	s := models.Storage{
		SelectedProject: "",
		Projects:        []models.Project{},
	}

	storage := os.Getenv("HOME") + "/projector-storage.json"

	f, err := os.ReadFile(storage)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			_, err := os.Create(storage)
			if err != nil {
				return s, &storageError{"failed to create a storage file", err.Error()}
			}

			Read()
		} else {
			return s, &storageError{"failed to read the storage file", err.Error()}
		}
	}

	if len(f) == 0 {
		return s, &storageError{"your storage is empty", "No system error."}
	}

	if err = json.Unmarshal(f, &s); err != nil {
		return s, &storageError{"failed to unmarshal file byte data into json", err.Error()}
	}

	return s, nil
}
