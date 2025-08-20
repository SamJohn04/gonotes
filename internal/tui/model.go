package tui

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	cursor int
	notes []string
	command string
}

func InitialModel() model {
	return model{
		cursor: 0,
		notes: []string{},
		command: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		cmd := m.updateKeyMessage(msg)
		if cmd != nil {
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Press ctrl+c to quit\n"

	if m.command != "" {
		s += m.command
	}

	s += "\n"

	return s
}

func (m *model) updateKeyMessage(msg tea.KeyMsg) (tea.Cmd) {
	switch msg.String() {
	// exit the program
	case "ctrl+c":
		return tea.Quit

	case "enter":
		m.command = ""
	
	case "backspace":
		m.command = m.command[:len(m.command)-1]

	default:
		m.command += msg.String()
	}

	return nil
}
