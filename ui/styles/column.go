package ui

import "github.com/charmbracelet/lipgloss"

var (
	ColumnHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorDefault)

	SelectedColumnHeaderStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(colorInProgress)

	ColumnStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorBorder).
			Padding(1).
			Width(30)

	SelectedColumnStyle = lipgloss.NewStyle().
				BorderForeground(colorSelected).
				BorderStyle(lipgloss.ThickBorder())
)

func GetColumnColor(colName string) lipgloss.Color {
	return colorDefault
}

func GetColumnStyle(colName string, isSelected bool) lipgloss.Style {
	color := GetColumnColor(colName)

	if isSelected {
		return SelectedColumnStyle.BorderForeground(color)
	}

	return ColumnStyle.BorderForeground(color)
}
