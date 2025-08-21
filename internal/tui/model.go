package tui

import tea "github.com/charmbracelet/bubbletea"

type sessionState int

const (
	startupView sessionState = iota
	editorView
)

type model struct {
	state sessionState

	width int
	height int

	cursorRow int
	cursorCol int

	command string
}

func InitialModel() model {
	return model{
		state: startupView,

		width: 80,
		height: 24,

		cursorRow: 0,
		cursorCol: 0,

		command: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	
	case tea.KeyMsg:
		switch m.state {
		case startupView:
			m, cmd := m.updateStartupView(msg)
			return m, cmd
		case editorView:
			m, cmd := m.updateEditorView(msg)
			return m, cmd
		}
	}
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case startupView:
		return m.viewStartupView()
	case editorView:
		return m.viewEditorView()
	}
	return ""
}
