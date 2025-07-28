package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

func (m SecretsModel) View() string {
	var s strings.Builder
	s.WriteString(lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#7D56F4")).Render("🔐 Configuration des Secrets"))
	s.WriteString("\n\n")
	s.WriteString(m.list.View())
	s.WriteString("\n")
	s.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render("• Entrée pour sélectionner • Échap pour retour • Ctrl+C pour quitter"))
	return s.String()
}
