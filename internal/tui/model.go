package tui

import (
	"path/filepath"

	"github.com/SamJohn04/nate/internal/config"
	"github.com/SamJohn04/nate/internal/explorer"
	"github.com/SamJohn04/nate/internal/files"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type sessionState int

const (
	startupView sessionState = iota
	editorView
	saveView
	explorerView
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

	// Explorer state
	currentDirectory string
	dirList list.Model
}

func InitialModel(filename string, styleCfg config.StyleConfig) model {
	var initialState sessionState

	ti := textarea.New()
	ti.SetValue(files.ReadFile(filename))

	SetBaseStyle(styleCfg)
	
	ti.FocusedStyle.Base = baseStyle
	ti.FocusedStyle.CursorLine = baseStyle

	if filename == "" {
		initialState = startupView
	} else {
		initialState = editorView
		ti.Focus()
	}

	saveTi := textinput.New()
	saveTi.PromptStyle = baseStyle
	saveTi.PlaceholderStyle = baseStyle
	saveTi.TextStyle = baseStyle

	currentDirectory := filepath.Dir(filename)
	dirList := list.New(explorer.ReadDir(currentDirectory), list.NewDefaultDelegate(), 0, 0)

	return model{
		state: initialState,

		width: 80,
		height: 24,

		textarea: ti,
		filename: filename,
		modified: false,

		save: saveTi,

		currentDirectory: currentDirectory,
		dirList: dirList,
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
		case explorerView:
			return m.updateExplorerView(msg)
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
	case explorerView:
		return m.viewExplorerView()
	default:
		panic("Something went horribly wrong.")
	}
}
