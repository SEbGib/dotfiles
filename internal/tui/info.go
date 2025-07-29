package tui

import (
	"fmt"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type InfoModel struct {
	systemInfo map[string]string
}

func NewInfoModel() InfoModel {
	info := map[string]string{
		"OS":           runtime.GOOS,
		"Architecture": runtime.GOARCH,
		"Go Version":   runtime.Version(),
		"Shell":        "$SHELL",
		"Home":         "$HOME",
		"User":         "$USER",
	}

	return InfoModel{systemInfo: info}
}

func (m InfoModel) Init() tea.Cmd {
	return nil
}

func (m InfoModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			return NewTwoColumnMainModel(), nil
		}
	}

	return m, nil
}

func (m InfoModel) View() string {
	var s strings.Builder

	// Beautiful header
	s.WriteString(CreateBanner("📊 Informations Système"))
	s.WriteString("\n\n")

	// System info in a beautiful table
	var systemContent strings.Builder
	for key, value := range m.systemInfo {
		systemContent.WriteString(fmt.Sprintf("%-15s: %s\n", key, value))
	}

	systemCard := CreateCard("🖥️ Informations Système", systemContent.String())
	s.WriteString(systemCard)
	s.WriteString("\n")

	// Paths info
	paths := map[string]string{
		"Dotfiles":   "~/.local/share/chezmoi",
		"Zsh Config": "~/.zshrc",
		"Git Config": "~/.gitconfig",
		"Neovim":     "~/.config/nvim",
		"tmux":       "~/.config/tmux",
		"Starship":   "~/.config/starship.toml",
	}

	var pathsContent strings.Builder
	for key, path := range paths {
		pathsContent.WriteString(fmt.Sprintf("%-15s: %s\n", key, path))
	}

	pathsCard := CreateCard("📁 Chemins Importants", pathsContent.String())
	s.WriteString(pathsCard)

	// Footer
	s.WriteString("\n")
	footerText := "• Entrée/Échap Retour au menu • Ctrl+C Quitter"
	s.WriteString(FooterStyle.Render(footerText))

	return AppStyle.Render(s.String())
}
