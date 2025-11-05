package ui

import "github.com/charmbracelet/lipgloss"

var (
	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBlue).
			Padding(0, 1).
			Width(100)

	FocusedInputStyle = InputStyle.
				UnsetForeground().
				BorderForeground(colorOrange)

	FormHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorHeader).
			MarginBottom(1)

	FormLabelStyle = lipgloss.NewStyle().
			Width(100).
			Bold(true).
			Foreground(colorBlue)

	ActiveFormLabelStyle = FormLabelStyle.
				UnsetForeground().
				Foreground(colorOrange)
)
