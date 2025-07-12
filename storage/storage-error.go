package storage

import "fmt"

type storageError struct {
	message     string
	systemError string
}

func (e *storageError) Error() string {
	return fmt.Sprintf("Projector Storage Error: %s.\n%s", e.message, e.systemError)
}
