package tui

import (
	"github.com/SamJohn04/gonotes/internal/files"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewEditorView() string {
	base := m.textarea.View()

	quitMessage := footerStyle.
		Width(m.width).
		Height(1).
		Render("(ctrl+z to quit; ctrl+s to save; ctrl+w to save as)")

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

	// save
	case "ctrl+s":
		if m.filename == "" {
			m.textarea.Blur()
			m.state = saveView
			m.save.Focus()
		} else {
			files.WriteFile(m.filename, m.textarea.Value())
		}
		return m, nil
	
	// save by new filename even if it exists
	case "ctrl+w":
		m.textarea.Blur()
		m.state = saveView
		m.save.Focus()

		return m, nil
	
	case "tab":
		m.textarea.InsertString("\t")
		return m, nil
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}
