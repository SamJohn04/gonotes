package tui

import tea "github.com/charmbracelet/bubbletea"

type model struct {
	cursor int
	notes []string
}

func InitialModel() model {
	return model{
		cursor: 0,
		notes: []string{},
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch message := message.(type) {
		case tea.KeyMsg:
			if message.String() == "ctrl+c" {
				return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := "Press ctrl+c to quit\n\n"

	return s
}
