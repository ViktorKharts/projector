package models

import (
	"fmt"
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
	ui "github.com/viktorkharts/projector/ui/styles"
)

type Board struct {
	Project            Project
	CurrentColumnIndex int
	CurrentTaskIndex   int
	Width              int
	Height             int
	Mode               BoardMode
	TitleInput         textinput.Model
	DescriptionInput   textinput.Model
	ColumnNameInput    textinput.Model
	FocusedInput       int
}

type BoardMode int

const (
	ViewMode BoardMode = iota
	ViewTaskMode
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
		case ViewTaskMode:
			return b.handleViewTaskMode(msg)
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
	case ViewTaskMode:
		return b.renderTaskView()
	case CreateTaskMode:
		return b.renderTaskForm("Create New Task")
	case EditTaskMode:
		return b.renderTaskForm("Edit Task")
	case CreateColumnMode:
		return b.renderColumnForm("Create New Column")
	case EditColumnMode:
		return b.renderColumnForm("Rename Column Name")
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
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		b.CurrentTaskIndex--
		tasksLength := len(b.Project.Columns[b.CurrentColumnIndex].Tasks)
		if b.CurrentTaskIndex < 0 {
			b.CurrentTaskIndex = tasksLength - 1
		}

	case "j":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		b.CurrentTaskIndex++
		tasksLength := len(b.Project.Columns[b.CurrentColumnIndex].Tasks)
		if b.CurrentTaskIndex >= tasksLength {
			b.CurrentTaskIndex = 0
		}

	case "n":
		b.Mode = CreateTaskMode
		b.TitleInput = textinput.New()
		b.TitleInput.Placeholder = "Fix all bugs"
		b.TitleInput.Focus()
		b.TitleInput.Width = b.Width

		b.DescriptionInput = textinput.New()
		b.DescriptionInput.Placeholder = "Just fix all bugs, not hard"
		b.DescriptionInput.Width = b.Width

	case "e", "r":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		column := &b.Project.Columns[b.CurrentColumnIndex]
		if len(column.Tasks) == 0 {
			return b, nil
		}

		task := column.Tasks[b.CurrentTaskIndex]

		b.Mode = EditTaskMode
		b.TitleInput = textinput.New()
		b.TitleInput.SetValue(task.Title)
		b.TitleInput.Focus()
		b.TitleInput.Width = b.Width

		b.DescriptionInput = textinput.New()
		b.DescriptionInput.SetValue(task.Description)
		b.DescriptionInput.Width = b.Width

	case "x":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		column := &b.Project.Columns[b.CurrentColumnIndex]
		if len(column.Tasks) == 0 {
			return b, nil
		}
		column.Tasks = slices.Delete(column.Tasks, b.CurrentTaskIndex, b.CurrentTaskIndex+1)
		if b.CurrentTaskIndex >= len(column.Tasks) && b.CurrentTaskIndex > 0 {
			b.CurrentTaskIndex--
		}

	case "v", "enter", " ":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		column := &b.Project.Columns[b.CurrentColumnIndex]
		if len(column.Tasks) == 0 {
			return b, nil
		}
		b.Mode = ViewTaskMode

	case "L":
		b.moveTaskRight()

	case "H":
		b.moveTaskLeft()

	case "K":
		b.moveTaskUp()

	case "J":
		b.moveTaskDown()

	case "+":
		b.Mode = CreateColumnMode
		b.ColumnNameInput = textinput.New()
		b.ColumnNameInput.Placeholder = "Ready for Testing"
		b.ColumnNameInput.Focus()
		b.ColumnNameInput.Width = b.Width

	case "-":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}
		b.Project.Columns = slices.Delete(b.Project.Columns, b.CurrentColumnIndex, b.CurrentColumnIndex+1)
		if b.CurrentColumnIndex >= len(b.Project.Columns) && len(b.Project.Columns) > 0 {
			b.CurrentColumnIndex = len(b.Project.Columns) - 1
		}
		b.CurrentTaskIndex = 0

	case "E", "R":
		if len(b.Project.Columns) == 0 {
			return b, nil
		}

		column := b.Project.Columns[b.CurrentColumnIndex]

		b.Mode = EditColumnMode
		b.ColumnNameInput = textinput.New()
		b.ColumnNameInput.SetValue(column.Name)
		b.ColumnNameInput.Focus()
		b.ColumnNameInput.Width = b.Width

	case "(":
		if b.CurrentColumnIndex-1 < 0 {
			return b, nil
		}

		b.Project.Columns[b.CurrentColumnIndex-1], b.Project.Columns[b.CurrentColumnIndex] =
			b.Project.Columns[b.CurrentColumnIndex], b.Project.Columns[b.CurrentColumnIndex-1]
		b.CurrentColumnIndex--

	case ")":
		if b.CurrentColumnIndex+1 >= len(b.Project.Columns) {
			return b, nil
		}

		b.Project.Columns[b.CurrentColumnIndex], b.Project.Columns[b.CurrentColumnIndex+1] =
			b.Project.Columns[b.CurrentColumnIndex+1], b.Project.Columns[b.CurrentColumnIndex]
		b.CurrentColumnIndex++

	}

	return b, nil
}

