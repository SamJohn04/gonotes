package tui

import (
	"path/filepath"

	"github.com/SamJohn04/nate/internal/explorer"
	"github.com/SamJohn04/nate/internal/files"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func (m model) viewExplorerView() string {
	helperText := "Browsing: " + m.currentDirectory
	seenDirList := m.dirList.View()
	return baseStyle.Render(lipgloss.JoinVertical(
		lipgloss.Left,
		helperText,
		seenDirList,
		))
}

func (m model) updateExplorerView(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "esc":
		m.switchToEditorView()
		return m, nil
	case "enter":
		i, ok := m.dirList.SelectedItem().(explorer.FileItem)
		if !ok {
			return m, nil
		} else if i.IsDir {
			newPath := filepath.Join(m.currentDirectory, i.Name)
			m.currentDirectory = newPath
			m.dirList.SetItems(explorer.ReadDir(m.currentDirectory))
			return m, nil
		} else {
			newPath := filepath.Join(m.currentDirectory, i.Name)
			m.filename = newPath
			m.textarea.SetValue(files.ReadFile(newPath))
			m.switchToEditorView()
			return m, nil
		}
	}

	m.dirList.Update(msg)
	return m, nil
}

func (m *model) switchToExplorerView() {
	m.textarea.Blur()
	m.save.Blur()
	m.state = explorerView
}
