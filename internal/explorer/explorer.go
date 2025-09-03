package explorer

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
)

type FileItem struct {
	Name  string
	IsDir bool
}

func ReadDir(path string) []list.Item {
	entries, err := os.ReadDir(path)
	if err != nil {
		return []list.Item{}
	}

	var items []list.Item
	for _, e := range entries {
		items = append(items, FileItem{
			Name: e.Name(),
			IsDir: e.IsDir(),
		})
	}

	return items
}

func (f FileItem) Title() string {
	return f.Name
}

func (f FileItem) Description() string {
	if f.IsDir {
		return "DIR"
	}
	return "file"
}

func (f FileItem) FilterValue() string {
	return f.Name
}
