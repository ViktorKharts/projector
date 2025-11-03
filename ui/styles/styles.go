package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorToDo       = lipgloss.Color("#5F87FF") // Blue
	colorInProgress = lipgloss.Color("#FFD700") // Yellow
	colorDone       = lipgloss.Color("#87FF5F") // Green
	colorSelected   = lipgloss.Color("#FFFF00") // Bright Yellow
	colorHeader     = lipgloss.Color("#00D7FF") // Cyan
	colorText       = lipgloss.Color("#FFFFFF") // White
	colorGrey       = lipgloss.Color("#D7D7D7") // Dark Grey
	colorBlack      = lipgloss.Color("#000000")
	colorBlue       = lipgloss.Color("#5F87FF")
	colorOrange     = lipgloss.Color("#FE7743")

	BaseStyle = lipgloss.NewStyle().Foreground(colorText)

	ProjectHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorHeader).
				Padding(0, 1).
				MarginBottom(1)

	HelpStyle = lipgloss.NewStyle().
			Align(lipgloss.Center).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorGrey).
			Foreground(colorGrey).
			Padding(0, 1).
			Width(100)
)
