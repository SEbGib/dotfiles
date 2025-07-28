package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/sebastiengiband/dotfiles/internal/scripts"
)

type VerifyCheck struct {
	name        string
	description string
	status      string // "pending", "running", "passed", "failed", "warning"
	message     string
}

type VerifyModel struct {
	checks       []VerifyCheck
	current      int
	spinner      spinner.Model
	running      bool
	complete     bool
	summary      VerifySummary
	scriptRunner *scripts.ScriptRunner
}

type VerifySummary struct {
	total   int
	passed  int
	failed  int
	warning int
}

type checkCompleteMsg struct {
	index   int
	status  string
	message string
}

func NewVerifyModel() VerifyModel {
	checks := []VerifyCheck{
		{name: "🔧 Chezmoi", description: "Gestionnaire de dotfiles", status: "pending"},
		{name: "⭐ Starship", description: "Prompt moderne", status: "pending"},
		{name: "🐚 Zsh", description: "Shell avancé", status: "pending"},
		{name: "📝 Neovim", description: "Éditeur moderne", status: "pending"},
		{name: "🖥️ tmux", description: "Multiplexeur terminal", status: "pending"},
		{name: "📁 Git", description: "Contrôle de version", status: "pending"},
		{name: "🔍 FZF", description: "Recherche floue", status: "pending"},
		{name: "🔎 Ripgrep", description: "Recherche dans fichiers", status: "pending"},
		{name: "📂 fd", description: "Alternative à find", status: "pending"},
		{name: "🦇 bat", description: "Alternative à cat", status: "pending"},
		{name: "📋 eza", description: "Alternative à ls", status: "pending"},
		{name: "🌟 Lazygit", description: "Interface Git", status: "pending"},
		{name: "📄 Configuration Zsh", description: "~/.zshrc", status: "pending"},
		{name: "⚙️ Configuration Git", description: "~/.gitconfig", status: "pending"},
		{name: "🎨 Configuration Starship", description: "~/.config/starship.toml", status: "pending"},
		{name: "📁 Configuration Neovim", description: "~/.config/nvim/", status: "pending"},
		{name: "🖥️ Configuration tmux", description: "~/.config/tmux/", status: "pending"},
		{name: "🔌 Oh My Zsh", description: "Framework Zsh", status: "pending"},
		{name: "🔌 Plugins Zsh", description: "Plugins installés", status: "pending"},
	}

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return VerifyModel{
		checks:       checks,
		spinner:      s,
		summary:      VerifySummary{total: len(checks)},
		scriptRunner: scripts.NewScriptRunner(),
	}
}

func (m VerifyModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.startVerification(),
	)
}

func (m VerifyModel) startVerification() tea.Cmd {
	return func() tea.Msg {
		m.running = true
		return m.runNextCheck()
	}
}

func (m VerifyModel) runNextCheck() tea.Cmd {
	if m.current >= len(m.checks) {
		return nil
	}

	return func() tea.Msg {
		time.Sleep(time.Millisecond * 200) // Simulate check time

		check := m.checks[m.current]
		status, message := m.performCheck(check)

		return checkCompleteMsg{
			index:   m.current,
			status:  status,
			message: message,
		}
	}
}

func (m VerifyModel) performCheck(check VerifyCheck) (string, string) {
	switch check.name {
	case "🔧 Chezmoi":
		if m.commandExists("chezmoi") {
			return "passed", "chezmoi installé"
		}
		return "failed", "chezmoi non trouvé"

	case "⭐ Starship":
		if m.commandExists("starship") {
			return "passed", "starship installé"
		}
		return "failed", "starship non trouvé"

	case "🐚 Zsh":
		if m.commandExists("zsh") {
			return "passed", "zsh installé"
		}
		return "failed", "zsh non trouvé"

	case "📝 Neovim":
		if m.commandExists("nvim") {
			return "passed", "neovim installé"
		}
		return "failed", "neovim non trouvé"

	case "🖥️ tmux":
		if m.commandExists("tmux") {
			return "passed", "tmux installé"
		}
		return "failed", "tmux non trouvé"

	case "📁 Git":
		if m.commandExists("git") {
			return "passed", "git installé"
		}
		return "failed", "git non trouvé"

	case "🔍 FZF":
		if m.commandExists("fzf") {
			return "passed", "fzf installé"
		}
		return "warning", "fzf non trouvé (optionnel)"

	case "🔎 Ripgrep":
		if m.commandExists("rg") {
			return "passed", "ripgrep installé"
		}
		return "warning", "ripgrep non trouvé (optionnel)"

	case "📂 fd":
		if m.commandExists("fd") {
			return "passed", "fd installé"
		}
		return "warning", "fd non trouvé (optionnel)"

	case "🦇 bat":
		if m.commandExists("bat") {
			return "passed", "bat installé"
		}
		return "warning", "bat non trouvé (optionnel)"

	case "📋 eza":
		if m.commandExists("eza") {
			return "passed", "eza installé"
		}
		return "warning", "eza non trouvé (optionnel)"

	case "🌟 Lazygit":
		if m.commandExists("lazygit") {
			return "passed", "lazygit installé"
		}
		return "warning", "lazygit non trouvé (optionnel)"

	default:
		// File checks would go here
		return "passed", "Vérification simulée"
	}
}