func (b Board) handleViewTaskMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.Mode = ViewMode
		return b, nil
	}

	return b, cmd
}

func (b Board) handleCreateTaskMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.FocusedInput = 0
		b.TitleInput.Focus()
		b.DescriptionInput.Blur()
		b.Mode = ViewMode
		return b, nil

	case "tab":
		b.FocusedInput = (b.FocusedInput + 1) % 2
		if b.FocusedInput == 0 {
			b.TitleInput.Focus()
			b.DescriptionInput.Blur()
		} else {
			b.TitleInput.Blur()
			b.DescriptionInput.Focus()
		}
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			tasks := b.Project.Columns[b.CurrentColumnIndex].Tasks
			newTask := Task{
				Id:          uuid.NewString(),
				Title:       b.TitleInput.Value(),
				Description: b.DescriptionInput.Value(),
				Index:       greatestIndex(tasks),
			}
			b.Project.Columns[b.CurrentColumnIndex].Tasks = append(tasks, newTask)
		}
		b.FocusedInput = 0
		b.TitleInput.Focus()
		b.DescriptionInput.Blur()
		b.CurrentTaskIndex = len(b.Project.Columns[b.CurrentColumnIndex].Tasks) - 1
		b.Mode = ViewMode
		return b, nil
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	if b.FocusedInput == 1 {
		b.DescriptionInput, cmd = b.DescriptionInput.Update(msg)
	}

	return b, cmd
}

func (b Board) handleEditTaskMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		b.FocusedInput = 0
		b.TitleInput.Focus()
		b.DescriptionInput.Blur()
		b.Mode = ViewMode
		return b, nil

	case "tab":
		b.FocusedInput = (b.FocusedInput + 1) % 2
		if b.FocusedInput == 0 {
			b.TitleInput.Focus()
			b.DescriptionInput.Blur()
		} else {
			b.TitleInput.Blur()
			b.DescriptionInput.Focus()
		}
		return b, nil

	case "enter":
		if b.TitleInput.Value() != "" {
			task := &b.Project.Columns[b.CurrentColumnIndex].Tasks[b.CurrentTaskIndex]
			task.Title = b.TitleInput.Value()
			task.Description = b.DescriptionInput.Value()
		}
		b.FocusedInput = 0
		b.TitleInput.Focus()
		b.DescriptionInput.Blur()
		b.Mode = ViewMode
		return b, nil
	}

	if b.FocusedInput == 0 {
		b.TitleInput, cmd = b.TitleInput.Update(msg)
	}

	if b.FocusedInput == 1 {
		b.DescriptionInput, cmd = b.DescriptionInput.Update(msg)
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

	header := ui.ProjectHeaderStyle.Render(b.Project.Name)
	s.WriteString("\n" + header + "\n")
	numColumns := len(b.Project.Columns)

	if numColumns == 0 {
		s.WriteString("No columns yet! Press '+' to add one.\n")
		s.WriteString("\n(esc: back to projects)")
		s.WriteString("\n(q: exit projector)")
		return s.String()
	}

	columnWidth := (b.Width / numColumns)
	columns := make([]string, numColumns)

	maxTasks := 0
	allTasks := 0
	for _, col := range b.Project.Columns {
		allTasks += len(col.Tasks)
		if len(col.Tasks) > maxTasks {
			maxTasks = len(col.Tasks)
		}
	}

	for colIdx, col := range b.Project.Columns {
		var columnContent strings.Builder

		columnContent.WriteString(ui.GetColumnHeaderStyle(columnWidth, colIdx == b.CurrentColumnIndex).Render(col.Name) + "\n")
		columnContent.WriteString(ui.TaskCounterStyle.Width(columnWidth).Render(fmt.Sprintf("(%d/%d)", len(col.Tasks), allTasks)) + "\n")
		columnContent.WriteString(ui.SeparatorStyle.Width(columnWidth).Render(strings.Repeat("-", columnWidth/2)) + "\n")

		slices.SortStableFunc(col.Tasks, func(a, b Task) int {
			return a.Index - b.Index
		})

		for taskIdx := range maxTasks {
			if taskIdx < len(col.Tasks) {
				task := col.Tasks[taskIdx]
				if len(task.Title) >= columnWidth {
					task.Title = task.Title[:columnWidth-3] + "..."
				}

				taskDisplay := ui.GetTaskStyle(columnWidth, false).Render("  " + task.Title)
				if colIdx == b.CurrentColumnIndex && taskIdx == b.CurrentTaskIndex {
					taskDisplay = ui.GetTaskStyle(columnWidth, true).Render("▶ " + task.Title)
				}

				columnContent.WriteString(taskDisplay + "\n")
			} else {
				columnContent.WriteString(strings.Repeat(" ", columnWidth) + "\n")
			}
		}

		isSelected := colIdx == b.CurrentColumnIndex
		columnStyle := ui.GetColumnStyle(isSelected)
		columns[colIdx] = columnStyle.Render(columnContent.String())
	}

	s.WriteString(lipgloss.JoinHorizontal(lipgloss.Top, columns...))

	help := ui.HelpStyle.Width(columnWidth * numColumns).Render(
		"h/l: column | j/k: task | H/L/J/K: move task | v/enter: view details\n" +
			"n: new task | e/r: edit | x: delete | -: del column | +: new column | R: rename column\n" +
			"esc: back | q: quit",
	)

	s.WriteString("\n" + help)

	return s.String()
}

