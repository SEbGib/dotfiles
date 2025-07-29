package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebastiengiband/dotfiles/internal/scripts"
)

type BackupModel struct {
	list list.Model
}

func NewBackupModel() BackupModel {
	items := []list.Item{
		MenuItem{title: "💾 Créer une sauvegarde", description: "Sauvegarder les configurations actuelles", action: "create_backup"},
		MenuItem{title: "📋 Lister les sauvegardes", description: "Voir toutes les sauvegardes disponibles", action: "list_backups"},
		MenuItem{title: "🔄 Restaurer une sauvegarde", description: "Restaurer depuis une sauvegarde", action: "restore_backup"},
		MenuItem{title: "🗑️ Supprimer une sauvegarde", description: "Supprimer une sauvegarde ancienne", action: "delete_backup"},
		MenuItem{title: "🔙 Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "💾 Sauvegarde & Restauration"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return BackupModel{list: l}
}

func (m BackupModel) Init() tea.Cmd {
	return nil
}

func (m BackupModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				case "create_backup":
					return NewBackupCreateModel(), nil
				case "list_backups":
					return NewBackupListModel(), nil
				case "restore_backup":
					return NewBackupRestoreModel(), nil
				case "delete_backup":
					return NewBackupDeleteModel(), nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m BackupModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("💾 Sauvegarde & Restauration"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Gérez vos sauvegardes de configuration"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour au menu • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// BackupCreateModel handles backup creation
type BackupCreateModel struct {
	status       string
	message      string
	creating     bool
	complete     bool
	scriptRunner *scripts.ScriptRunner
}

func NewBackupCreateModel() BackupCreateModel {
	return BackupCreateModel{
		status:       "Prêt à créer une sauvegarde",
		scriptRunner: scripts.NewScriptRunner(),
	}
}

func (m BackupCreateModel) Init() tea.Cmd {
	return m.createBackup()
}

func (m BackupCreateModel) createBackup() tea.Cmd {
	return func() tea.Msg {
		m.creating = true

		// Create backup directory with timestamp
		homeDir := os.Getenv("HOME")
		if homeDir == "" {
			return backupCompleteMsg{success: false, message: "Impossible de déterminer le répertoire home"}
		}

		timestamp := time.Now().Format("2006-01-02_15-04-05")
		backupDir := filepath.Join(homeDir, fmt.Sprintf(".dotfiles-backup-%s", timestamp))

		// Create backup directory
		if err := os.MkdirAll(backupDir, 0755); err != nil {
			return backupCompleteMsg{success: false, message: fmt.Sprintf("Erreur création répertoire: %v", err)}
		}

		// Files to backup
		filesToBackup := []string{
			".zshrc", ".gitconfig", ".aliases", ".tmux.conf",
			".config/starship.toml", ".config/nvim", ".config/tmux",
		}

		backedUp := 0
		for _, file := range filesToBackup {
			srcPath := filepath.Join(homeDir, file)
			if _, err := os.Stat(srcPath); err == nil {
				dstPath := filepath.Join(backupDir, file)

				// Create directory if needed
				if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err == nil {
					// Copy file (simplified - in real implementation would use proper copy)
					if data, err := os.ReadFile(srcPath); err == nil {
						if err := os.WriteFile(dstPath, data, 0644); err == nil {
							backedUp++
						}
					}
				}
			}
		}

		return backupCompleteMsg{
			success: true,
			message: fmt.Sprintf("Sauvegarde créée: %s (%d fichiers)", backupDir, backedUp),
		}
	}
}

func (m BackupCreateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			if m.complete {
				return NewBackupModel(), nil
			}
		}
	case backupCompleteMsg:
		m.creating = false
		m.complete = true
		m.message = msg.message
		if msg.success {
			m.status = "✅ Sauvegarde terminée"
		} else {
			m.status = "❌ Erreur lors de la sauvegarde"
		}
	}
	return m, nil
}

func (m BackupCreateModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("💾 Création de Sauvegarde"))
	s.WriteString("\n\n")

	if m.creating {
		s.WriteString(SubtitleStyle.Render("🔄 Création en cours..."))
	} else if m.complete {
		s.WriteString(SubtitleStyle.Render(m.status))
	}
	s.WriteString("\n\n")

	if m.message != "" {
		s.WriteString(CardStyle.Render(m.message))
		s.WriteString("\n")
	}

	if m.complete {
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
	} else {
		s.WriteString(FooterStyle.Render("• Création en cours... • Ctrl+C Quitter"))
	}

	return AppStyle.Render(s.String())
}

// BackupListModel handles listing backups
type BackupListModel struct {
	backups      []string
	list         list.Model
	scriptRunner *scripts.ScriptRunner
}

func NewBackupListModel() BackupListModel {
	scriptRunner := scripts.NewScriptRunner()
	backups, _ := scriptRunner.ListBackups()

	items := make([]list.Item, len(backups))
	for i, backup := range backups {
		items[i] = MenuItem{
			title:       fmt.Sprintf("📁 %s", backup),
			description: "Sauvegarde disponible",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       "🔙 Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "📋 Liste des Sauvegardes"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return BackupListModel{
		backups:      backups,
		list:         l,
		scriptRunner: scriptRunner,
	}
}

func (m BackupListModel) Init() tea.Cmd {
	return nil
}

func (m BackupListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewBackupModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewBackupModel(), nil
				}
				// Show backup details (simplified)
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m BackupListModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("📋 Liste des Sauvegardes"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde trouvée"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render("💡 Créez votre première sauvegarde depuis le menu principal"))
	} else {
		s.WriteString(SubtitleStyle.Render(fmt.Sprintf("%d sauvegarde(s) disponible(s)", len(m.backups))))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// BackupRestoreModel handles backup restoration
type BackupRestoreModel struct {
	backups      []string
	list         list.Model
	scriptRunner *scripts.ScriptRunner
}

func NewBackupRestoreModel() BackupRestoreModel {
	scriptRunner := scripts.NewScriptRunner()
	backups, _ := scriptRunner.ListBackups()

	items := make([]list.Item, len(backups))
	for i, backup := range backups {
		items[i] = MenuItem{
			title:       fmt.Sprintf("🔄 %s", backup),
			description: "Cliquez pour restaurer cette sauvegarde",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       "🔙 Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🔄 Restaurer une Sauvegarde"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return BackupRestoreModel{
		backups:      backups,
		list:         l,
		scriptRunner: scriptRunner,
	}
}

func (m BackupRestoreModel) Init() tea.Cmd {
	return nil
}

func (m BackupRestoreModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewBackupModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewBackupModel(), nil
				}
				// TODO: Implement actual restoration logic
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m BackupRestoreModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔄 Restaurer une Sauvegarde"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde disponible"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render("💡 Créez d'abord une sauvegarde pour pouvoir la restaurer"))
	} else {
		s.WriteString(SubtitleStyle.Render("⚠️ Sélectionnez une sauvegarde à restaurer"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Restaurer • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// BackupDeleteModel handles backup deletion
type BackupDeleteModel struct {
	backups      []string
	list         list.Model
	scriptRunner *scripts.ScriptRunner
}

func NewBackupDeleteModel() BackupDeleteModel {
	scriptRunner := scripts.NewScriptRunner()
	backups, _ := scriptRunner.ListBackups()

	items := make([]list.Item, len(backups))
	for i, backup := range backups {
		items[i] = MenuItem{
			title:       fmt.Sprintf("🗑️ %s", backup),
			description: "Cliquez pour supprimer cette sauvegarde",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       "🔙 Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🗑️ Supprimer une Sauvegarde"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return BackupDeleteModel{
		backups:      backups,
		list:         l,
		scriptRunner: scriptRunner,
	}
}

func (m BackupDeleteModel) Init() tea.Cmd {
	return nil
}

func (m BackupDeleteModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewBackupModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewBackupModel(), nil
				}
				// TODO: Implement actual deletion logic with confirmation
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m BackupDeleteModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🗑️ Supprimer une Sauvegarde"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde à supprimer"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render("💡 Aucune sauvegarde trouvée dans le système"))
	} else {
		s.WriteString(SubtitleStyle.Render("⚠️ Attention: Suppression définitive"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Supprimer • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// Message types for backup operations
type backupCompleteMsg struct {
	success bool
	message string
}
