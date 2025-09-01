package tui

import (
	"github.com/SamJohn04/nate/internal/files"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewSaveView() string {
	cmdText := "Enter file name:"
	base := m.save.View()
	helpText := "(enter to save; esc to cancel)"

	text := centerBothStyle.
		Height(m.height).
		Width(m.width).
		Render(lipgloss.JoinVertical(
			lipgloss.Left,
			cmdText,
			base,
			"",
			helpText,
			))

	return baseStyle.Render(text)
}

func (m model) updateSaveView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg.String() {
	// go to editor without saving
	case "esc":
		m.switchToEditorView()
		return m, nil
	
	// go to editor after saving
	case "enter":
		m.filename = m.save.Value()
		err := files.WriteNewFile(m.filename, m.textarea.Value())

		if err == nil {
			m.switchToEditorView()
		}

		return m, nil
	}

	m.save, cmd = m.save.Update(msg)
	return m, cmd
}

func (m *model) switchToSaveView() {
	m.textarea.Blur()
	m.state = saveView
	m.save.Focus()
}
