package ui

import "github.com/charmbracelet/lipgloss"

var (
	TaskStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginBottom(0)

	SelectedTaskStyle = lipgloss.NewStyle().
				Background(colorSelected).
				Foreground(colorBlack).
				Padding(0, 1)
)
