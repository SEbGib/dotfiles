package tui

import (
	"fmt"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
			return NewMainModel(), nil
		}
	}

	return m, nil
}

func (m InfoModel) View() string {
	var s strings.Builder

	// Header
	s.WriteString(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Render("📊 Informations Système"))
	s.WriteString("\n\n")

	// System info
	for key, value := range m.systemInfo {
		s.WriteString(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFAA00")).
			Render(fmt.Sprintf("%-15s: ", key)))
		s.WriteString(value)
		s.WriteString("\n")
	}

	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#7D56F4")).
		Render("📁 Chemins importants:"))
	s.WriteString("\n")

	paths := map[string]string{
		"Dotfiles":   "~/.local/share/chezmoi",
		"Zsh Config": "~/.zshrc",
		"Git Config": "~/.gitconfig",
		"Neovim":     "~/.config/nvim",
		"tmux":       "~/.config/tmux",
		"Starship":   "~/.config/starship.toml",
	}

	for key, path := range paths {
		s.WriteString(lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#04B575")).
			Render(fmt.Sprintf("%-15s: ", key)))
		s.WriteString(path)
		s.WriteString("\n")
	}

	// Footer
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Render("• Appuyez sur Entrée ou Échap pour revenir au menu principal"))

	return s.String()
}
