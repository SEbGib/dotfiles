package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// TwoColumnMenuItem represents a menu item with keyboard shortcut
type TwoColumnMenuItem struct {
	shortcut    string
	title       string
	description string
	action      string
}

// TwoColumnMainModel implements a two-column layout with keyboard shortcuts
type TwoColumnMainModel struct {
	items               []TwoColumnMenuItem
	selectedIndex       int
	searchMode          bool
	searchInput         textinput.Model
	filteredItems       []TwoColumnMenuItem
	originalItems       []TwoColumnMenuItem
	quitting            bool
	statusMsg           string
	width               int
	height              int
	notificationManager *NotificationManager
}

// NewTwoColumnMainModel creates a new two-column main model
func NewTwoColumnMainModel() TwoColumnMainModel {
	items := []TwoColumnMenuItem{
		{
			shortcut:    "1",
			title:       "🚀 Installation Interactive",
			description: "Guide d'installation complète étape par étape avec vérification automatique des prérequis, installation des outils essentiels et configuration personnalisée de votre environnement de développement.",
			action:      "install",
		},
		{
			shortcut:    "2",
			title:       "⚙️ Gestion de Configuration",
			description: "Éditer vos fichiers de configuration (.zshrc, .gitconfig, .tmux.conf, etc.) avec un éditeur intégré supportant la coloration syntaxique et la validation en temps réel.",
			action:      "config",
		},
		{
			shortcut:    "3",
			title:       "✅ Vérification du Système",
			description: "Vérifier l'installation et la santé du système avec des tests automatisés pour s'assurer que tous les outils sont correctement installés et configurés.",
			action:      "verify",
		},
		{
			shortcut:    "4",
			title:       "💾 Sauvegarde & Restauration",
			description: "Gérer les sauvegardes de vos configurations avec versioning automatique, restauration sélective et synchronisation cloud pour protéger vos paramètres personnalisés.",
			action:      "backup",
		},
		{
			shortcut:    "5",
			title:       "🔧 Gestion des Outils",
			description: "Installer, mettre à jour ou supprimer des outils de développement avec gestion des dépendances, vérification de compatibilité et installation automatisée.",
			action:      "tools",
		},
		{
			shortcut:    "6",
			title:       "🔐 Configuration des Secrets",
			description: "Configurer Bitwarden et la gestion des secrets avec chiffrement sécurisé, intégration CLI et synchronisation automatique pour protéger vos données sensibles.",
			action:      "secrets",
		},
		{
			shortcut:    "7",
			title:       "📊 Informations Système",
			description: "Afficher les informations détaillées sur votre système incluant les versions des outils installés, l'utilisation des ressources et les statistiques de performance.",
			action:      "info",
		},
		{
			shortcut:    "8",
			title:       "❌ Quitter",
			description: "Fermer l'application en sauvegardant automatiquement les modifications en cours et en nettoyant les fichiers temporaires.",
			action:      "quit",
		},
	}

	// Setup search input
	searchInput := textinput.New()
	searchInput.Placeholder = "Tapez pour rechercher..."
	searchInput.CharLimit = 50
	searchInput.Width = 50

	return TwoColumnMainModel{
		items:               items,
		selectedIndex:       0,
		searchMode:          false,
		searchInput:         searchInput,
		filteredItems:       items,
		originalItems:       items,
		width:               80,
		height:              24,
		notificationManager: NewNotificationManager(),
	}
}

// Init initializes the model
func (m TwoColumnMainModel) Init() tea.Cmd {
	return nil
}

