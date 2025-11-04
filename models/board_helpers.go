package models

import "slices"

func (b *Board) moveTaskRight() {
	if b.CurrentColumnIndex >= len(b.Project.Columns)-1 {
		return
	}

	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if b.CurrentTaskIndex >= len(currentColumn.Tasks) {
		return
	}

	task := currentColumn.Tasks[b.CurrentTaskIndex]
	currentColumn.Tasks = slices.Delete(currentColumn.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)

	leftColumn := &b.Project.Columns[b.CurrentColumnIndex+1]
	task.Index = greatestIndex(leftColumn.Tasks)
	leftColumn.Tasks = append(leftColumn.Tasks, task)

	b.CurrentColumnIndex++
	b.CurrentTaskIndex = len(leftColumn.Tasks) - 1
}

func (b *Board) moveTaskLeft() {
	if b.CurrentColumnIndex <= 0 {
		return
	}

	currentColumn := &b.Project.Columns[b.CurrentColumnIndex]
	if b.CurrentTaskIndex >= len(currentColumn.Tasks) {
		return
	}

	task := currentColumn.Tasks[b.CurrentTaskIndex]
	currentColumn.Tasks = slices.Delete(currentColumn.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)

	rightColumn := &b.Project.Columns[b.CurrentColumnIndex-1]
	task.Index = greatestIndex(rightColumn.Tasks)
	rightColumn.Tasks = append(rightColumn.Tasks, task)

	b.CurrentColumnIndex--
	b.CurrentTaskIndex = len(rightColumn.Tasks) - 1
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

	tasks[b.CurrentTaskIndex], tasks[newIdx] = tasks[newIdx], tasks[b.CurrentTaskIndex]

	tempIdx := tasks[newIdx].Index
	tasks[newIdx].Index = tasks[b.CurrentTaskIndex].Index
	tasks[b.CurrentTaskIndex].Index = tempIdx

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

	tasks[b.CurrentTaskIndex], tasks[newIdx] = tasks[newIdx], tasks[b.CurrentTaskIndex]

	tempIdx := tasks[newIdx].Index
	tasks[newIdx].Index = tasks[b.CurrentTaskIndex].Index
	tasks[b.CurrentTaskIndex].Index = tempIdx

	b.CurrentTaskIndex = newIdx
}

func greatestIndex(tasks []Task) int {
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