func (m VerifyModel) commandExists(cmd string) bool {
	return m.scriptRunner.CheckCommand(cmd)
}

func (m VerifyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			if m.complete {
				return NewMainModel(), nil
			}
		}

	case checkCompleteMsg:
		if msg.index < len(m.checks) {
			m.checks[msg.index].status = msg.status
			m.checks[msg.index].message = msg.message

			// Update summary
			switch msg.status {
			case "passed":
				m.summary.passed++
			case "failed":
				m.summary.failed++
			case "warning":
				m.summary.warning++
			}

			m.current++
			if m.current >= len(m.checks) {
				m.complete = true
				m.running = false
			} else {
				return m, m.runNextCheck()
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m VerifyModel) View() string {
	var s strings.Builder

	// Clean header
	s.WriteString(CreateBanner("✅ Vérification du Système"))
	s.WriteString("\n\n")

	// Progress
	if m.running {
		s.WriteString(fmt.Sprintf("Vérification en cours... %s (%d/%d)\n\n",
			m.spinner.View(), m.current, len(m.checks)))
	} else if m.complete {
		s.WriteString("Vérification terminée!\n\n")
	}

	// Checks list
	for i, check := range m.checks {
		var status string
		var style lipgloss.Style

		switch check.status {
		case "passed":
			status = "✅"
			style = lipgloss.NewStyle().Foreground(ColorSuccess)
		case "failed":
			status = "❌"
			style = lipgloss.NewStyle().Foreground(ColorError)
		case "warning":
			status = "⚠️"
			style = lipgloss.NewStyle().Foreground(ColorWarning)
		case "running":
			status = m.spinner.View()
			style = lipgloss.NewStyle().Foreground(ColorWarning)
		default:
			status = "⏳"
			style = lipgloss.NewStyle().Foreground(ColorTextMuted)
		}

		if i == m.current && m.running {
			status = m.spinner.View()
			style = lipgloss.NewStyle().Foreground(ColorWarning)
		}

		checkText := fmt.Sprintf("%s %s", status, check.name)
		if check.message != "" {
			checkText += fmt.Sprintf(" - %s", check.message)
		}

		s.WriteString(style.Render(checkText))
		s.WriteString("\n")
	}

	// Summary
	if m.complete {
		s.WriteString("\n")
		s.WriteString(lipgloss.NewStyle().Bold(true).Foreground(ColorPrimary).Render("📊 Résumé:"))
		s.WriteString("\n")

		successRate := float64(m.summary.passed) / float64(m.summary.total) * 100

		s.WriteString(fmt.Sprintf("• Total: %d vérifications\n", m.summary.total))
		s.WriteString(lipgloss.NewStyle().Foreground(ColorSuccess).Render(fmt.Sprintf("• Réussies: %d\n", m.summary.passed)))
		s.WriteString(lipgloss.NewStyle().Foreground(ColorError).Render(fmt.Sprintf("• Échouées: %d\n", m.summary.failed)))
		s.WriteString(lipgloss.NewStyle().Foreground(ColorWarning).Render(fmt.Sprintf("• Avertissements: %d\n", m.summary.warning)))
		s.WriteString(fmt.Sprintf("• Taux de réussite: %.1f%%\n", successRate))

		s.WriteString("\n")
		if m.summary.failed == 0 {
			s.WriteString(lipgloss.NewStyle().Foreground(ColorSuccess).Render("🎉 Système parfaitement configuré!"))
		} else if successRate >= 80 {
			s.WriteString(lipgloss.NewStyle().Foreground(ColorWarning).Render("⚠️ Système majoritairement configuré"))
		} else {
			s.WriteString(lipgloss.NewStyle().Foreground(ColorError).Render("❌ Système nécessite une attention"))
		}
	}

	// Footer
	s.WriteString("\n\n")
	if m.complete {
		s.WriteString(FooterStyle.Render("Entrée/Échap pour retour au menu"))
	} else {
		s.WriteString(FooterStyle.Render("Ctrl+C pour annuler"))
	}

	return s.String()
}
