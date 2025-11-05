package ui

import "github.com/charmbracelet/lipgloss"

var (
	ProjectsHeader = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorHeader)

	ProjectDefault = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorBlue)

	ProjectSelected = ProjectDefault.
			UnsetForeground().
			Foreground(colorOrange)

	ProjectsContainer = lipgloss.NewStyle().
				Align(lipgloss.Left)
)
