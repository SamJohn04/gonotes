package tui

import "github.com/charmbracelet/lipgloss"

var (
	baseStyle = lipgloss.NewStyle().
		Padding(1, 2).
		Align(lipgloss.Left, lipgloss.Top)

	centerStyle = lipgloss.NewStyle().
		Align(lipgloss.Center)

	centerVertStyle = lipgloss.NewStyle().
		AlignVertical(lipgloss.Center)

	topVertStyle = lipgloss.NewStyle().
		AlignVertical(lipgloss.Top)

	bottomVertStyle = lipgloss.NewStyle().
		AlignVertical(lipgloss.Bottom)

	// Startup Styles
	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00d7af")).
		Bold(true).
		Align(lipgloss.Center)

	// Editor Styles
	footerStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#353535")).
		Align(lipgloss.Center, lipgloss.Bottom)

	// Save Styles
	bothCenterStyle = lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center)
)
