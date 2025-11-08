package commands

import "github.com/viktorkharts/projector/ui"

type MoveTaskUpDownCommand struct {
	ColumnIndex int
	TaskIndex   int
	Direction   int
}

func (mtc *MoveTaskUpDownCommand) Execute(b *ui.Board) error {
	currentColumn := &b.Project.Columns[mtc.ColumnIndex]
	if len(currentColumn.Tasks) == 0 {
		return nil
	}

	if mtc.TaskIndex == 0 && mtc.Direction < 0 {
		return nil
	}

	if mtc.TaskIndex == len(currentColumn.Tasks)-1 && mtc.Direction > 0 {
		return nil
	}

	tasks := currentColumn.Tasks
	newIdx := mtc.TaskIndex + mtc.Direction

	tasks[mtc.TaskIndex], tasks[newIdx] = tasks[newIdx], tasks[mtc.TaskIndex]

	tempIdx := tasks[newIdx].Index
	tasks[newIdx].Index = tasks[mtc.TaskIndex].Index
	tasks[mtc.TaskIndex].Index = tempIdx

	b.CurrentTaskIndex = newIdx
	return nil
}

func (mtc *MoveTaskUpDownCommand) Undo(b *ui.Board) error {
	currentColumn := &b.Project.Columns[mtc.ColumnIndex]
	tasks := currentColumn.Tasks

	currentPosition := mtc.TaskIndex + mtc.Direction
	originalPosition := mtc.TaskIndex

	tasks[currentPosition], tasks[originalPosition] = tasks[originalPosition], tasks[currentPosition]

	tempIdx := tasks[originalPosition].Index
	tasks[originalPosition].Index = tasks[currentPosition].Index
	tasks[currentPosition].Index = tempIdx

	b.CurrentTaskIndex = mtc.TaskIndex
	return nil
}
