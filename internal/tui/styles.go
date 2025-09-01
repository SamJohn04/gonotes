package tui

import (
	"github.com/SamJohn04/nate/internal/config"
	"github.com/charmbracelet/lipgloss"
)

var (
	baseStyle = lipgloss.NewStyle()

	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00d7af")).
		Bold(true).
		AlignHorizontal(lipgloss.Center)

	centerHorizontalStyle = lipgloss.NewStyle().
		AlignHorizontal(lipgloss.Center)

	centerVerticalStyle = lipgloss.NewStyle().
		AlignVertical(lipgloss.Center)

	footerStyle = lipgloss.NewStyle().
		Background(lipgloss.Color("#353535")).
		Align(lipgloss.Center, lipgloss.Bottom)

	centerBothStyle = lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center)
)

func SetBaseStyle(styleCfg config.StyleConfig) {
	baseStyle = baseStyle.
		Foreground(lipgloss.Color(styleCfg.ForegroundColor)).
		Background(lipgloss.Color(styleCfg.BackgroundColor))

	titleStyle = baseStyle.
		Foreground(lipgloss.Color(styleCfg.HeadingColor)).
		Bold(true).
		AlignHorizontal(lipgloss.Center)

	centerHorizontalStyle = baseStyle.
		AlignHorizontal(lipgloss.Center)

	centerVerticalStyle = baseStyle.
		AlignVertical(lipgloss.Center)

	centerBothStyle = baseStyle.
		Align(lipgloss.Center, lipgloss.Center)
}
