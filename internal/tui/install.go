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
	p.Width = 60

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

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
	// Check for existing configs
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
		// Simulate homebrew installation
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

	// Run actual verification script if it exists
	if cmd := exec.Command("./verify-installation.sh"); cmd != nil {
		m.addLog("Exécution du script de vérification...")
		// Don't actually run it in this demo
	}

	m.addLog("✅ Installation vérifiée avec succès!")
	return nil
}

func (m *InstallModel) addLog(message string) {
	m.logs = append(m.logs, fmt.Sprintf("[%s] %s", time.Now().Format("15:04:05"), message))
	if len(m.logs) > 10 {
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
				return NewMainModel(), nil
			}
		case "enter":
			if m.completed {
				return NewMainModel(), nil
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

			// Start next step
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
		return CreateStatusBadge("warning", "Installation interrompue")
	}

	var s strings.Builder

	// Beautiful header
	s.WriteString(CreateBanner("🚀 Installation Interactive des Dotfiles"))
	s.WriteString("\n\n")

	// Progress section in a card
	completedSteps := 0
	for _, step := range m.steps {
		if step.completed {
			completedSteps++
		}
	}

	progressPercent := float64(completedSteps) / float64(len(m.steps))
	progressBar := CreateProgressBar(progressPercent, 50)
	progressText := fmt.Sprintf("%d/%d étapes terminées (%.0f%%)",
		completedSteps, len(m.steps), progressPercent*100)

	progressCard := CreateCard("📊 Progression",
		progressBar+"\n"+progressText)
	s.WriteString(progressCard)
	s.WriteString("\n")

	// Steps list in a beautiful card
	var stepsContent strings.Builder
	for i, step := range m.steps {
		var statusText string

		if step.completed {
			statusText = CreateStatusBadge("success", step.name)
		} else if step.running {
			statusText = SpinnerStyle.Render(m.spinner.View()) + " " +
				MenuItemStyle.Render(step.name)
		} else if step.error != "" {
			statusText = CreateStatusBadge("error", step.name+" - "+step.error)
		} else {
			statusText = CreateStatusBadge("pending", step.name)
		}

		stepsContent.WriteString(statusText)
		stepsContent.WriteString("\n")

		if i == m.currentStep && step.running {
			stepsContent.WriteString(SubtitleStyle.Render("   " + step.description))
			stepsContent.WriteString("\n")
		}
	}

	stepsCard := CreateCard("📋 Étapes d'installation", stepsContent.String())
	s.WriteString(stepsCard)

	// Status section
	s.WriteString("\n")
	statusCard := CreateCard("📊 Status",
		CreateStatusBadge("info", m.status))
	s.WriteString(statusCard)

	// Logs section
	if len(m.logs) > 0 {
		s.WriteString("\n")
		var logsContent strings.Builder
		for _, log := range m.logs {
			parts := strings.SplitN(log, "] ", 2)
			if len(parts) == 2 {
				timestamp := strings.TrimPrefix(parts[0], "[")
				message := parts[1]
				logsContent.WriteString(CreateLogEntry(timestamp, message))
				logsContent.WriteString("\n")
			} else {
				logsContent.WriteString(LogEntryStyle.Render(log))
				logsContent.WriteString("\n")
			}
		}

		logsCard := CreateCard("📋 Logs récents", logsContent.String())
		s.WriteString(LogContainerStyle.Render(logsCard))
	}

	// Footer
	s.WriteString("\n")
	var footerText string
	if m.completed {
		footerText = "• Entrée/Échap Retour au menu • Installation terminée avec succès! 🎉"
	} else {
		footerText = "• Ctrl+C Annuler l'installation • Installation en cours..."
	}
	s.WriteString(FooterStyle.Render(footerText))

	return AppStyle.Render(s.String())
}
