package models

type Task struct {
	Id         string
	Value      string
	IsComplete bool
}

func (t *Task) ToggleIsCompleteTask() {
	t.IsComplete = !t.IsComplete
}

func (t *Task) EditTaskValue(nv string) {
	t.Value = nv
}
