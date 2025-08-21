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

	// Editor state
	lines []string
	cursorRow int
	cursorCol int
	scrollY int
	filename string
	modified bool

	command string
}

func InitialModel() model {
	return model{
		state: startupView,

		width: 80,
		height: 24,

		lines: []string{""},
		cursorRow: 0,
		cursorCol: 0,
		scrollY: 0,
		filename: "filename.txt",
		modified: false,

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
