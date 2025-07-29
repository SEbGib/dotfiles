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
		MenuItem{title: " Créer une sauvegarde", description: "Sauvegarder les configurations actuelles", action: "create_backup"},
		MenuItem{title: " Lister les sauvegardes", description: "Voir toutes les sauvegardes disponibles", action: "list_backups"},
		MenuItem{title: " Restaurer une sauvegarde", description: "Restaurer depuis une sauvegarde", action: "restore_backup"},
		MenuItem{title: " Supprimer une sauvegarde", description: "Supprimer une sauvegarde ancienne", action: "delete_backup"},
		MenuItem{title: " Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Sauvegarde & Restauration"
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

	s.WriteString(CreateBanner(" Sauvegarde & Restauration"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Gérez vos sauvegardes de configuration"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour au menu • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// BackupCreateModel handles backup creation
type BackupCreateModel struct {
	progress     UnifiedProgressModel
	complete     bool
	scriptRunner *scripts.ScriptRunner
}

func NewBackupCreateModel() BackupCreateModel {
	progress := NewBackupProgress()

	return BackupCreateModel{
		progress:     progress,
		scriptRunner: scripts.NewScriptRunner(),
	}
}

func (m BackupCreateModel) Init() tea.Cmd {
	return tea.Batch(
		m.progress.Init(),
		m.createBackup(),
	)
}

func (m BackupCreateModel) createBackup() tea.Cmd {
	return func() tea.Msg {
		var logs []string
		addLog := func(msg string) {
			timestamp := time.Now().Format("15:04:05")
			logs = append(logs, fmt.Sprintf("[%s] %s", timestamp, msg))
		}

		// Step 1: Initialize
		addLog("Démarrage de la sauvegarde...")

		homeDir := os.Getenv("HOME")
		if homeDir == "" {
			addLog("Erreur: Impossible de déterminer le répertoire home")
			return ProgressFinishedMsg{
				Success: false,
				Message: "Impossible de déterminer le répertoire home",
			}
		}

		addLog(fmt.Sprintf("Répertoire home: %s", homeDir))

		filesToBackup := []string{
			".zshrc", ".gitconfig", ".aliases", ".tmux.conf",
			".config/starship.toml", ".config/nvim", ".config/tmux",
		}

		addLog(fmt.Sprintf("Fichiers à sauvegarder: %d", len(filesToBackup)))

		// Step 2: Create backup directory
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		backupDir := filepath.Join(homeDir, fmt.Sprintf(".dotfiles-backup-%s", timestamp))

		addLog(fmt.Sprintf("Création du répertoire: %s", backupDir))
		if err := os.MkdirAll(backupDir, 0755); err != nil {
			addLog(fmt.Sprintf("Erreur création répertoire: %v", err))
			return ProgressFinishedMsg{
				Success: false,
				Message: fmt.Sprintf("Erreur création répertoire: %v", err),
			}
		}

		// Step 3: Copy files
		backedUp := 0
		skipped := 0

		for _, file := range filesToBackup {
			srcPath := filepath.Join(homeDir, file)

			addLog(fmt.Sprintf("Vérification: %s", file))

			// Check if source exists
			srcInfo, err := os.Stat(srcPath)
			if err != nil {
				addLog(fmt.Sprintf("Ignoré (n'existe pas): %s", file))
				skipped++
				continue
			}

			dstPath := filepath.Join(backupDir, file)

			// Create destination directory
			if err := os.MkdirAll(filepath.Dir(dstPath), 0755); err != nil {
				addLog(fmt.Sprintf("Erreur création dossier pour %s: %v", file, err))
				continue
			}

			// Copy file or directory
			if srcInfo.IsDir() {
				addLog(fmt.Sprintf("Copie dossier: %s", file))
				if err := copyDir(srcPath, dstPath); err == nil {
					backedUp++
					addLog(fmt.Sprintf("✓ Dossier copié: %s", file))
				} else {
					addLog(fmt.Sprintf("Erreur copie dossier %s: %v", file, err))
				}
			} else {
				addLog(fmt.Sprintf("Copie fichier: %s", file))
				if data, err := os.ReadFile(srcPath); err == nil {
					if err := os.WriteFile(dstPath, data, srcInfo.Mode()); err == nil {
						backedUp++
						addLog(fmt.Sprintf("✓ Fichier copié: %s", file))
					} else {
						addLog(fmt.Sprintf("Erreur écriture %s: %v", file, err))
					}
				} else {
					addLog(fmt.Sprintf("Erreur lecture %s: %v", file, err))
				}
			}

			time.Sleep(time.Millisecond * 200) // Allow UI to update
		}

		// Step 4: Finalize
		addLog(fmt.Sprintf("Sauvegarde terminée: %d copiés, %d ignorés", backedUp, skipped))

		return BackupCompleteMsg{
			Success: true,
			Message: fmt.Sprintf("Sauvegarde créée: %s (%d fichiers copiés, %d ignorés)", filepath.Base(backupDir), backedUp, skipped),
			Logs:    logs,
		}
	}
}

// BackupCompleteMsg represents completion of backup operation
type BackupCompleteMsg struct {
	Success bool
	Message string
	Logs    []string
}

// copyDir recursively copies a directory
func copyDir(src, dst string) error {
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Calculate destination path
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Copy file
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(dstPath, data, info.Mode())
	})
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
	case BackupCompleteMsg:
		m.complete = true
		// Add logs to progress model
		for _, log := range msg.Logs {
			m.progress.AddLog(log)
		}
		m.progress.Message = msg.Message
	case ProgressFinishedMsg:
		m.complete = true
	default:
		var cmd tea.Cmd
		m.progress, cmd = m.progress.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m BackupCreateModel) View() string {
	if m.complete {
		var s strings.Builder
		s.WriteString(CreateBanner(" Création de Sauvegarde"))
		s.WriteString("\n\n")
		s.WriteString(SubtitleStyle.Render(" Sauvegarde terminée avec succès!"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.progress.Message))
		s.WriteString("\n\n")
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
		return AppStyle.Render(s.String())
	}

	return m.progress.View()
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
			title:       fmt.Sprintf(" %s", backup),
			description: "Sauvegarde disponible",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Liste des Sauvegardes"
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

	s.WriteString(CreateBanner(" Liste des Sauvegardes"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde trouvée"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(" Créez votre première sauvegarde depuis le menu principal"))
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
			title:       fmt.Sprintf(" %s", backup),
			description: "Cliquez pour restaurer cette sauvegarde",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Restaurer une Sauvegarde"
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

	s.WriteString(CreateBanner(" Restaurer une Sauvegarde"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde disponible"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(" Créez d'abord une sauvegarde pour pouvoir la restaurer"))
	} else {
		s.WriteString(SubtitleStyle.Render(" Sélectionnez une sauvegarde à restaurer"))
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
			title:       fmt.Sprintf(" %s", backup),
			description: "Cliquez pour supprimer cette sauvegarde",
			action:      backup,
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu sauvegardes",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Supprimer une Sauvegarde"
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

	s.WriteString(CreateBanner(" Supprimer une Sauvegarde"))
	s.WriteString("\n\n")

	if len(m.backups) == 0 {
		s.WriteString(SubtitleStyle.Render("Aucune sauvegarde à supprimer"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(" Aucune sauvegarde trouvée dans le système"))
	} else {
		s.WriteString(SubtitleStyle.Render(" Attention: Suppression définitive"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Supprimer • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// Message types for backup operations (legacy - now using ProgressFinishedMsg)
// type backupCompleteMsg struct {
// 	success bool
// 	message string
// }
