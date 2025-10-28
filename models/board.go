package models

import (
	"slices"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Board struct {
	Project            Project
	CurrentColumnIndex int
	CurrentTaskIndex   int
	Width              int
	Height             int
	Mode               BoardMode
	TitleInput         textinput.Model
	ColumnNameInput    textinput.Model
	FocusedInput       int
}

type BoardMode int

const (
	ViewMode BoardMode = iota
	CreateTaskMode
	EditTaskMode
	CreateColumnMode
	EditColumnMode
)

func (b Board) Init() tea.Cmd {
	return textinput.Blink
}

func (b Board) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.Width = msg.Width
		b.Height = msg.Height

	case tea.KeyMsg:
		switch b.Mode {
		case ViewMode:
			return b.handleViewMode(msg)
		case CreateTaskMode:
			return b.handleCreateTaskMode(msg)
		case EditTaskMode:
			return b.handleEditTaskMode(msg)
		case CreateColumnMode:
			return b.handleCreateColumnMode(msg)
		case EditColumnMode:
			return b.handleEditColumnMode(msg)
		}
	}

	return b, cmd
}

func (b Board) View() string {
	switch b.Mode {
	case ViewMode:
		return b.renderBoard()
	case CreateTaskMode:
		return b.renderTaskForm("create")
	case EditTaskMode:
		return b.renderTaskForm("edit")
	case CreateColumnMode:
		return b.renderColumnForm("create")
	case EditColumnMode:
		return b.renderColumnForm("edit")
	}

	return ""
}

func (b Board) handleViewMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		return b, tea.Quit

	case "h":
		b.CurrentColumnIndex--
		if b.CurrentColumnIndex < 0 {
			b.CurrentColumnIndex = len(b.Project.Columns) - 1
		}
		b.CurrentTaskIndex = 0

	case "l":
		b.CurrentColumnIndex++
		if b.CurrentColumnIndex >= len(b.Project.Columns)-1 {
			b.CurrentColumnIndex = 0
		}
		b.CurrentTaskIndex = 0

	case "k":
		b.CurrentTaskIndex--
		tasksLength := len(b.Project.Columns[b.CurrentColumnIndex].Tasks)
		if b.CurrentTaskIndex < 0 {
			b.CurrentTaskIndex = tasksLength - 1
		}

	case "j":
		b.CurrentTaskIndex++
		tasksLength := len(b.Project.Columns[b.CurrentColumnIndex].Tasks)
		if b.CurrentTaskIndex >= tasksLength-1 {
			b.CurrentTaskIndex = 0
		}

	case "n":
		b.Mode = CreateTaskMode
		b.TitleInput = textinput.New()
		b.TitleInput.Placeholder = "Task title..."
		b.TitleInput.Focus()
		b.TitleInput.Width = 40

		// TODO: add description field

	case "e":
		column := b.Project.Columns[b.CurrentColumnIndex]
		task := column.Tasks[b.CurrentTaskIndex]

		b.Mode = EditColumnMode
		b.TitleInput = textinput.New()
		b.TitleInput.SetValue(task.Title)
		b.TitleInput.Focus()
		b.TitleInput.Width = 40

		// TODO: add description field

	case "x":
		column := &b.Project.Columns[b.CurrentColumnIndex]
		column.Tasks = slices.Delete(column.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)
		if b.CurrentTaskIndex >= len(column.Tasks) && b.CurrentTaskIndex > 0 {
			b.CurrentTaskIndex--
		}

	case "shift+l":
		b.moveTaskToNextColumn()

	case "shift+h":
		b.moveTaskToPrevColumn()

	case "+":
		b.Mode = CreateColumnMode
		b.ColumnNameInput = textinput.New()
		b.ColumnNameInput.Placeholder = "Column name..."
		b.ColumnNameInput.Focus()
		b.ColumnNameInput.Width = 40
	}

	return b, nil
}

func (b Board) handleCreateTaskMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.Mode = ViewMode
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			newTask := Task{
				Id:    uuid.NewString(),
				Title: b.TitleInput.Value(),
			}
			b.Project.Columns[b.CurrentColumnIndex].Tasks = append(
				b.Project.Columns[b.CurrentColumnIndex].Tasks,
				newTask,
			)
		}
		b.Mode = ViewMode
		return b, nil

		// TODO: add handle description
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	return b, cmd
}

func (b Board) handleEditTaskMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.Mode = ViewMode
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			b.Project.Columns[b.CurrentColumnIndex].Tasks[b.CurrentTaskIndex].Title = b.TitleInput.Value()
		}
		b.Mode = ViewMode
		return b, nil

		// TODO: add handle description
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	return b, cmd
}

func (b Board) handleCreateColumnMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.Mode = ViewMode
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			newColumn := Column{
				Id:    uuid.NewString(),
				Name:  b.TitleInput.Value(),
				Tasks: []Task{},
			}
			b.Project.Columns = append(
				b.Project.Columns,
				newColumn,
			)
		}
		b.Mode = ViewMode
		return b, nil
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	return b, cmd
}

func (b Board) handleEditColumnMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.Mode = ViewMode
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			b.Project.Columns[b.CurrentColumnIndex].Name = b.TitleInput.Value()
		}
		b.Mode = ViewMode
		return b, nil
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	return b, cmd
}

func (b Board) renderBoard() string {
	return ""
}

func (b Board) renderTaskForm(t string) string {
	return ""
}

func (b Board) renderColumnForm(t string) string {
	return ""
}
