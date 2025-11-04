package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/viktorkharts/projector/models"
	"github.com/viktorkharts/projector/storage"
)

func main() {
	s, err := storage.Read()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	p := tea.NewProgram(s, tea.WithAltScreen())

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
