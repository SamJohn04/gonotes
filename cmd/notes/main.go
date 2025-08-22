package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SamJohn04/gonotes/internal/tui"
)

func main() {
	var filename string

	if len(os.Args) == 1 {
		filename = ""
	} else {
		filename = os.Args[1]
	}

	p := tea.NewProgram(tui.InitialModel(filename), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
