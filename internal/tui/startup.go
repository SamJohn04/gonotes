package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) viewStartupView() string {
	title := titleStyle.
		Width(m.Width).
		Render("GONOTES")
	subTitle := centerStyle.
		Width(m.Width).
		Render("Notes made quick, simple, and easy.")
	
	return centerVertStyle.
		Height(m.Height).
		Render(lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			"",
			subTitle,
			))
}

func (m Model) updateStartupView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit
	
	// change from Startup to Editor
	case "enter":
		m.State = EditorView
	}
	return m, nil
}
