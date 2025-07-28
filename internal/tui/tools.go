package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ToolsModel struct {
	list list.Model
}

func NewToolsModel() ToolsModel {
	items := []list.Item{
		MenuItem{title: "📦 Installer des outils", description: "Installer de nouveaux outils", action: "install_tools"},
		MenuItem{title: "🔄 Mettre à jour les outils", description: "Mettre à jour tous les outils", action: "update_tools"},
		MenuItem{title: "📋 Lister les outils installés", description: "Voir tous les outils installés", action: "list_tools"},
		MenuItem{title: "🗑️ Désinstaller des outils", description: "Supprimer des outils", action: "uninstall_tools"},
		MenuItem{title: "🔙 Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "🔧 Gestion des Outils"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ToolsModel{list: l}
}

func (m ToolsModel) Init() tea.Cmd {
	return nil
}

func (m ToolsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m ToolsModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner("🔧 Gestion des Outils"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Installez et gérez vos outils de développement"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}
