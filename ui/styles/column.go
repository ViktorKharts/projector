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
			Padding(1)

	SelectedColumnStyle = ColumnStyle

	TaskCounterStyle = lipgloss.NewStyle().
				Foreground(colorBlue).
				Align(lipgloss.Center)
)

func GetColumnHeaderStyle(width int, isSelected bool) lipgloss.Style {
	if isSelected {
		return SelectedColumnHeaderStyle.Width(width).Align(lipgloss.Center)
	}

	return ColumnHeaderStyle.Width(width).Align(lipgloss.Center)
}

func GetColumnStyle(isSelected bool) lipgloss.Style {
	if isSelected {
		return SelectedColumnStyle.BorderForeground(colorOrange)
	}

	return ColumnStyle.BorderForeground(colorBlue)
}
