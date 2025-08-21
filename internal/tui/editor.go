package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) viewEditorView() string {
	s := "Press ctrl+z to quit\n"
	s += m.command
	s += "\n"

	return s
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
