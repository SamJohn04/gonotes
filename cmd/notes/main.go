package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/SamJohn04/nate/internal/config"
	"github.com/SamJohn04/nate/internal/tui"
)

func main() {
	var filename string

	if len(os.Args) == 1 {
		filename = ""
	} else {
		filename = os.Args[1]
	}

	styleCfg := config.Load()

	p := tea.NewProgram(tui.InitialModel(filename, styleCfg), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running TUI:", err)
		os.Exit(1)
	}
}
