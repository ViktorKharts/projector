package storage

import (
	"encoding/json"
	"errors"
	"io/fs"
	"os"

	"github.com/viktorkharts/projector/models"
)

func Read() (models.FileData, error) {
	fd := models.FileData{}
	storage := os.Getenv("HOME") + "/projector-storage.json"

	f, err := os.ReadFile(storage)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			_, err := os.Create(storage)
			if err != nil {
				return fd, &storageError{"failed to create a storage file", err.Error()}
			}

			Read()
		} else {
			return fd, &storageError{"failed to read the storage file", err.Error()}
		}
	}

	if len(f) == 0 {
		return fd, &storageError{"your storage is empty", "No system error."}
	}

	if err = json.Unmarshal(f, &fd); err != nil {
		return fd, &storageError{"failed to unmarshal file byte data into json", err.Error()}
	}

	return fd, nil
}
