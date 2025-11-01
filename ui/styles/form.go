package ui

import "github.com/charmbracelet/lipgloss"

var (
	InputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorHeader).
			Padding(0, 1).
			Width(50)

	FocusedInputStyle = lipgloss.NewStyle().
				BorderForeground(colorSelected).
				BorderStyle(lipgloss.ThickBorder())

	FormTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorHeader).
			MarginBottom(1)

	FormLabelStyle = lipgloss.NewStyle().
			Foreground(colorText).
			MarginTop(1)
)
