package models

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Board struct {
	Project            Project
	CurrentColumnIndex int
	CurrentTaskIndex   int
	Width              int
	Height             int
	Mode               BoardMode
	TitleInput         textinput.Model
	ColumnNameInpput   textinput.Model
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
