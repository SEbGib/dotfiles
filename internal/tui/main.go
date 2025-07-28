package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Import the new styles (styles are now in styles.go)
// Menu items
type MenuItem struct {
	title       string
	description string
	action      string
}

func (m MenuItem) Title() string       { return m.title }
func (m MenuItem) Description() string { return m.description }
func (m MenuItem) FilterValue() string { return m.title }

// Main model
type MainModel struct {
	list        list.Model
	choice      string
	quitting    bool
	currentView string
	statusMsg   string
}

// Initialize the main model
func NewMainModel() MainModel {
	items := []list.Item{
		MenuItem{
			title:       "🚀 Installation Interactive",
			description: "Guide d'installation complète étape par étape",
			action:      "install",
		},
		MenuItem{
			title:       "⚙️ Gestion de Configuration",
			description: "Modifier et gérer vos configurations dotfiles",
			action:      "config",
		},
		MenuItem{
			title:       "✅ Vérification du Système",
			description: "Vérifier l'installation et la santé du système",
			action:      "verify",
		},
		MenuItem{
			title:       "💾 Sauvegarde & Restauration",
			description: "Gérer les sauvegardes de vos configurations",
			action:      "backup",
		},
		MenuItem{
			title:       "🔧 Gestion des Outils",
			description: "Installer, mettre à jour ou supprimer des outils",
			action:      "tools",
		},
		MenuItem{
			title:       "🔐 Configuration des Secrets",
			description: "Configurer Bitwarden et la gestion des secrets",
			action:      "secrets",
		},
		MenuItem{
			title:       "📊 Informations Système",
			description: "Afficher les informations sur votre système",
			action:      "info",
		},
		MenuItem{
			title:       "❌ Quitter",
			description: "Fermer l'application",
			action:      "quit",
		},
	}

	const defaultWidth = 80
	const listHeight = 14

	l := list.New(items, list.NewDefaultDelegate(), defaultWidth, listHeight)
	l.Title = "🏠 Dotfiles Manager - Menu Principal"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = HeaderStyle
	l.Styles.PaginationStyle = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	l.Styles.HelpStyle = HelpStyle

	return MainModel{
		list:        l,
		currentView: "main",
	}
}

// Initialize the model
func (m MainModel) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "ctrl+c":
			m.quitting = true
			return m, tea.Quit

		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				m.choice = i.action
				switch i.action {
				case "quit":
					m.quitting = true
					return m, tea.Quit
				case "install":
					return NewInstallModel(), nil
				case "verify":
					return NewVerifyModel(), nil
				case "config":
					return NewConfigModel(), nil
				case "backup":
					return NewBackupModel(), nil
				case "tools":
					return NewToolsModel(), nil
				case "secrets":
					return NewSecretsModel(), nil
				case "info":
					return NewInfoModel(), nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

// View renders the model
func (m MainModel) View() string {
	if m.quitting {
		return CreateStatusBadge("info", "Au revoir! 👋")
	}

	var s strings.Builder

	// Beautiful header with banner
	s.WriteString(CreateBanner("🚀 Dotfiles Manager - Configuration Moderne"))
	s.WriteString("\n\n")

	// Subtitle
	s.WriteString(SubtitleStyle.Render("Interface moderne pour la gestion de vos dotfiles"))
	s.WriteString("\n\n")

	// Main content in a card
	listContent := m.list.View()
	s.WriteString(CardStyle.Render(listContent))

	// Beautiful footer
	footerText := "• ↑/↓ Navigation • Entrée Sélectionner • Ctrl+C Quitter"
	s.WriteString(FooterStyle.Render(footerText))

	// Status message if present
	if m.statusMsg != "" {
		s.WriteString("\n")
		s.WriteString(CreateStatusBadge("info", m.statusMsg))
	}

	return AppStyle.Render(s.String())
}

// Key bindings
func (m MainModel) ShortHelp() []key.Binding {
	return []key.Binding{
		key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "sélectionner"),
		),
		key.NewBinding(
			key.WithKeys("ctrl+c"),
			key.WithHelp("ctrl+c", "quitter"),
		),
	}
}

func (m MainModel) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			key.NewBinding(
				key.WithKeys("enter"),
				key.WithHelp("enter", "sélectionner l'option"),
			),
			key.NewBinding(
				key.WithKeys("ctrl+c"),
				key.WithHelp("ctrl+c", "quitter l'application"),
			),
		},
	}
}
