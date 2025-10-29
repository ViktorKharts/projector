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

	nextColumn := b.Project.Columns[b.CurrentColumnIndex+1]
	nextColumn.Tasks = append(nextColumn.Tasks, task)

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

	prevColumn := b.Project.Columns[b.CurrentColumnIndex-1]
	prevColumn.Tasks = append(prevColumn.Tasks, task)

	b.CurrentColumnIndex--
	b.CurrentTaskIndex = len(prevColumn.Tasks) - 1
}
