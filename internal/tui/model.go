package tui

import tea "github.com/charmbracelet/bubbletea"

type SessionState int

const (
	StartupView SessionState = iota
	EditorView
)

type Model struct {
	State SessionState

	Cursor int
	Command string
}

func InitialModel() Model {
	return Model{
		State: StartupView,

		Cursor: 0,
		Command: "",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m.State {
	case StartupView:
		m, cmd := m.updateStartupView(msg)
		return m, cmd
	case EditorView:
		m, cmd := m.updateEditorView(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) View() string {
	switch m.State {
	case StartupView:
		return m.viewStartupView()
	case EditorView:
		return m.viewEditorView()
	}
	return ""
}
