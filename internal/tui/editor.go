package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewEditorView() string {
	base := baseStyle.
		Width(m.width).
		Height(m.height-2).
		Render(m.command)

	quitMessage := footerStyle.
		Width(m.width).
		Height(1).
		Render("(ctrl+z to quit)")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		base,
		"",
		quitMessage,
		)
}

func (m model) updateEditorView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit
	
	case "enter":
		m.command += "\n"
	
	case "backspace":
		if m.command != "" {
			m.command = m.command[:len(m.command)-1]
		}

	default:
		m.command += msg.String()
	}
	return m, nil
}
