package models

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Storage struct {
	Cursor          int
	Widht           int
	Height          int
	SelectedProject string
	Projects        []Project
	ProjectInput    textinput.Model
	CurrentBoard    Board
	ShowingBoard    bool
	Mode            ProjectsMode
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

	// TODO: make this work
	// if m.SelectedProject != "" {
	// 	m.ShowingBoard = true
	//
	// 	var idx int
	// 	for i, p := range m.Projects {
	// 		if p.Name == m.SelectedProject {
	// 			idx = i
	// 			break
	// 		}
	// 	}
	//
	// 	m.CurrentBoard = Board{
	// 		Project: m.Projects[idx],
	// 		Width:   80,
	// 		Height:  24,
	// 	}
	// }

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
		m.Widht = msg.Width
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

	// if m.IsNewProject || m.IsProjectEdit {
	// 	m.TextBubble, cmd = m.TextBubble.Update(msg)
	// }

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

	// case "n":
	// 	m.IsNewProject = true
	//
	// 	ti := textinput.New()
	// 	ti.Placeholder = "foo bar"
	// 	ti.Focus()
	// 	ti.CharLimit = 140
	// 	ti.Width = 20
	// 	m.TextBubble = ti
	//
	// 	return m, nil

	// case "r":
	// 	m.IsProjectEdit = true
	//
	// 	ti := textinput.New()
	// 	ti.SetValue(m.Projects[m.Cursor].Name)
	// 	ti.Focus()
	// 	ti.CharLimit = 140
	// 	ti.Width = 20
	// 	m.TextBubble = ti
	//
	// 	return m, nil

	// case "esc", "enter":
	// 	if m.IsNewProject {
	// 		m.IsNewProject = false
	// 		p := Project{
	// 			Id:   uuid.NewString(),
	// 			Name: m.TextBubble.Value(),
	// 			Columns: []Column{
	// 				{Id: uuid.NewString(), Name: "To Do", Tasks: []Task{}},
	// 				{Id: uuid.NewString(), Name: "In Progress", Tasks: []Task{}},
	// 				{Id: uuid.NewString(), Name: "Done", Tasks: []Task{}},
	// 			},
	// 		}
	// 		m.Projects = append(m.Projects, p)
	// 		m.TextBubble = textinput.Model{}
	// 	}
	// 	if m.IsProjectEdit {
	// 		m.IsProjectEdit = false
	// 		m.Projects[m.Cursor].Name = m.TextBubble.Value()
	// 		m.TextBubble = textinput.Model{}
	// 	}

	case " ":
		m.ShowingBoard = true
		m.SelectedProject = m.Projects[m.Cursor].Name
		m.CurrentBoard = Board{
			Project: m.Projects[m.Cursor],
			Width:   80,
			Height:  24,
		}
		return m, nil
	}

	return m, nil
}

func (m Storage) handleCreateProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Storage) handleEditProjectMode(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m Storage) renderProjectsView() string {
	var s strings.Builder

	s.WriteString("These are all the projects you have:\n\n")

	for i, v := range m.Projects {
		if m.Cursor == i {
			s.WriteString(" > ")
		} else if m.Cursor != i {
			s.WriteString("   ")
		}

		if m.SelectedProject == m.Projects[i].Name {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}

		s.WriteString(v.Name)
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

func (m Storage) renderCreateNewProjectView() string {
	var s strings.Builder

	s.WriteString("Provide a name for a new Project.\n\n")
	s.WriteString(m.ProjectInput.View() + "\n")
	s.WriteString("\n\n(enter to save)\n")
	s.WriteString("(esc to return)\n")

	return s.String()
}

func (m Storage) renderEditProjectView() string {
	var s strings.Builder

	project := m.Projects[m.Cursor]

	s.WriteString("Change the name for the project '" + project.Name + "'" + "\n\n")
	s.WriteString(m.ProjectInput.View() + "\n")
	s.WriteString("\n\n(enter to save)\n")
	s.WriteString("(esc to return)\n")

	return s.String()
}
