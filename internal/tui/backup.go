package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
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
			return NewMainModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok && i.action == "back" {
				return NewMainModel(), nil
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
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}
