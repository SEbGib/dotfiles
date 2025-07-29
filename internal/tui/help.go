package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// HelpOverlay represents a help overlay that can be shown over any model
type HelpOverlay struct {
	visible   bool
	shortcuts []Shortcut
	title     string
}

// Shortcut represents a keyboard shortcut
type Shortcut struct {
	Key         string
	Description string
	Category    string
}

// NewHelpOverlay creates a new help overlay
func NewHelpOverlay(title string, shortcuts []Shortcut) HelpOverlay {
	return HelpOverlay{
		visible:   false,
		shortcuts: shortcuts,
		title:     title,
	}
}

// Toggle toggles the help overlay visibility
func (h *HelpOverlay) Toggle() {
	h.visible = !h.visible
}

// Show shows the help overlay
func (h *HelpOverlay) Show() {
	h.visible = true
}

// Hide hides the help overlay
func (h *HelpOverlay) Hide() {
	h.visible = false
}

// IsVisible returns whether the help overlay is visible
func (h *HelpOverlay) IsVisible() bool {
	return h.visible
}

// Render renders the help overlay
func (h *HelpOverlay) Render(baseView string) string {
	if !h.visible {
		return baseView
	}

	helpContent := h.renderHelpContent()

	// Create overlay style
	overlayStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(ColorPrimary).
		Background(ColorBgSecondary).
		Padding(1, 2).
		Margin(2, 4)

	overlay := overlayStyle.Render(helpContent)

	// Position overlay in center
	return lipgloss.Place(
		lipgloss.Width(baseView),
		lipgloss.Height(baseView),
		lipgloss.Center,
		lipgloss.Center,
		overlay,
		lipgloss.WithWhitespaceChars(" "),
		lipgloss.WithWhitespaceForeground(lipgloss.Color("0")),
	)
}

// renderHelpContent renders the help content
func (h *HelpOverlay) renderHelpContent() string {
	var content strings.Builder

	// Title
	title := lipgloss.NewStyle().
		Bold(true).
		Foreground(ColorPrimary).
		Align(lipgloss.Center).
		Render(" " + h.title)
	content.WriteString(title)
	content.WriteString("\n\n")

	// Group shortcuts by category
	categories := h.groupShortcutsByCategory()

	for category, shortcuts := range categories {
		if category != "" {
			categoryStyle := lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorSecondary).
				Render(category)
			content.WriteString(categoryStyle)
			content.WriteString("\n")
		}

		for _, shortcut := range shortcuts {
			keyStyle := lipgloss.NewStyle().
				Foreground(ColorPrimary).
				Bold(true).
				Width(12).
				Render(shortcut.Key)

			descStyle := lipgloss.NewStyle().
				Foreground(ColorTextPrimary).
				Render(shortcut.Description)

			content.WriteString("  " + keyStyle + " " + descStyle + "\n")
		}
		content.WriteString("\n")
	}

	// Footer
	footer := lipgloss.NewStyle().
		Foreground(ColorTextMuted).
		Italic(true).
		Align(lipgloss.Center).
		Render("Appuyez sur ? ou F1 pour fermer l'aide")
	content.WriteString(footer)

	return content.String()
}

// groupShortcutsByCategory groups shortcuts by their category
func (h *HelpOverlay) groupShortcutsByCategory() map[string][]Shortcut {
	categories := make(map[string][]Shortcut)

	for _, shortcut := range h.shortcuts {
		category := shortcut.Category
		if category == "" {
			category = "Général"
		}
		categories[category] = append(categories[category], shortcut)
	}

	return categories
}

// WithHelp wraps a model with help overlay capabilities
type WithHelp struct {
	Model tea.Model
	Help  HelpOverlay
}

// NewWithHelp creates a new model with help overlay
func NewWithHelp(model tea.Model, title string, shortcuts []Shortcut) WithHelp {
	return WithHelp{
		Model: model,
		Help:  NewHelpOverlay(title, shortcuts),
	}
}

// Init initializes the wrapped model
func (wh WithHelp) Init() tea.Cmd {
	return wh.Model.Init()
}

// Update handles messages for the wrapped model and help overlay
func (wh WithHelp) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// Handle help toggle
		if msg.String() == "?" || msg.String() == "f1" {
			wh.Help.Toggle()
			return wh, nil
		}

		// If help is visible, hide it on any other key
		if wh.Help.IsVisible() {
			wh.Help.Hide()
			return wh, nil
		}
	}

	// Update the wrapped model only if help is not visible
	if !wh.Help.IsVisible() {
		var cmd tea.Cmd
		updatedModel, cmd := wh.Model.Update(msg)

		// Check if the model type has changed (indicating navigation)
		// by comparing type names using reflection-free approach
		originalType := fmt.Sprintf("%T", wh.Model)
		updatedType := fmt.Sprintf("%T", updatedModel)

		// If the model type changed, return the new model directly (navigation)
		if originalType != updatedType {
			return updatedModel, cmd
		}

		wh.Model = updatedModel
		return wh, cmd
	}
	return wh, nil
}

// View renders the wrapped model with help overlay
func (wh WithHelp) View() string {
	baseView := wh.Model.View()
	return wh.Help.Render(baseView)
}

// Predefined shortcut sets for common models

// GetMainMenuShortcuts returns shortcuts for the main menu
func GetMainMenuShortcuts() []Shortcut {
	return []Shortcut{
		{"↑/↓", "Naviguer dans le menu", "Navigation"},
		{"Entrée", "Sélectionner l'option", "Navigation"},
		{"/", "Rechercher dans le menu", "Navigation"},
		{"Ctrl+C", "Quitter l'application", "Général"},
		{"?", "Afficher/masquer l'aide", "Général"},
	}
}

// GetEditorShortcuts returns shortcuts for the editor
func GetEditorShortcuts() []Shortcut {
	return []Shortcut{
		{"E/Entrée", "Ouvrir dans l'éditeur externe", "Édition"},
		{"Ctrl+S", "Sauvegarder le fichier", "Édition"},
		{"Ctrl+L", "Basculer les numéros de ligne", "Affichage"},
		{"Ctrl+M", "Basculer les métadonnées", "Affichage"},
		{"Ctrl+P", "Plus de lignes d'aperçu", "Affichage"},
		{"Ctrl+O", "Moins de lignes d'aperçu", "Affichage"},
		{"Échap", "Retour au menu principal", "Navigation"},
		{"Ctrl+C", "Quitter l'application", "Général"},
		{"?", "Afficher/masquer l'aide", "Général"},
	}
}

// GetVerificationShortcuts returns shortcuts for the verification screen
func GetVerificationShortcuts() []Shortcut {
	return []Shortcut{
		{"Entrée", "Démarrer la vérification", "Actions"},
		{"R", "Relancer la vérification", "Actions"},
		{"Échap", "Retour au menu principal", "Navigation"},
		{"Ctrl+C", "Quitter l'application", "Général"},
		{"?", "Afficher/masquer l'aide", "Général"},
	}
}

// GetInstallationShortcuts returns shortcuts for the installation screen
func GetInstallationShortcuts() []Shortcut {
	return []Shortcut{
		{"Entrée", "Démarrer l'installation", "Actions"},
		{"Espace", "Pause/Reprendre", "Actions"},
		{"S", "Ignorer l'étape actuelle", "Actions"},
		{"Échap", "Retour au menu principal", "Navigation"},
		{"Ctrl+C", "Annuler l'installation", "Général"},
		{"?", "Afficher/masquer l'aide", "Général"},
	}
}
