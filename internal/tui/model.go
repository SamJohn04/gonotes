package tui

import (
	"github.com/SamJohn04/gonotes/internal/files"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	startupView sessionState = iota
	editorView
	saveView
)

type model struct {
	state sessionState

	width int
	height int

	// Editor state
	textarea textarea.Model
	filename string
	modified bool

	// Save state
	save textarea.Model
}

func InitialModel(filename string) model {
	var initialState sessionState

	ti := textarea.New()
	ti.SetValue(files.ReadFile(filename))

	if ti.Value() == "" {
		initialState = startupView
	} else {
		initialState = editorView
	}

	saveTi := textarea.New()

	return model{
		state: initialState,

		width: 80,
		height: 24,

		textarea: ti,
		filename: filename,
		modified: false,

		save: saveTi,
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
