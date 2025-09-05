package tui

import (
	"fmt"
	"io"
	"strings"

	"github.com/SamJohn04/nate/internal/explorer"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type explorerDelegate struct{}

func (d explorerDelegate) Height() int {
	return 1
}

func (d explorerDelegate) Spacing() int {
	return 0
}

func (d explorerDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd {
	return nil
}

func (d explorerDelegate) Render(w io.Writer, m list.Model, index int, item list.Item) {
	i, ok := item.(explorer.FileItem)
	if !ok {
		return
	}

	line := i.Name
	if i.IsDir {
		line += "/"
	}

	if index == m.Index() {
		line = "> " + line
	} else {
		line = "  " + line
	}

	width := m.Width()
	if len(line) < width {
		line = line + strings.Repeat(" ", width-len(line))
	}

	fmt.Fprint(w, baseStyle.Render(line))
}
