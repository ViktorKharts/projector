package ui

import "github.com/charmbracelet/lipgloss"

var (
	ColumnHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(colorBlue)

	SelectedColumnHeaderStyle = lipgloss.NewStyle().
					Bold(true).
					Foreground(colorOrange)

	ColumnStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colorGrey).
			Padding(1).
			Width(30)

	SelectedColumnStyle = lipgloss.NewStyle().
				BorderForeground(colorOrange).
				BorderStyle(lipgloss.ThickBorder())
)

func GetColumnStyle(colName string, isSelected bool) lipgloss.Style {
	if isSelected {
		return SelectedColumnStyle.BorderForeground(colorOrange)
	}

	return ColumnStyle.BorderForeground(colorBlue)
}
