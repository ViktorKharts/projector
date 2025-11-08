package ui

import (
	"slices"
	"strings"

	"github.com/viktorkharts/projector/commands"
	"github.com/viktorkharts/projector/models"
	ui "github.com/viktorkharts/projector/ui/styles"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/google/uuid"
)

type Main struct {
	Cursor          int
	Width           int
	Height          int
	SelectedProject string
	Projects        []models.Project
	ProjectInput    textinput.Model
	CurrentBoard    Board
	ShowingBoard    bool
	Mode            ProjectsMode
	FocusedInput    int
}

type ProjectsMode int

const (
	ProjectsViewMode ProjectsMode = iota
	CreateProjectMode
	EditProjectMode
)

func (m Main) Init() tea.Cmd {
	return textinput.Blink
}

func (m Main) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if m.ShowingBoard {
		if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "esc" && m.CurrentBoard.Mode == ViewMode {
			m.ShowingBoard = false
			m.Projects[m.Cursor] = m.CurrentBoard.Project
			return m, nil
		}

		boardModel, cmd := m.CurrentBoard.Update(msg)
		m.CurrentBoard = boardModel.(Board)
		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height

	case tea.KeyMsg:
		switch m.Mode {
		case ProjectsViewMode:
			return m.handleProjectsViewMode(msg)
		case CreateProjectMode:
			return m.handleCreateProjectMode(msg)
		case EditProjectMode:
			return m.handleEditProjectMode(msg)
		}
	}

	return m, cmd
}

func (m Main) View() string {
	if m.ShowingBoard {
		return m.CurrentBoard.View()
	}

	switch m.Mode {
	case ProjectsViewMode:
		return m.renderProjectsView()
	case CreateProjectMode:
		return m.renderCreateNewProjectView()
	case EditProjectMode:
		return m.renderEditProjectView()
	}

	return ""
}

func (m Main) handleProjectsViewMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit

	case "s":
		m.SelectedProject = m.Projects[m.Cursor].Name
		return m, nil

	case "j":
		m.Cursor++
		if m.Cursor >= len(m.Projects) {
			m.Cursor = 0
		}

	case "k":
		m.Cursor--
		if m.Cursor < 0 {
			m.Cursor = len(m.Projects) - 1
		}

	case "x":
		if len(m.Projects) == 0 {
			return m, nil
		}
		m.Projects = slices.Delete(m.Projects, m.Cursor, m.Cursor+1)
		if m.Cursor > 0 {
			m.Cursor--
		}

	case "n":
		m.Mode = CreateProjectMode
		m.ProjectInput = textinput.New()
		m.ProjectInput.Placeholder = "Project name..."
		m.ProjectInput.Focus()
		m.ProjectInput.Width = 40
		m.ProjectInput.CharLimit = 140

	case "r", "e":
		if len(m.Projects) == 0 {
			return m, nil
		}
		m.Mode = EditProjectMode
		m.ProjectInput = textinput.New()
		m.ProjectInput.SetValue(m.Projects[m.Cursor].Name)
		m.ProjectInput.Focus()
		m.ProjectInput.Width = 40
		m.ProjectInput.CharLimit = 140

	case " ", "enter":
		if len(m.Projects) == 0 {
			return m, nil
		}
		m.ShowingBoard = true
		m.SelectedProject = m.Projects[m.Cursor].Name
		m.CurrentBoard = Board{
			Project: m.Projects[m.Cursor],
			Width:   m.Width,
			Height:  m.Height,
			History: commands.NewCommandBoardHistory(),
		}
	}

	return m, nil
}

func (m Main) handleCreateProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Mode = ProjectsViewMode
		return m, nil

	case "enter":
		if m.ProjectInput.Value() != "" {
			p := models.Project{
				Id:   uuid.NewString(),
				Name: m.ProjectInput.Value(),
				Columns: []models.Column{
					{Id: uuid.NewString(), Name: "To Do", Tasks: []models.Task{}},
					{Id: uuid.NewString(), Name: "In Progress", Tasks: []models.Task{}},
					{Id: uuid.NewString(), Name: "Done", Tasks: []models.Task{}},
				},
			}
			m.Projects = append(m.Projects, p)
			m.SelectedProject = m.ProjectInput.Value()
		}
		m.Mode = ProjectsViewMode
		return m, nil
	}

	if m.FocusedInput == 0 {
		m.ProjectInput, cmd = m.ProjectInput.Update(msg)
	}

	return m, cmd
}

func (m Main) handleEditProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Mode = ProjectsViewMode
		return m, nil

	case "enter":
		if m.ProjectInput.Value() != "" {
			p := &m.Projects[m.Cursor]
			p.Name = m.ProjectInput.Value()
		}
		m.Mode = ProjectsViewMode
		m.SelectedProject = m.ProjectInput.Value()
		return m, nil
	}

	if m.FocusedInput == 0 {
		m.ProjectInput, cmd = m.ProjectInput.Update(msg)
	}

	return m, cmd
}

func (m Main) renderProjectsView() string {
	var s strings.Builder

	if len(m.Projects) == 0 {
		help := ui.HelpStyle.Render("No projects yet! Press 'n' to create your first project.\n" +
			"\n(n: new project | q: quit)")
		s.WriteString("\n" + help)
		return s.String()
	}

	s.WriteString(ui.ProjectHeaderStyle.Render("Projector Wecomes You. These are your projects:") + "\n")

	var leftAlignedBlock strings.Builder
	for i, v := range m.Projects {
		var line string

		cursor := "○ "
		line = ui.ProjectDefault.Render(cursor + v.Name)
		if m.Cursor == i {
			cursor = "● "
			line = ui.ProjectSelected.Render(cursor + v.Name)
		}

		leftAlignedBlock.WriteString(line + "\n")
	}
	s.WriteString(ui.ProjectsContainer.Render(leftAlignedBlock.String()) + "\n\n")

	help := ui.HelpStyle.Render("(press q to quit)")
	s.WriteString("\n" + help + "\n")

	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Center, s.String())
}

func (m Main) renderCreateNewProjectView() string {
	var s strings.Builder

	s.WriteString(ui.FormHeaderStyle.Render("Provide a name for a new Project") + "\n")

	s.WriteString(ui.ActiveFormLabelStyle.Render("Project Title:") + "\n")
	s.WriteString(ui.FocusedInputStyle.Render(m.ProjectInput.View()) + "\n")

	help := ui.HelpStyle.Render("(enter: save | esc: return)")
	s.WriteString(help)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Top, s.String())
}

func (m Main) renderEditProjectView() string {
	var s strings.Builder

	project := m.Projects[m.Cursor]

	s.WriteString("Change name for " + ui.FormHeaderStyle.Render(project.Name) + "\n\n")
	s.WriteString(ui.ActiveFormLabelStyle.Render("New Project Title:") + "\n")
	s.WriteString(ui.FocusedInputStyle.Render(m.ProjectInput.View()) + "\n")

	help := ui.HelpStyle.Render("(enter: save | esc: return)")
	s.WriteString(help)
	return lipgloss.Place(m.Width, m.Height, lipgloss.Center, lipgloss.Top, s.String())
}
