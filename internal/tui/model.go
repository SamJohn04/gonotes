package tui

import (
	"github.com/SamJohn04/nate/internal/config"
	"github.com/SamJohn04/nate/internal/files"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	save textinput.Model
}

func InitialModel(filename string, styleCfg config.StyleConfig) model {
	var initialState sessionState

	ti := textarea.New()
	ti.SetValue(files.ReadFile(filename))
	
	ti.FocusedStyle.Base = lipgloss.NewStyle().
		Foreground(lipgloss.Color(styleCfg.ForegroundColor)).
		Background(lipgloss.Color(styleCfg.BackgroundColor))
	ti.FocusedStyle.CursorLine = lipgloss.NewStyle().
		Foreground(lipgloss.Color(styleCfg.ForegroundColor)).
		Background(lipgloss.Color(styleCfg.BackgroundColor))

	if filename == "" {
		initialState = startupView
	} else {
		initialState = editorView
		ti.Focus()
	}

	saveTi := textinput.New()

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
		m.textarea.SetHeight(m.height-1)
	
	case tea.KeyMsg:
		switch m.state {
		case startupView:
			return m.updateStartupView(msg)
		case editorView:
			return m.updateEditorView(msg)
		case saveView:
			return m.updateSaveView(msg)
		default:
			panic("Something went horribly wrong.")
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
	case saveView:
		return m.viewSaveView()
	default:
		panic("Something went horribly wrong.")
	}
}
