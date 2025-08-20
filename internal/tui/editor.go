package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) viewEditorView() string {
	s := "Press ctrl+z to quit\n"
	s += m.Command
	s += "\n"

	return s
}

func (m Model) updateEditorView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, cmd := m.updateEditorKeyMsg(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) updateEditorKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit
	
	case "enter":
		m.Command += "\n"
	
	case "backspace":
		m.Command = m.Command[:len(m.Command)-1]

	default:
		m.Command += msg.String()
	}
	return m, nil
}
