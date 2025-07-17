package models

type Storage struct {
	SelectedProject string
	Projects        map[string]Project
}
