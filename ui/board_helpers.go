package ui

import (
	"slices"

	"github.com/viktorkharts/projector/models"
)

func (b *Board) moveTaskLeftRight(direction int) {
	if b.CurrentColumnIndex >= len(b.Project.Columns)-1 && direction > 0 {
		return
	}

	if b.CurrentColumnIndex <= 0 && direction < 0 {
		return
	}

	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if b.CurrentTaskIndex >= len(currentColumn.Tasks) {
		return
	}

	task := currentColumn.Tasks[b.CurrentTaskIndex]
	currentColumn.Tasks = slices.Delete(currentColumn.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)

	nextColumn := &b.Project.Columns[b.CurrentColumnIndex+direction]
	task.Index = greatestIndex(nextColumn.Tasks)
	nextColumn.Tasks = append(nextColumn.Tasks, task)

	b.CurrentColumnIndex += direction
	b.CurrentTaskIndex = len(nextColumn.Tasks) - 1
}

func (b *Board) moveTaskUpDown(direction int) {
	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if len(currentColumn.Tasks) == 0 {
		return
	}

	if b.CurrentTaskIndex == 0 && direction < 0 {
		return
	}

	if b.CurrentTaskIndex == len(currentColumn.Tasks)-1 && direction > 0 {
		return
	}

	tasks := currentColumn.Tasks
	newIdx := b.CurrentTaskIndex + direction

	tasks[b.CurrentTaskIndex], tasks[newIdx] = tasks[newIdx], tasks[b.CurrentTaskIndex]

	tempIdx := tasks[newIdx].Index
	tasks[newIdx].Index = tasks[b.CurrentTaskIndex].Index
	tasks[b.CurrentTaskIndex].Index = tempIdx

	b.CurrentTaskIndex = newIdx
}

func greatestIndex(tasks []models.Task) int {
	if len(tasks) == 0 {
		return 0
	}

	greatest := tasks[0].Index
	for _, t := range tasks {
		if t.Index > greatest {
			greatest = t.Index
		}
	}

	return greatest + 1
}
