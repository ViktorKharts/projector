package models

import (
	"fmt"
	"slices"
	"strings"

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
		return b.renderTaskForm("Create Task")
	case EditTaskMode:
		return b.renderTaskForm("Edit Task")
	case CreateColumnMode:
		return b.renderColumnForm("Create Column")
	case EditColumnMode:
		return b.renderColumnForm("Rename Column")
	}

	return ""
}

func (b Board) handleViewMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "q":
		return b, tea.Quit

	case "h":
		b.CurrentColumnIndex--
		if b.CurrentColumnIndex < 0 {
			b.CurrentColumnIndex = len(b.Project.Columns) - 1
		}
		b.CurrentTaskIndex = 0

	case "l":
		b.CurrentColumnIndex++
		if b.CurrentColumnIndex >= len(b.Project.Columns) {
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
		if b.CurrentTaskIndex >= tasksLength {
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

		b.Mode = EditTaskMode
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

	case "L":
		b.moveTaskToNextColumn()

	case "H":
		b.moveTaskToPrevColumn()

	case "+":
		b.Mode = CreateColumnMode
		b.ColumnNameInput = textinput.New()
		b.ColumnNameInput.Placeholder = "Column name..."
		b.ColumnNameInput.Focus()
		b.ColumnNameInput.Width = 40

	case "-":
		b.Project.Columns = slices.Delete(b.Project.Columns, b.CurrentColumnIndex, b.CurrentColumnIndex+1)
		if b.CurrentColumnIndex > 0 {
			b.CurrentColumnIndex--
		}

	case "E":
		column := b.Project.Columns[b.CurrentColumnIndex]

		b.Mode = EditColumnMode
		b.ColumnNameInput = textinput.New()
		b.ColumnNameInput.SetValue(column.Name)
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
		if b.ColumnNameInput.Value() != "" {
			newColumn := Column{
				Id:    uuid.NewString(),
				Name:  b.ColumnNameInput.Value(),
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
		b.ColumnNameInput, cmd = b.ColumnNameInput.Update(msg)
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
		if b.ColumnNameInput.Value() != "" {
			b.Project.Columns[b.CurrentColumnIndex].Name = b.ColumnNameInput.Value()
		}
		b.Mode = ViewMode
		return b, nil
	}

	if b.FocusedInput == 0 {
		b.ColumnNameInput, cmd = b.ColumnNameInput.Update(msg)
	}

	return b, cmd
}

func (b Board) renderBoard() string {
	var s strings.Builder

	s.WriteString(fmt.Sprintf("Projects: %s\n\n", b.Project.Name))
	numColumns := len(b.Project.Columns)

	if numColumns == 0 {
		s.WriteString("No columns yet! Press '+' to add one.\n")
		s.WriteString("\n(esc: back to projects)")
		s.WriteString("\n(q: exit projector)")
		return s.String()
	}

	columnWidth := (b.Width / numColumns)

	for i, col := range b.Project.Columns {
		colHeader := col.Name
		if i == b.CurrentColumnIndex {
			colHeader = "> " + colHeader + " <"
		}
		s.WriteString(fmt.Sprintf("%-*s", columnWidth, colHeader))
	}
	s.WriteString("\n")

	for range numColumns {
		s.WriteString(strings.Repeat("-", columnWidth))
	}
	s.WriteString("\n")

	maxTasks := 0
	for _, col := range b.Project.Columns {
		if len(col.Tasks) > maxTasks {
			maxTasks = len(col.Tasks)
		}
	}

	for taskRow := range maxTasks {
		for colIndex, col := range b.Project.Columns {
			if taskRow < len(col.Tasks) {
				task := col.Tasks[taskRow]
				taskDisplay := task.Title

				if colIndex == b.CurrentColumnIndex && taskRow == b.CurrentTaskIndex {
					taskDisplay = "* " + taskDisplay + " *"
				}

				if len(taskDisplay) > columnWidth-2 {
					taskDisplay = taskDisplay[:columnWidth-5] + "..."
				}

				s.WriteString(fmt.Sprintf("%-*s", columnWidth, taskDisplay))
			} else {
				s.WriteString(strings.Repeat(" ", columnWidth))
			}
		}
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString("h/l: change column | j/k: change task\n")
	s.WriteString("n: new task | e: edit | x: delete | shift+l: move right | shift+h: move left | +: new column\n")
	s.WriteString("esc: back to projects\n")
	s.WriteString("q: quit\n")

	return s.String()
}

func (b Board) renderTaskForm(t string) string {
	var s strings.Builder

	s.WriteString(t + "\n\n")

	s.WriteString("Task Title:\n")
	if b.FocusedInput == 0 {
		s.WriteString(b.TitleInput.View() + " <- focused\n")
	} else {
		s.WriteString(b.TitleInput.View() + "\n")
	}

	s.WriteString("\n(enter: save | esc: cancel)")

	return s.String()
}

func (b Board) renderColumnForm(t string) string {
	var s strings.Builder

	s.WriteString(t + "\n\n")

	s.WriteString("Column Title:\n")
	if b.FocusedInput == 0 {
		s.WriteString(b.ColumnNameInput.View() + "\n")
	} else {
		s.WriteString(b.ColumnNameInput.View() + "\n")
	}

	s.WriteString("\n(enter: save | esc: cancel)")

	return s.String()
}
