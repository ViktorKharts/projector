package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
	// "github.com/viktorkharts/projector/cmd"
)

func main() {
	// cmd.Execute()
	s, _ := storage.Read()

	p := tea.NewProgram(s)

	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
	}

	if m, ok := m.(models.Storage); ok && m.SelectedProject != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.SelectedProject)
	}
}
