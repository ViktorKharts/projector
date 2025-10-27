package models

type Task struct {
	Id          string
	Title       string
	Description string
	IsComplete  bool
}

func (t *Task) ToggleIsCompleteTask() {
	t.IsComplete = !t.IsComplete
}

func (t *Task) EditTaskValue(nt string) {
	t.Title = nt
}

func (t *Task) EditTaskDescription(nd string) {
	t.Title = nd
}
