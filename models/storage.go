package models

import (
	"slices"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Storage struct {
	SelectedProject string
	Projects        []Project
	Cursor          int
}

func (m Storage) Init() tea.Cmd {
	return nil
}

func (m Storage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "s", " ":
			m.SelectedProject = m.Projects[m.Cursor].Name
			return m, nil

		case "down", "j":
			m.Cursor++
			if m.Cursor >= len(m.Projects) {
				m.Cursor = 0
			}

		case "up", "k":
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
		}
	}

	return m, nil
}

func (m Storage) View() string {
	s := strings.Builder{}
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
