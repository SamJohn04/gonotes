package tui

import "github.com/charmbracelet/lipgloss"

var (
	baseStyle = lipgloss.NewStyle().
		Padding(1, 2)

	// Startup Styles
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("86")).
		Bold(true).
		Align(lipgloss.Center, lipgloss.Center)
)
