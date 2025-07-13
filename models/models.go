package models

type Storage struct {
	SelectedProject string
	Projects        map[string]Project
}

type Project struct {
	Id    string
	Name  string
	Tasks []Task
}

type Task struct {
	Id         string
	Value      string
	IsComplete bool
}
