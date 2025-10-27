package models

type Column struct {
	Id    string
	Name  string
	Tasks []Task
}

func (c *Column) OverWriteTasks(t []Task) {
	c.Tasks = t
}
