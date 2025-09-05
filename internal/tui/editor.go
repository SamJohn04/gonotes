package tui

import (
	"strings"

	"github.com/SamJohn04/nate/internal/files"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewEditorView() string {
	base := m.textarea.View()

	quitMessage := footerStyle.
		Width(m.width).
		Height(1).
		Render("(ctrl+z to quit; ctrl+s to save; ctrl+w to save as; ctrl+e for explorer)")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		base,
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
			m.switchToSaveView()
		} else {
			files.WriteFile(m.filename, m.textarea.Value())
		}
		return m, nil
	
	// save by new filename even if it exists
	case "ctrl+w":
		m.switchToSaveView()
		return m, nil

	// explorer
	case "ctrl+e":
		m.switchToExplorerView()
		return m, nil
	
	case "tab":
		m.textarea.InsertString("\t")
		return m, nil
	
	case "enter":
		leading := m.findLeadingWhitespace(m.textarea.Line())
		m.textarea, cmd = m.textarea.Update(msg)
		m.textarea.InsertString(leading)
		return m, cmd
	}

	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func (m *model) switchToEditorView() {
	m.save.Blur()
	m.state = editorView
	m.textarea.Focus()
}

func (m model) findLeadingWhitespace(index int) string {
	line := strings.SplitN(m.textarea.Value(), "\n", index+1)[index]
	leading := ""
	for _, r := range line {
		if r == '\t' || r == ' ' {
			leading += string(r)
		} else {
			break
		}
	}
	return leading
}
