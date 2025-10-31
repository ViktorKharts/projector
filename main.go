package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
	// "github.com/viktorkharts/projector/cmd"
)

func main() {
	// cmd.Execute()
	s, err := initializeStorage()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(s)

	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	if m, ok := m.(models.Storage); ok {
		if err = storage.Write(m); err != nil {
			fmt.Printf("Error: %v", err)
			os.Exit(1)
		}
	}
}

func initializeStorage() (models.Storage, error) {
	s, err := storage.Read()
	if err != nil {
		return models.Storage{}, err
	}

	ti := textinput.New()
	ti.Placeholder = "foo bar"
	ti.Focus()
	ti.CharLimit = 140
	ti.Width = 20
	s.TextBubble = ti

	items := make([]list.Item, len(s.Projects))
	for i, p := range s.Projects {
		items[i] = p
	}

	delegate := list.NewDefaultDelegate()

	li := list.New(items, delegate, 0, 0)
	li.Title = "Projector"
	// li.AdditionalFullHelpKeys = func() []key.Binding {
	// 	return []key.Binding{
	// 		listKeys.toggleSpinner,
	// 		listKeys.insertItem,
	// 		listKeys.toggleTitleBar,
	// 		listKeys.toggleStatusBar,
	// 		listKeys.togglePagination,
	// 		listKeys.toggleHelpMenu,
	// 	}
	// }
	s.ListBubble = li

	return s, nil
}
