package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewEditorView() string {
	base := m.textarea.View()

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
	var cmd tea.Cmd

	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit

	case "ctrl+s":
		m.state = saveView
		return m, nil
	
	case "tab":
		m.textarea.InsertString("\t")
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}
