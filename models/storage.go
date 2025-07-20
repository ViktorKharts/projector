package models

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Storage struct {
	Cursor           int
	IsWithinSelected bool
	IsNewProject     bool
	SelectedProject  string
	Projects         []Project
	textInput        textinput.Model
}

func (m Storage) Init() tea.Cmd {
	return textinput.Blink
}

func (m Storage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "s", " ":
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
			toDelete := m.Projects[m.Cursor]
			m.Projects = slices.DeleteFunc(m.Projects, func(p Project) bool {
				return p.Id == toDelete.Id
			})
			if m.SelectedProject == toDelete.Name {
				m.SelectedProject = m.Projects[0].Name
			}
			if m.Cursor >= 1 {
				m.Cursor--
			} else {
				m.Cursor = 0
			}

		case "n":
			m.IsNewProject = true

			ti := textinput.New()
			ti.Placeholder = "foo bar"
			ti.Focus()
			ti.CharLimit = 140
			ti.Width = 20
			m.textInput = ti

			return m, nil

		case "esc", "enter":
			m.IsNewProject = false
			p := Project{
				Id:    uuid.NewString(),
				Name:  m.textInput.Value(),
				Tasks: []Task{},
			}
			m.Projects = append(m.Projects, p)
			m.textInput = textinput.Model{}
		}
	}

	if m.IsNewProject {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

func (m Storage) View() string {
	s := strings.Builder{}

	// Create a new Project
	if m.IsNewProject {
		s.WriteString("A new Project has to have a name!\n\n")
		s.WriteString(m.textInput.View())
		s.WriteString("\n\n(esc to return)\n")
		return s.String()
	}

	// Select Project
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
