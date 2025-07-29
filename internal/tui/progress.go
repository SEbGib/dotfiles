package tui

import (
	"fmt"
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// ProgressType defines different types of progress indicators
type ProgressType int

const (
	ProgressTypeDeterminate   ProgressType = iota // Known steps/duration
	ProgressTypeIndeterminate                     // Unknown duration with spinner
	ProgressTypeMultiStep                         // Multiple phases with sub-progress
)

// ProgressStep represents a single step in a multi-step operation
type ProgressStep struct {
	Name        string
	Description string
	Status      string // "pending", "running", "completed", "failed"
	Error       string
	Progress    float64 // 0.0 to 1.0 for sub-progress
}

// UnifiedProgressModel provides a consistent progress interface
type UnifiedProgressModel struct {
	Type        ProgressType
	Title       string
	Steps       []ProgressStep
	CurrentStep int

	// Progress components
	ProgressBar progress.Model
	Spinner     spinner.Model

	// State
	Running  bool
	Complete bool
	Success  bool
	Message  string
	Logs     []string

	// Styling
	Width int
}

// Progress messages
type ProgressStartMsg struct {
	Step int
}

type ProgressUpdateMsg struct {
	Step     int
	Progress float64
	Message  string
}

type ProgressCompleteMsg struct {
	Step    int
	Success bool
	Error   string
}

type ProgressFinishedMsg struct {
	Success bool
	Message string
}

// NewUnifiedProgress creates a new unified progress model
func NewUnifiedProgress(progressType ProgressType, title string, steps []ProgressStep) UnifiedProgressModel {
	// Initialize progress bar
	p := progress.New(progress.WithDefaultGradient())
	p.Width = 50

	// Initialize spinner
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(ColorWarning)

	return UnifiedProgressModel{
		Type:        progressType,
		Title:       title,
		Steps:       steps,
		ProgressBar: p,
		Spinner:     s,
		Width:       80,
		Logs:        make([]string, 0),
	}
}

// NewBackupProgress creates progress for backup operations
func NewBackupProgress() UnifiedProgressModel {
	steps := []ProgressStep{
		{Name: " Analyse", Description: "Analyse des fichiers à sauvegarder", Status: "pending"},
		{Name: " Création", Description: "Création du répertoire de sauvegarde", Status: "pending"},
		{Name: " Copie", Description: "Copie des fichiers de configuration", Status: "pending"},
		{Name: " Finalisation", Description: "Finalisation de la sauvegarde", Status: "pending"},
	}

	return NewUnifiedProgress(ProgressTypeMultiStep, " Création de Sauvegarde", steps)
}

// NewVerificationProgress creates progress for verification operations
func NewVerificationProgress() UnifiedProgressModel {
	steps := []ProgressStep{
		{Name: " Outils", Description: "Vérification des outils installés", Status: "pending"},
		{Name: " Configurations", Description: "Vérification des fichiers de config", Status: "pending"},
		{Name: " Plugins", Description: "Vérification des plugins et extensions", Status: "pending"},
		{Name: " Résumé", Description: "Génération du rapport final", Status: "pending"},
	}

	return NewUnifiedProgress(ProgressTypeMultiStep, " Vérification du Système", steps)
}

// NewToolsProgress creates progress for tools operations
func NewToolsProgress(operation string) UnifiedProgressModel {
	var title string
	var steps []ProgressStep

	switch operation {
	case "install":
		title = " Installation d'Outils"
		steps = []ProgressStep{
			{Name: " Détection", Description: "Détection du système et des outils", Status: "pending"},
			{Name: " Téléchargement", Description: "Téléchargement des paquets", Status: "pending"},
			{Name: " Installation", Description: "Installation des outils", Status: "pending"},
			{Name: " Vérification", Description: "Vérification de l'installation", Status: "pending"},
		}
	case "update":
		title = " Mise à Jour d'Outils"
		steps = []ProgressStep{
			{Name: " Analyse", Description: "Analyse des outils installés", Status: "pending"},
			{Name: " Téléchargement", Description: "Téléchargement des mises à jour", Status: "pending"},
			{Name: " Mise à jour", Description: "Application des mises à jour", Status: "pending"},
			{Name: " Vérification", Description: "Vérification des mises à jour", Status: "pending"},
		}
	default:
		title = " Opération sur les Outils"
		steps = []ProgressStep{
			{Name: " Traitement", Description: "Traitement en cours", Status: "pending"},
		}
	}

	return NewUnifiedProgress(ProgressTypeMultiStep, title, steps)
}

func (m UnifiedProgressModel) Init() tea.Cmd {
	return tea.Batch(
		m.Spinner.Tick,
		func() tea.Msg {
			return ProgressStartMsg{Step: 0}
		},
	)
}

func (m UnifiedProgressModel) Update(msg tea.Msg) (UnifiedProgressModel, tea.Cmd) {
	switch msg := msg.(type) {
	case ProgressStartMsg:
		if msg.Step < len(m.Steps) {
			m.CurrentStep = msg.Step
			m.Steps[msg.Step].Status = "running"
			m.Running = true
		}

	case ProgressUpdateMsg:
		if msg.Step < len(m.Steps) {
			m.Steps[msg.Step].Progress = msg.Progress
			if msg.Message != "" {
				m.AddLog(msg.Message)
			}
		}

	case ProgressCompleteMsg:
		if msg.Step < len(m.Steps) {
			if msg.Success {
				m.Steps[msg.Step].Status = "completed"
			} else {
				m.Steps[msg.Step].Status = "failed"
				m.Steps[msg.Step].Error = msg.Error
			}

			// Move to next step
			nextStep := msg.Step + 1
			if nextStep < len(m.Steps) {
				return m, func() tea.Msg {
					return ProgressStartMsg{Step: nextStep}
				}
			} else {
				// All steps complete
				m.Complete = true
				m.Running = false
				m.Success = msg.Success
			}
		}

	case ProgressFinishedMsg:
		m.Complete = true
		m.Running = false
		m.Success = msg.Success
		m.Message = msg.Message

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd

	case progress.FrameMsg:
		progressModel, cmd := m.ProgressBar.Update(msg)
		m.ProgressBar = progressModel.(progress.Model)
		return m, cmd
	}

	return m, nil
}

func (m UnifiedProgressModel) View() string {
	var s lipgloss.Style = lipgloss.NewStyle().Width(m.Width)
	var content string

	// Title
	content += CreateBanner(m.Title) + "\n\n"

	switch m.Type {
	case ProgressTypeDeterminate:
		content += m.viewDeterminate()
	case ProgressTypeIndeterminate:
		content += m.viewIndeterminate()
	case ProgressTypeMultiStep:
		content += m.viewMultiStep()
	}

	// Logs
	if len(m.Logs) > 0 {
		content += "\n" + lipgloss.NewStyle().Bold(true).Foreground(ColorPrimary).Render(" Activité récente:") + "\n"
		for _, log := range m.Logs {
			content += LogEntryStyle.Render("  "+log) + "\n"
		}
	}

	// Footer
	content += "\n"
	if m.Complete {
		if m.Success {
			content += FooterStyle.Render("• Entrée/Échap Retour • Opération terminée avec succès! ")
		} else {
			content += FooterStyle.Render("• Entrée/Échap Retour • Opération échouée ")
		}
	} else {
		content += FooterStyle.Render("• Échap Annuler • Opération en cours...")
	}

	return s.Render(content)
}

func (m UnifiedProgressModel) viewDeterminate() string {
	completedSteps := 0
	for _, step := range m.Steps {
		if step.Status == "completed" {
			completedSteps++
		}
	}

	progressPercent := float64(completedSteps) / float64(len(m.Steps))
	return m.ProgressBar.ViewAs(progressPercent) +
		fmt.Sprintf(" %d/%d étapes terminées\n", completedSteps, len(m.Steps))
}

func (m UnifiedProgressModel) viewIndeterminate() string {
	if m.Running {
		return m.Spinner.View() + " " + m.Message + "\n"
	}
	return m.Message + "\n"
}

func (m UnifiedProgressModel) viewMultiStep() string {
	var content string

	// Overall progress bar
	completedSteps := 0
	for _, step := range m.Steps {
		if step.Status == "completed" {
			completedSteps++
		}
	}

	progressPercent := float64(completedSteps) / float64(len(m.Steps))
	content += m.ProgressBar.ViewAs(progressPercent) +
		fmt.Sprintf(" %d/%d étapes terminées\n\n", completedSteps, len(m.Steps))

	// Steps list
	for i, step := range m.Steps {
		var status string
		var style lipgloss.Style

		switch step.Status {
		case "completed":
			status = ""
			style = lipgloss.NewStyle().Foreground(ColorSuccess)
		case "running":
			status = m.Spinner.View()
			style = lipgloss.NewStyle().Foreground(ColorWarning)
		case "failed":
			status = ""
			style = lipgloss.NewStyle().Foreground(ColorError)
		default:
			status = ""
			style = lipgloss.NewStyle().Foreground(ColorTextMuted)
		}

		stepText := fmt.Sprintf("%s %s", status, step.Name)
		if step.Error != "" {
			stepText += fmt.Sprintf(" - Erreur: %s", step.Error)
		}

		content += style.Render(stepText) + "\n"

		// Show description for current step
		if i == m.CurrentStep && step.Status == "running" {
			content += SubtitleStyle.Render("   "+step.Description) + "\n"

			// Show sub-progress if available
			if step.Progress > 0 {
				subProgress := progress.New(progress.WithDefaultGradient())
				subProgress.Width = 40
				content += "   " + subProgress.ViewAs(step.Progress) + "\n"
			}
		}
	}

	return content
}

// AddLog adds a timestamped log entry
func (m *UnifiedProgressModel) AddLog(message string) {
	timestamp := time.Now().Format("15:04:05")
	m.Logs = append(m.Logs, fmt.Sprintf("[%s] %s", timestamp, message))

	// Keep only last 6 log entries
	if len(m.Logs) > 6 {
		m.Logs = m.Logs[1:]
	}
}

// StartStep marks a step as running
func (m *UnifiedProgressModel) StartStep(step int) {
	if step < len(m.Steps) {
		m.CurrentStep = step
		m.Steps[step].Status = "running"
		m.Running = true
	}
}

// CompleteStep marks a step as completed
func (m *UnifiedProgressModel) CompleteStep(step int, success bool, errorMsg string) {
	if step < len(m.Steps) {
		if success {
			m.Steps[step].Status = "completed"
		} else {
			m.Steps[step].Status = "failed"
			m.Steps[step].Error = errorMsg
		}
	}
}

// UpdateStepProgress updates the progress of a specific step
func (m *UnifiedProgressModel) UpdateStepProgress(step int, progress float64, message string) {
	if step < len(m.Steps) {
		m.Steps[step].Progress = progress
		if message != "" {
			m.AddLog(message)
		}
	}
}

// IsComplete returns whether all steps are complete
func (m UnifiedProgressModel) IsComplete() bool {
	for _, step := range m.Steps {
		if step.Status != "completed" && step.Status != "failed" {
			return false
		}
	}
	return true
}

// IsSuccessful returns whether all steps completed successfully
func (m UnifiedProgressModel) IsSuccessful() bool {
	for _, step := range m.Steps {
		if step.Status == "failed" {
			return false
		}
	}
	return m.IsComplete()
}
