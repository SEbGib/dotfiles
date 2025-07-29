package tui

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// SearchableMainModel extends MainModel with search functionality
type SearchableMainModel struct {
	MainModel
	searchMode    bool
	searchInput   textinput.Model
	filteredItems []MenuItem
	originalItems []MenuItem
}

// NewSearchableMainModel creates a new searchable main model
func NewSearchableMainModel() SearchableMainModel {
	mainModel := NewMainModel()

	// Extract original items
	originalItems := make([]MenuItem, len(mainModel.list.Items()))
	for i, item := range mainModel.list.Items() {
		if menuItem, ok := item.(MenuItem); ok {
			originalItems[i] = menuItem
		}
	}

	// Setup search input
	searchInput := textinput.New()
	searchInput.Placeholder = "Tapez pour rechercher..."
	searchInput.CharLimit = 50
	searchInput.Width = 50

	return SearchableMainModel{
		MainModel:     mainModel,
		searchMode:    false,
		searchInput:   searchInput,
		filteredItems: originalItems,
		originalItems: originalItems,
	}
}

// Update handles messages for the searchable main model
func (m SearchableMainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
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
				if len(m.filteredItems) > 0 {
					selectedItem := m.filteredItems[0] // Select first match
					return m.handleMenuSelection(selectedItem)
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
			case "enter":
				// Handle normal menu selection
				i, ok := m.list.SelectedItem().(MenuItem)
				if ok {
					return m.handleMenuSelection(i)
				}
			default:
				// Normal navigation
				updatedModel, cmd := m.MainModel.Update(msg)
				m.MainModel = updatedModel.(MainModel)
				return m, cmd
			}
		}
	}

	if !m.searchMode {
		updatedModel, cmd := m.MainModel.Update(msg)
		m.MainModel = updatedModel.(MainModel)
		return m, cmd
	}
	return m, cmd
}

// filterItems filters menu items based on search query
func (m *SearchableMainModel) filterItems(query string) {
	if query == "" {
		m.filteredItems = m.originalItems
		m.updateListItems()
		return
	}

	query = strings.ToLower(query)
	m.filteredItems = make([]MenuItem, 0)

	for _, item := range m.originalItems {
		if m.matchesQuery(item, query) {
			m.filteredItems = append(m.filteredItems, item)
		}
	}

	m.updateListItems()
}

// matchesQuery checks if an item matches the search query using fuzzy matching
func (m *SearchableMainModel) matchesQuery(item MenuItem, query string) bool {
	title := strings.ToLower(item.title)
	description := strings.ToLower(item.description)

	// Exact match in title or description
	if strings.Contains(title, query) || strings.Contains(description, query) {
		return true
	}

	// Fuzzy match - check if all characters of query appear in order
	return m.fuzzyMatch(title, query) || m.fuzzyMatch(description, query)
}

// fuzzyMatch performs fuzzy string matching
func (m *SearchableMainModel) fuzzyMatch(text, pattern string) bool {
	if len(pattern) == 0 {
		return true
	}
	if len(text) == 0 {
		return false
	}

	textRunes := []rune(text)
	patternRunes := []rune(pattern)

	textIdx := 0
	patternIdx := 0

	for textIdx < len(textRunes) && patternIdx < len(patternRunes) {
		if unicode.ToLower(textRunes[textIdx]) == unicode.ToLower(patternRunes[patternIdx]) {
			patternIdx++
		}
		textIdx++
	}

	return patternIdx == len(patternRunes)
}

// updateListItems updates the list with filtered items
func (m *SearchableMainModel) updateListItems() {
	items := make([]list.Item, len(m.filteredItems))
	for i, item := range m.filteredItems {
		items[i] = item
	}
	m.list.SetItems(items)
}

// resetFilter resets the filter to show all items
func (m *SearchableMainModel) resetFilter() {
	m.filteredItems = m.originalItems
	m.updateListItems()
}

// handleMenuSelection handles menu item selection
func (m SearchableMainModel) handleMenuSelection(item MenuItem) (tea.Model, tea.Cmd) {
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

// View renders the searchable main model
func (m SearchableMainModel) View() string {
	if m.quitting {
		return "Au revoir! 👋\n"
	}

	var s strings.Builder

	// Header
	s.WriteString("\n")
	s.WriteString(CreateBanner("🚀 Dotfiles Manager"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Interface moderne pour la gestion de vos dotfiles"))
	s.WriteString("\n\n")

	// Search mode indicator and input
	if m.searchMode {
		s.WriteString(lipgloss.NewStyle().
			Foreground(ColorInfo).
			Bold(true).
			Render("🔍 Mode Recherche"))
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
	} else {
		// Add separator for better visual separation
		s.WriteString(CreateSeparator(60))
		s.WriteString("\n\n")
	}

	// Main content
	s.WriteString(m.list.View())
	s.WriteString("\n")

	// Add separator before footer
	s.WriteString(CreateSeparator(60))
	s.WriteString("\n")

	// Footer with search instructions
	var footerText string
	if m.searchMode {
		footerText = "↑/↓ Navigation • Entrée Sélectionner • Échap Quitter recherche • Ctrl+C Quitter"
	} else {
		footerText = "↑/↓ Navigation • Entrée Sélectionner • / Rechercher • Ctrl+C Quitter"
	}
	s.WriteString(FooterStyle.Render(footerText))

	// Status message if present
	if m.statusMsg != "" {
		s.WriteString("\n\n")
		s.WriteString(CreateStatusBadge("info", m.statusMsg))
	}

	return s.String()
}
