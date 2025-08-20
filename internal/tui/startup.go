package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m Model) viewStartupView() string {
	title := titleStyle.
		Width(m.Width).
		Render("GONOTES")
	subTitle := centerStyle.
		Width(m.Width).
		Render("Notes made quick, simple, and easy.")
	
	return centerVertStyle.
		Height(m.Height).
		Render(lipgloss.JoinVertical(
			lipgloss.Center,
			title,
			"",
			subTitle,
			))
}

func (m Model) updateStartupView(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
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
