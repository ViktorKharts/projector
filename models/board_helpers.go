package models

import "slices"

func (b *Board) moveTaskToNextColumn() {
	if b.CurrentColumnIndex >= len(b.Project.Columns)-1 {
		return
	}

	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if b.CurrentTaskIndex >= len(currentColumn.Tasks) {
		return
	}

	task := currentColumn.Tasks[b.CurrentTaskIndex]
	currentColumn.Tasks = slices.Delete(currentColumn.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)

	nextColumn := &b.Project.Columns[b.CurrentColumnIndex+1]
	nextColumn.Tasks = append(nextColumn.Tasks, task)
	task.Index = len(nextColumn.Tasks) - 1

	b.CurrentColumnIndex++
	b.CurrentTaskIndex = len(nextColumn.Tasks) - 1
}

func (b *Board) moveTaskToPrevColumn() {
	if b.CurrentColumnIndex <= 0 {
		return
	}

	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if b.CurrentTaskIndex >= len(currentColumn.Tasks) {
		return
	}

	task := currentColumn.Tasks[b.CurrentTaskIndex]
	currentColumn.Tasks = slices.Delete(currentColumn.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)

	prevColumn := &b.Project.Columns[b.CurrentColumnIndex-1]
	prevColumn.Tasks = append(prevColumn.Tasks, task)
	task.Index = len(prevColumn.Tasks) - 1

	b.CurrentColumnIndex--
	b.CurrentTaskIndex = len(prevColumn.Tasks) - 1
}

func (b *Board) moveTaskUp() {
	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if len(currentColumn.Tasks) == 0 {
		return
	}

	if b.CurrentTaskIndex == 0 {
		return
	}

	tasks := currentColumn.Tasks
	newIdx := b.CurrentTaskIndex - 1
	tasks[b.CurrentTaskIndex].Index--
	tasks[newIdx].Index++
	tasks[b.CurrentTaskIndex], tasks[newIdx] = tasks[newIdx], tasks[b.CurrentTaskIndex]
	b.CurrentTaskIndex = newIdx
}

func (b *Board) moveTaskDown() {
	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if len(currentColumn.Tasks) == 0 {
		return
	}

	if b.CurrentTaskIndex == len(currentColumn.Tasks)-1 {
		return
	}

	tasks := currentColumn.Tasks
	newIdx := b.CurrentTaskIndex + 1
	tasks[b.CurrentTaskIndex].Index++
	tasks[newIdx].Index--
	tasks[b.CurrentTaskIndex], tasks[newIdx] = tasks[newIdx], tasks[b.CurrentTaskIndex]
	b.CurrentTaskIndex = newIdx
}
