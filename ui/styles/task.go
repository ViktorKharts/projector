package ui

import "github.com/charmbracelet/lipgloss"

var (
	TaskStyle = lipgloss.NewStyle().
			Foreground(colorBlue).
			Align(lipgloss.Center).
			Padding(0, 1).
			MarginBottom(0)

	SelectedTaskStyle = lipgloss.NewStyle().
				Foreground(colorOrange).
				Align(lipgloss.Center).
				Padding(0, 1).
				MarginBottom(0)
)