// Update handles messages
func (m TwoColumnMainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case NotificationExpiredMsg:
		m.notificationManager.RemoveNotification(msg.ID)
		return m, nil

	case tea.KeyMsg:
		if m.searchMode {
			switch msg.String() {
			case "esc":
				// Exit search mode
				m.searchMode = false
				m.searchInput.SetValue("")
				m.resetFilter()
				return m, nil
			case "enter":
				// Select filtered item if any
				if len(m.filteredItems) > 0 && m.selectedIndex < len(m.filteredItems) {
					return m.handleMenuSelection(m.filteredItems[m.selectedIndex])
				}
				return m, nil
			case "up", "k":
				if m.selectedIndex > 0 {
					m.selectedIndex--
				}
				return m, nil
			case "down", "j":
				if m.selectedIndex < len(m.filteredItems)-1 {
					m.selectedIndex++
				}
				return m, nil
			default:
				// Update search input
				m.searchInput, cmd = m.searchInput.Update(msg)
				m.filterItems(m.searchInput.Value())
				return m, cmd
			}
		} else {
			switch msg.String() {
			case "/", "ctrl+f":
				// Enter search mode
				m.searchMode = true
				m.searchInput.Focus()
				return m, textinput.Blink
			case "ctrl+c":
				m.quitting = true
				return m, tea.Quit
			case "up", "k":
				if m.selectedIndex > 0 {
					m.selectedIndex--
				}
				return m, nil
			case "down", "j":
				if m.selectedIndex < len(m.filteredItems)-1 {
					m.selectedIndex++
				}
				return m, nil
			case "enter":
				if m.selectedIndex < len(m.filteredItems) {
					return m.handleMenuSelection(m.filteredItems[m.selectedIndex])
				}
				return m, nil
			case "?", "f1":
				// Show help overlay
				shortcuts := GetMainMenuShortcuts()
				helpModel := NewWithHelp(m, "Aide - Menu Principal", shortcuts)
				return helpModel, nil
			default:
				// Check for number shortcuts (1-8)
				if num, err := strconv.Atoi(msg.String()); err == nil && num >= 1 && num <= len(m.filteredItems) {
					return m.handleMenuSelection(m.filteredItems[num-1])
				}
			}
		}
	}

	return m, cmd
}

// filterItems filters menu items based on search query
func (m *TwoColumnMainModel) filterItems(query string) {
	if query == "" {
		m.filteredItems = m.originalItems
		m.selectedIndex = 0
		return
	}

	query = strings.ToLower(query)
	m.filteredItems = make([]TwoColumnMenuItem, 0)

	for _, item := range m.originalItems {
		if m.matchesQuery(item, query) {
			m.filteredItems = append(m.filteredItems, item)
		}
	}

	// Reset selection to first item
	m.selectedIndex = 0
}

// matchesQuery checks if an item matches the search query
func (m *TwoColumnMainModel) matchesQuery(item TwoColumnMenuItem, query string) bool {
	title := strings.ToLower(item.title)
	description := strings.ToLower(item.description)

	return strings.Contains(title, query) || strings.Contains(description, query)
}

// resetFilter resets the filter to show all items
func (m *TwoColumnMainModel) resetFilter() {
	m.filteredItems = m.originalItems
	m.selectedIndex = 0
}

// handleMenuSelection handles menu item selection
func (m TwoColumnMainModel) handleMenuSelection(item TwoColumnMenuItem) (tea.Model, tea.Cmd) {
	switch item.action {
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
	default:
		return m, nil
	}
}

