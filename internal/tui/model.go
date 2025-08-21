package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/textarea"
)

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
	textarea textarea.Model
	filename string
	modified bool
}

func InitialModel() model {
	ti := textarea.New()

	return model{
		state: startupView,

		width: 80,
		height: 24,

		textarea: ti,
		// lines: []string{""},
		// cursorRow: 0,
		// cursorCol: 0,
		// scrollY: 0,
		filename: "",
		modified: false,
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

		m.textarea.SetWidth(m.width)
		m.textarea.SetHeight(m.height-2)
	
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
