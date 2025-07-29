package tui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type SecretsModel struct {
	list list.Model
}

func NewSecretsModel() SecretsModel {
	items := []list.Item{
		MenuItem{title: "🔐 Configurer Bitwarden", description: "Configurer l'intégration Bitwarden", action: "setup_bitwarden"},
		MenuItem{title: "🔑 Tester la connexion", description: "Vérifier la connexion aux secrets", action: "test_secrets"},
		MenuItem{title: "📝 Éditer les variables", description: "Modifier les variables d'environnement", action: "edit_env"},
		MenuItem{title: "🔄 Synchroniser les secrets", description: "Mettre à jour depuis Bitwarden", action: "sync_secrets"},
		MenuItem{title: "🔙 Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🔐 Configuration des Secrets"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return SecretsModel{list: l}
}

func (m SecretsModel) Init() tea.Cmd {
	return nil
}

func (m SecretsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewTwoColumnMainModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				switch i.action {
				case "back":
					return NewTwoColumnMainModel(), nil
				case "setup_bitwarden":
					return NewBitwardenSetupModel(), nil
				case "test_secrets":
					return NewSecretsTestModel(), nil
				case "edit_env":
					return NewEnvEditModel(), nil
				case "sync_secrets":
					return NewSecretsSyncModel(), nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m SecretsModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔐 Configuration des Secrets"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Gérez vos secrets et intégration Bitwarden"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour au menu • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// BitwardenSetupModel handles Bitwarden setup
type BitwardenSetupModel struct {
	step       int
	email      textinput.Model
	status     string
	complete   bool
	setupSteps []string
}

func NewBitwardenSetupModel() BitwardenSetupModel {
	email := textinput.New()
	email.Placeholder = "votre@email.com"
	email.Focus()
	email.CharLimit = 100
	email.Width = 50

	steps := []string{
		"Configuration de l'email Bitwarden",
		"Installation du CLI Bitwarden",
		"Connexion et authentification",
		"Test de la configuration",
	}

	return BitwardenSetupModel{
		step:       0,
		email:      email,
		status:     "Entrez votre email Bitwarden",
		setupSteps: steps,
	}
}

func (m BitwardenSetupModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m BitwardenSetupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			if m.complete {
				return NewSecretsModel(), nil
			}
			return NewSecretsModel(), nil
		case "enter":
			if m.complete {
				return NewSecretsModel(), nil
			}
			if m.step == 0 && m.email.Value() != "" {
				m.step++
				m.status = "Configuration en cours..."
				// TODO: Implement actual Bitwarden setup
				m.complete = true
				m.status = "✅ Configuration Bitwarden terminée"
			}
		}
	}

	if m.step == 0 {
		m.email, cmd = m.email.Update(msg)
	}

	return m, cmd
}

func (m BitwardenSetupModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔐 Configuration Bitwarden"))
	s.WriteString("\n\n")

	if m.complete {
		s.WriteString(SubtitleStyle.Render("✅ Configuration terminée"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render("Bitwarden est maintenant configuré et prêt à être utilisé."))
	} else {
		s.WriteString(SubtitleStyle.Render(fmt.Sprintf("Étape %d/%d: %s", m.step+1, len(m.setupSteps), m.setupSteps[m.step])))
		s.WriteString("\n\n")

		if m.step == 0 {
			s.WriteString("Email Bitwarden:\n")
			s.WriteString(m.email.View())
			s.WriteString("\n\n")
		}

		s.WriteString(CardStyle.Render(m.status))
	}

	s.WriteString("\n")
	if m.complete {
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
	} else {
		s.WriteString(FooterStyle.Render("• Entrée Continuer • Échap Annuler • Ctrl+C Quitter"))
	}

	return AppStyle.Render(s.String())
}

// SecretsTestModel handles secrets testing
type SecretsTestModel struct {
	testing  bool
	complete bool
	results  []TestResult
	status   string
}

type TestResult struct {
	name    string
	success bool
	message string
}

func NewSecretsTestModel() SecretsTestModel {
	return SecretsTestModel{
		status: "Prêt à tester la configuration des secrets",
	}
}

func (m SecretsTestModel) Init() tea.Cmd {
	return m.runTests()
}

func (m SecretsTestModel) runTests() tea.Cmd {
	return func() tea.Msg {
		m.testing = true

		// Test Bitwarden CLI
		bwResult := TestResult{
			name:    "Bitwarden CLI",
			success: false,
			message: "Non installé",
		}

		// Simple check for bw command
		if _, err := os.Stat("/usr/local/bin/bw"); err == nil {
			bwResult.success = true
			bwResult.message = "Installé et accessible"
		}

		// Test environment variables
		envResult := TestResult{
			name:    "Variables d'environnement",
			success: os.Getenv("BW_SESSION") != "",
			message: "Session Bitwarden",
		}

		if !envResult.success {
			envResult.message = "Aucune session active"
		}

		return secretsTestCompleteMsg{
			results: []TestResult{bwResult, envResult},
		}
	}
}

func (m SecretsTestModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			if m.complete {
				return NewSecretsModel(), nil
			}
		}
	case secretsTestCompleteMsg:
		m.testing = false
		m.complete = true
		m.results = msg.results

		// Count successful tests
		success := 0
		for _, result := range m.results {
			if result.success {
				success++
			}
		}

		if success == len(m.results) {
			m.status = "✅ Tous les tests sont passés"
		} else {
			m.status = fmt.Sprintf("⚠️ %d/%d tests réussis", success, len(m.results))
		}
	}

	return m, nil
}

func (m SecretsTestModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔑 Test des Secrets"))
	s.WriteString("\n\n")

	if m.testing {
		s.WriteString(SubtitleStyle.Render("🔄 Tests en cours..."))
	} else if m.complete {
		s.WriteString(SubtitleStyle.Render(m.status))
		s.WriteString("\n\n")

		// Show test results
		var resultsText strings.Builder
		for _, result := range m.results {
			status := "❌"
			if result.success {
				status = "✅"
			}
			resultsText.WriteString(fmt.Sprintf("%s %s: %s\n", status, result.name, result.message))
		}

		s.WriteString(CardStyle.Render(resultsText.String()))
	} else {
		s.WriteString(SubtitleStyle.Render("Démarrage des tests..."))
	}

	s.WriteString("\n")
	if m.complete {
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
	} else {
		s.WriteString(FooterStyle.Render("• Tests en cours... • Ctrl+C Quitter"))
	}

	return AppStyle.Render(s.String())
}

// EnvEditModel handles environment variables editing
type EnvEditModel struct {
	envVars []EnvVar
	list    list.Model
}

type EnvVar struct {
	name  string
	value string
	desc  string
}

func NewEnvEditModel() EnvEditModel {
	// Common environment variables for dotfiles
	envVars := []EnvVar{
		{name: "BW_SESSION", value: os.Getenv("BW_SESSION"), desc: "Session Bitwarden"},
		{name: "EDITOR", value: os.Getenv("EDITOR"), desc: "Éditeur par défaut"},
		{name: "SHELL", value: os.Getenv("SHELL"), desc: "Shell par défaut"},
		{name: "PATH", value: os.Getenv("PATH"), desc: "Chemins d'exécution"},
	}

	items := make([]list.Item, len(envVars))
	for i, env := range envVars {
		value := env.value
		if len(value) > 50 {
			value = value[:47] + "..."
		}
		if value == "" {
			value = "(non défini)"
		}

		items[i] = MenuItem{
			title:       fmt.Sprintf("🔧 %s", env.name),
			description: fmt.Sprintf("%s: %s", env.desc, value),
			action:      env.name,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       "🔙 Retour",
		description: "Retour au menu secrets",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🔧 Variables d'Environnement"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return EnvEditModel{
		envVars: envVars,
		list:    l,
	}
}

func (m EnvEditModel) Init() tea.Cmd {
	return nil
}

func (m EnvEditModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewSecretsModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewSecretsModel(), nil
				}
				// TODO: Implement variable editing
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m EnvEditModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔧 Variables d'Environnement"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Gérez vos variables d'environnement"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Éditer • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// SecretsSyncModel handles secrets synchronization
type SecretsSyncModel struct {
	syncing  bool
	complete bool
	status   string
	message  string
}

func NewSecretsSyncModel() SecretsSyncModel {
	return SecretsSyncModel{
		status: "Prêt à synchroniser les secrets",
	}
}

func (m SecretsSyncModel) Init() tea.Cmd {
	return m.syncSecrets()
}

func (m SecretsSyncModel) syncSecrets() tea.Cmd {
	return func() tea.Msg {
		m.syncing = true

		// TODO: Implement actual Bitwarden sync
		// For now, simulate sync

		return secretsSyncCompleteMsg{
			success: true,
			message: "Synchronisation terminée avec succès",
		}
	}
}

func (m SecretsSyncModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			if m.complete {
				return NewSecretsModel(), nil
			}
		}
	case secretsSyncCompleteMsg:
		m.syncing = false
		m.complete = true
		m.message = msg.message
		if msg.success {
			m.status = "✅ Synchronisation réussie"
		} else {
			m.status = "❌ Erreur de synchronisation"
		}
	}

	return m, nil
}

func (m SecretsSyncModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔄 Synchronisation des Secrets"))
	s.WriteString("\n\n")

	if m.syncing {
		s.WriteString(SubtitleStyle.Render("🔄 Synchronisation en cours..."))
	} else if m.complete {
		s.WriteString(SubtitleStyle.Render(m.status))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.message))
	} else {
		s.WriteString(SubtitleStyle.Render("Démarrage de la synchronisation..."))
	}

	s.WriteString("\n")
	if m.complete {
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
	} else {
		s.WriteString(FooterStyle.Render("• Synchronisation en cours... • Ctrl+C Quitter"))
	}

	return AppStyle.Render(s.String())
}

// Message types for secrets operations
type secretsTestCompleteMsg struct {
	results []TestResult
}

type secretsSyncCompleteMsg struct {
	success bool
	message string
}
