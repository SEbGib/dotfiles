package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/sebastiengiband/dotfiles/internal/scripts"
	"github.com/sebastiengiband/dotfiles/internal/tui"
)

func main() {
	// Setup logging
	log.SetLevel(log.DebugLevel)

	// Initialize cache system
	scriptRunner := scripts.NewScriptRunner()
	tui.InitializeCache(scriptRunner)

	// Create the two-column main TUI model with search and notifications
	twoColumnModel := tui.NewTwoColumnMainModel()

	// Wrap with help overlay
	shortcuts := tui.GetMainMenuShortcuts()
	finalModel := tui.NewWithHelp(twoColumnModel, "Aide - Menu Principal", shortcuts)

	// Create the Bubble Tea program
	p := tea.NewProgram(
		finalModel,
		tea.WithAltScreen(),
		tea.WithMouseCellMotion(),
	)

	// Run the program
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running TUI: %v\n", err)
		os.Exit(1)
	}
}
