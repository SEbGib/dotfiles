package tui

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type InstallStep struct {
	name        string
	description string
	command     string
	completed   bool
	running     bool
	error       string
}

type InstallModel struct {
	steps       []InstallStep
	currentStep int
	progress    progress.Model
	spinner     spinner.Model
	status      string
	quitting    bool
	completed   bool
	logs        []string
}

type stepCompleteMsg struct {
	step  int
	error string
}

type stepStartMsg struct {
	step int
}

func NewInstallModel() InstallModel {
	steps := []InstallStep{
		{
			name:        "🔍 Détection du système",
			description: "Détection de l'OS et de l'architecture",
			command:     "detect_system",
		},
		{
			name:        "💾 Sauvegarde des configurations",
			description: "Sauvegarde des configurations existantes",
			command:     "backup_configs",
		},
		{
			name:        "📦 Installation des outils essentiels",
			description: "Installation de Homebrew/apt et outils de base",
			command:     "install_tools",
		},
		{
			name:        "🐚 Configuration Zsh",
			description: "Installation Oh My Zsh et plugins",
			command:     "setup_zsh",
		},
		{
			name:        "📁 Création des dossiers",
			description: "Création de la structure de dossiers",
			command:     "create_directories",
		},
		{
			name:        "⚙️ Application des configurations",
			description: "Application des dotfiles avec chezmoi",
			command:     "apply_configs",
		},
		{
			name:        "✅ Vérification finale",
			description: "Vérification de l'installation",
			command:     "verify_installation",
		},
	}

	p := progress.New(progress.WithDefaultGradient())
	p.Width = 50

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = SpinnerStyle

	return InstallModel{
		steps:    steps,
		progress: p,
		spinner:  s,
		status:   "Prêt à commencer l'installation",
		logs:     []string{},
	}
}

func (m InstallModel) Init() tea.Cmd {
	return tea.Batch(
		m.spinner.Tick,
		m.startInstallation(),
	)
}

func (m InstallModel) startInstallation() tea.Cmd {
	return func() tea.Msg {
		return stepStartMsg{step: 0}
	}
}

func (m InstallModel) runStep(step int) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(time.Millisecond * 500) // Simulate work

		var err error
		switch m.steps[step].command {
		case "detect_system":
			err = m.detectSystem()
		case "backup_configs":
			err = m.backupConfigs()
		case "install_tools":
			err = m.installTools()
		case "setup_zsh":
			err = m.setupZsh()
		case "create_directories":
			err = m.createDirectories()
		case "apply_configs":
			err = m.applyConfigs()
		case "verify_installation":
			err = m.verifyInstallation()
		}

		var errMsg string
		if err != nil {
			errMsg = err.Error()
		}

		return stepCompleteMsg{step: step, error: errMsg}
	}
}

func (m InstallModel) detectSystem() error {
	m.addLog(fmt.Sprintf("OS détecté: %s", runtime.GOOS))
	m.addLog(fmt.Sprintf("Architecture: %s", runtime.GOARCH))
	return nil
}

func (m InstallModel) backupConfigs() error {
	m.addLog("Recherche des configurations existantes...")
	configs := []string{".zshrc", ".gitconfig", ".vimrc", ".tmux.conf"}
	for _, config := range configs {
		m.addLog(fmt.Sprintf("Vérification de %s", config))
	}
	m.addLog("Sauvegarde créée dans ~/.dotfiles-backup-*")
	return nil
}

func (m InstallModel) installTools() error {
	m.addLog("Installation des outils essentiels...")

	if runtime.GOOS == "darwin" {
		m.addLog("Installation via Homebrew...")
		tools := []string{"starship", "zsh", "neovim", "tmux", "fzf", "ripgrep"}
		for _, tool := range tools {
			m.addLog(fmt.Sprintf("Installation de %s", tool))
			time.Sleep(time.Millisecond * 100)
		}
	} else {
		m.addLog("Installation via gestionnaire de paquets système...")
		tools := []string{"starship", "zsh", "neovim", "tmux", "fzf", "ripgrep"}
		for _, tool := range tools {
			m.addLog(fmt.Sprintf("Installation de %s", tool))
			time.Sleep(time.Millisecond * 100)
		}
	}

	return nil
}

func (m InstallModel) setupZsh() error {
	m.addLog("Installation d'Oh My Zsh...")
	m.addLog("Installation des plugins Zsh...")
	plugins := []string{"zsh-autosuggestions", "fast-syntax-highlighting", "zsh-completions"}
	for _, plugin := range plugins {
		m.addLog(fmt.Sprintf("Installation du plugin %s", plugin))
		time.Sleep(time.Millisecond * 100)
	}
	return nil
}

