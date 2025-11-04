package models

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	ui "github.com/viktorkharts/projector/ui/styles"
)

type Storage struct {
	Cursor          int
	Width           int
	Height          int
	SelectedProject string
	Projects        []Project
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

func (m Storage) Init() tea.Cmd {
	return textinput.Blink
}

func (m Storage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m Storage) View() string {
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

func (m Storage) handleProjectsViewMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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

	case " ":
		if len(m.Projects) == 0 {
			return m, nil
		}
		m.ShowingBoard = true
		m.SelectedProject = m.Projects[m.Cursor].Name
		m.CurrentBoard = Board{
			Project: m.Projects[m.Cursor],
			Width:   m.Width,
			Height:  m.Height,
		}
	}

	return m, nil
}

func (m Storage) handleCreateProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	case "esc":
		m.Mode = ProjectsViewMode
		return m, nil

	case "enter":
		if m.ProjectInput.Value() != "" {
			p := Project{
				Id:   uuid.NewString(),
				Name: m.ProjectInput.Value(),
				Columns: []Column{
					{Id: uuid.NewString(), Name: "To Do", Tasks: []Task{}},
					{Id: uuid.NewString(), Name: "In Progress", Tasks: []Task{}},
					{Id: uuid.NewString(), Name: "Done", Tasks: []Task{}},
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

func (m Storage) handleEditProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
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

func (m Storage) renderProjectsView() string {
	var s strings.Builder

	if len(m.Projects) == 0 {
		help := ui.HelpStyle.Render("No projects yet! Press 'n' to create your first project.\n" +
			"\n(n: new project | q: quit)")
		s.WriteString("\n" + help)
		return s.String()
	}

	s.WriteString(ui.ProjectHeaderStyle.Render("Projector Wecomes You. These are your projects:") + "\n")

	for i, v := range m.Projects {
		var line string

		cursor := "   "
		if m.Cursor == i {
			cursor = " ▶ "
		}

		selection := "○ "
		if m.SelectedProject == m.Projects[i].Name {
			selection = "● "
		}

		line = cursor + selection + v.Name
		s.WriteString(line + "\n")
	}
	s.WriteString("\n\n")

	help := ui.HelpStyle.Render("(press q to quit)")
	s.WriteString("\n" + help + "\n")
	return s.String()
}

func (m Storage) renderCreateNewProjectView() string {
	var s strings.Builder

	s.WriteString("Provide a name for a new Project\n\n")

	s.WriteString(m.ProjectInput.View() + "\n")

	s.WriteString("\n(enter to save)\n")
	s.WriteString("(esc to return)\n")

	return s.String()
}

func (m Storage) renderEditProjectView() string {
	var s strings.Builder

	project := m.Projects[m.Cursor]

	s.WriteString("Change the name for the project '" + project.Name + "'" + "\n\n")

	s.WriteString(m.ProjectInput.View() + "\n")

	s.WriteString("\n(enter to save)\n")
	s.WriteString("(esc to return)\n")

	return s.String()
}
