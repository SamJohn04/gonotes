package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewStartupView() string {
	title := titleStyle.
		Width(m.width).
		Render("GONOTES")
	subTitle := centerStyle.
		Width(m.width).
		Render("Notes made quick, simple, and easy.")
	
	return centerVertStyle.
		Height(m.height).
		Render(lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			"",
			subTitle,
			))
}

func (m model) updateStartupView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit
	
	// change from Startup to Editor
	case "enter":
		m.state = editorView
	}
	return m, nil
}
