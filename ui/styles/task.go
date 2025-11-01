package styles

import "github.com/charmbracelet/lipgloss"

var (
	taskStyle = lipgloss.NewStyle().
			Padding(0, 1).
			MarginBottom(0)

	selectedTaskStyle = lipgloss.NewStyle().
				Background(colorSelected).
				Foreground(colorBlack).
				Padding(0, 1)
)