// View renders the two-column layout
func (m TwoColumnMainModel) View() string {
	if m.quitting {
		return "Au revoir! 👋\n"
	}

	var s strings.Builder

	// Header - consistent with other screens
	s.WriteString(CreateBanner("🏠 Dotfiles Manager - Menu Principal"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Interface moderne pour la gestion de vos dotfiles"))
	s.WriteString("\n\n")

	// Search mode indicator and input
	if m.searchMode {
		searchHeader := lipgloss.NewStyle().
			Foreground(ColorInfo).
			Bold(true).
			Render("🔍 Mode Recherche")
		s.WriteString(searchHeader)
		s.WriteString("\n\n")

		searchBox := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(0, 1).
			Render(m.searchInput.View())
		s.WriteString(searchBox)
		s.WriteString("\n\n")

		// Show number of results
		resultCount := lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Render(fmt.Sprintf("%d résultat(s) trouvé(s)", len(m.filteredItems)))
		s.WriteString(resultCount)
		s.WriteString("\n\n")
	}

	// Two-column layout wrapped in a card for consistency
	twoColumnContent := m.renderTwoColumns()
	s.WriteString(CardStyle.Render(twoColumnContent))
	s.WriteString("\n")

	// Footer with instructions - consistent style
	var footerText string
	if m.searchMode {
		footerText = "• ↑/↓ Navigation • Entrée Sélectionner • Échap Quitter recherche • Ctrl+C Quitter"
	} else {
		footerText = "• ↑/↓ Navigation • 1-8 Raccourcis • Entrée Sélectionner • / Rechercher • Ctrl+C Quitter"
	}
	s.WriteString(FooterStyle.Render(footerText))

	// Status message if present
	if m.statusMsg != "" {
		s.WriteString("\n\n")
		s.WriteString(CreateStatusBadge("info", m.statusMsg))
	}

	baseView := AppStyle.Render(s.String())

	// Add notifications overlay
	notifications := m.notificationManager.RenderNotifications()
	if notifications != "" {
		return lipgloss.JoinVertical(lipgloss.Left, notifications, baseView)
	}

	return baseView
}

// renderTwoColumns renders the two-column layout
func (m TwoColumnMainModel) renderTwoColumns() string {
	if len(m.filteredItems) == 0 {
		return lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Padding(2).
			Align(lipgloss.Center).
			Render("🔍 Aucun résultat trouvé")
	}

	// Calculate column widths - more responsive
	totalWidth := m.width - 8 // Account for card padding
	if totalWidth < 80 {
		totalWidth = 80
	}
	leftWidth := totalWidth / 3      // 1/3 for menu items
	rightWidth := totalWidth * 2 / 3 // 2/3 for descriptions

	// Left column - menu items with improved styling
	var leftColumn strings.Builder

	// Header for left column
	leftHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Background(ColorBgSecondary).
		Padding(0, 1).
		Width(leftWidth - 2).
		Align(lipgloss.Center).
		Render("📋 Options")
	leftColumn.WriteString(leftHeader)
	leftColumn.WriteString("\n\n")

	// Menu items
	for i, item := range m.filteredItems {
		var itemStyle lipgloss.Style
		var prefix string

		if i == m.selectedIndex {
			itemStyle = MenuItemSelectedStyle.Copy().Width(leftWidth - 4)
			prefix = "▶ "
		} else {
			itemStyle = MenuItemStyle.Copy().Width(leftWidth - 4)
			prefix = "  "
		}

		menuText := fmt.Sprintf("%s%s %s", prefix, item.shortcut, item.title)
		leftColumn.WriteString(itemStyle.Render(menuText))
		leftColumn.WriteString("\n")
	}

	// Right column - description with improved styling
	var rightColumn strings.Builder

	// Header for right column
	rightHeader := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Background(ColorBgSecondary).
		Padding(0, 1).
		Width(rightWidth - 2).
		Align(lipgloss.Center).
		Render("📖 Description détaillée")
	rightColumn.WriteString(rightHeader)
	rightColumn.WriteString("\n\n")

	if m.selectedIndex < len(m.filteredItems) {
		selectedItem := m.filteredItems[m.selectedIndex]

		// Title with improved styling
		titleStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorSuccess).
			Background(ColorBgTertiary).
			Padding(0, 1).
			Width(rightWidth - 4).
			MarginBottom(1)
		rightColumn.WriteString(titleStyle.Render("🎯 " + selectedItem.title))
		rightColumn.WriteString("\n\n")

		// Description with better formatting
		descStyle := lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Width(rightWidth - 4).
			Padding(1).
			Background(ColorBgSecondary)
		rightColumn.WriteString(descStyle.Render(selectedItem.description))
		rightColumn.WriteString("\n\n")

		// Action hint with improved styling
		actionHint := lipgloss.NewStyle().
			Foreground(ColorInfo).
			Bold(true).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorInfo).
			Padding(0, 1).
			Background(ColorBgPrimary).
			Render(fmt.Sprintf("⌨️  Raccourci: %s  ou  ↵ Entrée", selectedItem.shortcut))
		rightColumn.WriteString(actionHint)
	}

	// Style columns with consistent borders
	leftColumnStyled := lipgloss.NewStyle().
		Width(leftWidth).
		Height(12). // Fixed height for consistency
		Padding(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorSecondary).
		Background(ColorBgPrimary).
		Render(leftColumn.String())

	rightColumnStyled := lipgloss.NewStyle().
		Width(rightWidth).
		Height(12). // Fixed height for consistency
		Padding(1).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(ColorSecondary).
		Background(ColorBgPrimary).
		Render(rightColumn.String())

	// Join columns horizontally with spacing
	return lipgloss.JoinHorizontal(lipgloss.Top, leftColumnStyled, " ", rightColumnStyled)
}

// Notification helper methods
func (m *TwoColumnMainModel) ShowNotification(notifType NotificationType, title, message string) tea.Cmd {
	return m.notificationManager.AddNotification(notifType, title, message)
}

func (m *TwoColumnMainModel) ShowSuccess(title, message string) tea.Cmd {
	return m.ShowNotification(NotificationSuccess, title, message)
}

func (m *TwoColumnMainModel) ShowWarning(title, message string) tea.Cmd {
	return m.ShowNotification(NotificationWarning, title, message)
}

func (m *TwoColumnMainModel) ShowError(title, message string) tea.Cmd {
	return m.ShowNotification(NotificationError, title, message)
}

func (m *TwoColumnMainModel) ShowInfo(title, message string) tea.Cmd {
	return m.ShowNotification(NotificationInfo, title, message)
}
