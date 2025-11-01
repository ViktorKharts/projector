package styles

import "github.com/charmbracelet/lipgloss"

var (
	columnHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Padding(0, 1).
				Align(lipgloss.Center)

	selectedColumnHeaderStyle = lipgloss.NewStyle().
					Bold(true).
					Underline(true).
					Foreground(colorSelected).
					Align(lipgloss.Center)

	columnStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1).
			Width(30)

	selectedColumnStyle = lipgloss.NewStyle().
				BorderForeground(colorSelected).
				BorderStyle(lipgloss.ThickBorder())
)

func GetColumnColor(colName string) lipgloss.Color {
	switch colName {
	case "To Do":
		return colorToDo
	case "In Progress":
		return colorInProgress
	case "Done":
		return colorDone
	default:
		return colorDefault
	}
}

func GetColumnStyle(colName string, isSelected bool) lipgloss.Style {
	color := GetColumnColor(colName)

	if isSelected {
		return selectedColumnStyle.BorderForeground(color)
	}

	return columnStyle.BorderForeground(color)
}
