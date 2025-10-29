package models

import (
	"slices"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
)

type Storage struct {
	Cursor          int
	IsNewProject    bool
	IsProjectEdit   bool
	SelectedProject string
	Projects        []Project
	ListBubble      list.Model
	TextBubble      textinput.Model
	CurrentBoard    Board
	ShowingBoard    bool
	ViewMode        BoardMode
}

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
		boardModel, cmd := m.CurrentBoard.Update(msg)
		m.CurrentBoard = boardModel.(Board)

		if msg, ok := msg.(tea.KeyMsg); ok && msg.String() == "esc" {
			m.ShowingBoard = false
			m.Projects[m.Cursor] = m.CurrentBoard.Project
		}

		return m, cmd
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
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
			m.IsNewProject = true

			ti := textinput.New()
			ti.Placeholder = "foo bar"
			ti.Focus()
			ti.CharLimit = 140
			ti.Width = 20
			m.TextBubble = ti

			return m, nil

		case "r":
			m.IsProjectEdit = true

			ti := textinput.New()
			ti.SetValue(m.Projects[m.Cursor].Name)
			ti.Focus()
			ti.CharLimit = 140
			ti.Width = 20
			m.TextBubble = ti

			return m, nil

		case "esc", "enter":
			if m.IsNewProject {
				m.IsNewProject = false
				p := Project{
					Id:   uuid.NewString(),
					Name: m.TextBubble.Value(),
					Columns: []Column{
						{Id: uuid.NewString(), Name: "To Do", Tasks: []Task{}},
						{Id: uuid.NewString(), Name: "In Progress", Tasks: []Task{}},
						{Id: uuid.NewString(), Name: "Done", Tasks: []Task{}},
					},
				}
				m.Projects = append(m.Projects, p)
				m.TextBubble = textinput.Model{}
			}
			if m.IsProjectEdit {
				m.IsProjectEdit = false
				m.Projects[m.Cursor].Name = m.TextBubble.Value()
				m.TextBubble = textinput.Model{}
			}

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
	}

	if m.IsNewProject || m.IsProjectEdit {
		m.TextBubble, cmd = m.TextBubble.Update(msg)
	}

	return m, cmd
}

func (m Storage) View() string {
	if m.ShowingBoard {
		return m.CurrentBoard.View()
	}

	s := strings.Builder{}

	// Create a new Project
	if m.IsNewProject {
		s.WriteString("A new Project has to have a name!\n\n")
		s.WriteString(m.TextBubble.View())
		s.WriteString("\n\n(esc to return)\n")
		return s.String()
	}

	// Edit a Project
	if m.IsProjectEdit {
		s.WriteString("Here, you can provide a new name for the Project!\n\n")
		s.WriteString(m.TextBubble.View())
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
