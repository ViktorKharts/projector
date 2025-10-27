package models

type Project struct {
	Id      string
	Name    string
	Columns []Column
}

// required to implement the list.Item interface
func (p Project) FilterValue() string {
	return p.Name
}
