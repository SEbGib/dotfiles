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

	// Create the enhanced main TUI model with search
	searchableModel := tui.NewSearchableMainModel()

	// Wrap with notifications
	modelWithNotifications := tui.NewWithNotifications(searchableModel)

	// Wrap with help overlay
	shortcuts := tui.GetMainMenuShortcuts()
	finalModel := tui.NewWithHelp(modelWithNotifications, "Aide - Menu Principal", shortcuts)

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