func (b Board) renderTaskView() string {
	var s strings.Builder

	task := b.Project.Columns[b.CurrentColumnIndex].Tasks[b.CurrentTaskIndex]

	s.WriteString("\n" + ui.FormLabelStyle.Render("Title:") + "\n")
	s.WriteString(ui.BaseStyle.Render(task.Title) + "\n")
	s.WriteString(ui.FormLabelStyle.Render("Task Description:") + "\n")
	if task.Description != "" {
		s.WriteString(ui.BaseStyle.Render(task.Description) + "\n")
	} else {
		s.WriteString(ui.BaseStyle.Render("(no description)") + "\n")
	}

	help := ui.HelpStyle.Render("\n(esc to return)")
	s.WriteString(help)
	return lipgloss.Place(b.Width, b.Height, lipgloss.Center, lipgloss.Top, s.String())
}

func (b Board) renderTaskForm(t string) string {
	var s strings.Builder

	s.WriteString("\n" + ui.FormHeaderStyle.Render(t) + "\n")

	if b.FocusedInput == 0 {
		s.WriteString(ui.ActiveFormLabelStyle.Render("Task Title:") + "\n")
		s.WriteString(ui.FocusedInputStyle.Render(b.TitleInput.View()) + "\n")
	} else {
		s.WriteString(ui.FormLabelStyle.Render("Task Title:") + "\n")
		s.WriteString(ui.InputStyle.Render(b.TitleInput.View()) + "\n")
	}

	if b.FocusedInput == 1 {
		s.WriteString(ui.ActiveFormLabelStyle.Render("Task Description:") + "\n")
		s.WriteString(ui.FocusedInputStyle.Render(b.DescriptionInput.View()) + "\n")
	} else {
		s.WriteString(ui.FormLabelStyle.Render("Task Description:") + "\n")
		s.WriteString(ui.InputStyle.Render(b.DescriptionInput.View()) + "\n")
	}

	help := ui.HelpStyle.Render("(enter: save | esc: cancel)")
	s.WriteString(help)
	return lipgloss.Place(b.Width, b.Height, lipgloss.Center, lipgloss.Top, s.String())
}

func (b Board) renderColumnForm(t string) string {
	var s strings.Builder

	s.WriteString("\n" + ui.FormHeaderStyle.Render(t) + "\n")
	s.WriteString(ui.ActiveFormLabelStyle.Render("Column Title:") + "\n")
	s.WriteString(ui.FocusedInputStyle.Render(b.ColumnNameInput.View()) + "\n")

	help := ui.HelpStyle.Render("(enter: save | esc: cancel)")
	s.WriteString(help)
	return lipgloss.Place(b.Width, b.Height, lipgloss.Center, lipgloss.Top, s.String())
}