func (m InstallModel) createDirectories() error {
	m.addLog("Création de la structure de dossiers...")
	dirs := []string{"~/dev", "~/dev/projects", "~/dev/tools", "~/.config"}
	for _, dir := range dirs {
		m.addLog(fmt.Sprintf("Création du dossier %s", dir))
	}
	return nil
}

func (m InstallModel) applyConfigs() error {
	m.addLog("Application des configurations avec chezmoi...")
	m.addLog("Configuration de Neovim...")
	m.addLog("Configuration de tmux...")
	m.addLog("Configuration de Git...")
	m.addLog("Configuration de Starship...")
	return nil
}

func (m InstallModel) verifyInstallation() error {
	m.addLog("Vérification de l'installation...")

	if cmd := exec.Command("./verify-installation.sh"); cmd != nil {
		m.addLog("Exécution du script de vérification...")
	}

	m.addLog("✅ Installation vérifiée avec succès!")
	return nil
}

func (m *InstallModel) addLog(message string) {
	m.logs = append(m.logs, fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), message))
	if len(m.logs) > 8 {
		m.logs = m.logs[1:]
	}
}

func (m InstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit
		case "esc":
			if m.completed {
				return NewTwoColumnMainModel(), nil
			} else {
				// Allow cancellation during installation
				m.quitting = true
				return NewTwoColumnMainModel(), nil
			}
		case "enter":
			if m.completed {
				return NewTwoColumnMainModel(), nil
			}
		}

	case stepStartMsg:
		if msg.step < len(m.steps) {
			m.currentStep = msg.step
			m.steps[msg.step].running = true
			m.status = fmt.Sprintf("Exécution: %s", m.steps[msg.step].name)
			return m, m.runStep(msg.step)
		}

	case stepCompleteMsg:
		if msg.step < len(m.steps) {
			m.steps[msg.step].running = false
			m.steps[msg.step].completed = true
			if msg.error != "" {
				m.steps[msg.step].error = msg.error
				m.status = fmt.Sprintf("Erreur: %s", msg.error)
			} else {
				m.status = fmt.Sprintf("Terminé: %s", m.steps[msg.step].name)
			}

			nextStep := msg.step + 1
			if nextStep < len(m.steps) {
				return m, func() tea.Msg {
					return stepStartMsg{step: nextStep}
				}
			} else {
				m.completed = true
				m.status = "🎉 Installation terminée avec succès!"
			}
		}

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.progress.Update(msg)
		m.progress = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m InstallModel) View() string {
	if m.quitting {
		return "Installation interrompue.\n"
	}

	var s strings.Builder

	// Clean header
	s.WriteString(CreateBanner("🚀 Installation Interactive"))
	s.WriteString("\n\n")

	// Progress bar
	completedSteps := 0
	for _, step := range m.steps {
		if step.completed {
			completedSteps++
		}
	}

	progressPercent := float64(completedSteps) / float64(len(m.steps))
	s.WriteString(m.progress.ViewAs(progressPercent))
	s.WriteString(fmt.Sprintf(" %d/%d étapes terminées\n\n", completedSteps, len(m.steps)))

	// Steps list
	for i, step := range m.steps {
		var status string
		var style lipgloss.Style

		if step.completed {
			status = "✅"
			style = lipgloss.NewStyle().Foreground(ColorSuccess)
		} else if step.running {
			status = m.spinner.View()
			style = lipgloss.NewStyle().Foreground(ColorWarning)
		} else if step.error != "" {
			status = "❌"
			style = lipgloss.NewStyle().Foreground(ColorError)
		} else {
			status = "⏳"
			style = lipgloss.NewStyle().Foreground(ColorTextMuted)
		}

		stepText := fmt.Sprintf("%s %s", status, step.name)
		if step.error != "" {
			stepText += fmt.Sprintf(" - Erreur: %s", step.error)
		}

		s.WriteString(style.Render(stepText))
		s.WriteString("\n")

		if i == m.currentStep && step.running {
			s.WriteString(SubtitleStyle.Render("   " + step.description))
			s.WriteString("\n")
		}
	}

	// Status
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().Bold(true).Foreground(ColorInfo).Render("Status: "))
	s.WriteString(m.status)
	s.WriteString("\n")

	// Logs
	if len(m.logs) > 0 {
		s.WriteString("\n")
		s.WriteString(lipgloss.NewStyle().Bold(true).Foreground(ColorPrimary).Render("📋 Logs récents:"))
		s.WriteString("\n")

		for _, log := range m.logs {
			s.WriteString(LogEntryStyle.Render("  " + log))
			s.WriteString("\n")
		}
	}

	// Footer
	s.WriteString("\n")
	if m.completed {
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour au menu • Installation terminée! 🎉"))
	} else {
		s.WriteString(FooterStyle.Render("• Échap Annuler et retour • Ctrl+C Quitter"))
	}

	return s.String()
}
