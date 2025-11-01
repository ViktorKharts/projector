package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorToDo       = lipgloss.Color("#5F87FF") // Blue
	colorInProgress = lipgloss.Color("#FFD700") // Yellow
	colorDone       = lipgloss.Color("#87FF5F") // Green
	colorSelected   = lipgloss.Color("#FFFF00") // Bright Yellow
	colorHeader     = lipgloss.Color("#00D7FF") // Cyan
	colorText       = lipgloss.Color("#FFFFFF") // White
	colorBorder     = lipgloss.Color("#444444") // Dark Grey
	colorBlack      = lipgloss.Color("#000000") // Black
	colorDefault    = lipgloss.Color("#5F87FF") // Blue

	BaseStyle = lipgloss.NewStyle().Foreground(colorText)

	ProjectHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorHeader).
				Padding(0, 1).
				MarginBottom(1)

	HelpStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorHeader).
			Padding(0, 1).
			Width(50)
)
