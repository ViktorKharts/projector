package models

type Project struct {
	Id    string
	Name  string
	Tasks []Task
}

func (p *Project) OverWriteTasks(t []Task) {
	p.Tasks = t
}

func (p Project) FilterValue() string {
	return p.Name
}
