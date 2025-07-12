package models

type FileData struct {
	SelectedProject string
	Projects        []Project
}

type Project struct {
	Name  string
	Tasks []Task
}

type Task struct {
	Value      string
	IsComplete bool
}
