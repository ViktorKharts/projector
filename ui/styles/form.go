package styles

import "github.com/charmbracelet/lipgloss"

var (
	inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorHeader).
			Padding(0, 1).
			Width(50)

	focusedInputStyle = lipgloss.NewStyle().
				BorderForeground(colorSelected).
				BorderStyle(lipgloss.ThickBorder())

	formTitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(colorHeader).
			MarginBottom(1)

	formLabelStyle = lipgloss.NewStyle().
			Foreground(colorText).
			MarginTop(1)
)
