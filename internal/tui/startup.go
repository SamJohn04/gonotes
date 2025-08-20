package tui

import tea "github.com/charmbracelet/bubbletea"

func (m Model) viewStartupView() string {
	s := "Welcome, and Good Day!\n\n"

	return s
}

func (m Model) updateStartupView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		m, cmd := m.updateStartupKeyMsg(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) updateStartupKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+z":
		return m, tea.Quit
	
	// change from StartupView to EditorView
	case "enter":
		m.State = EditorView
	}
	return m, nil
}
