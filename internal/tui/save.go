package tui

import (
	"github.com/SamJohn04/gonotes/internal/files"
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) viewSaveView() string {
	base := m.save.View()
	return bothCenterStyle.
		Height(m.height).
		Width(m.width).
		Render(base)
}

func (m model) updateSaveView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	// go to editor without saving
	case "ctrl+z":
		m.save.Blur()
		m.state = editorView
		m.textarea.Focus()

		return m, nil
	
	// go to editor after saving
	case "ctrl+s":
		m.filename = m.save.Value()
		files.WriteFile(m.filename, m.textarea.Value())

		m.save.Blur()
		m.state = editorView
		m.textarea.Focus()

		return m, nil
	}

	m.save, cmd = m.save.Update(msg)
	return m, cmd
}
