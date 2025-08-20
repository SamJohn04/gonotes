package tui

import "github.com/charmbracelet/lipgloss"

var (
	centerVertStyle = lipgloss.NewStyle().
		AlignVertical(lipgloss.Center)

	baseStyle = lipgloss.NewStyle().
		Padding(1, 2)

	centerStyle = lipgloss.NewStyle().
		Align(lipgloss.Center)

	// Startup Styles
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Align(lipgloss.Center)
)
