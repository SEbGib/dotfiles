package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// EnhancedEditorModel extends EditorModel with better preview and features
type EnhancedEditorModel struct {
	EditorModel
	showLineNumbers bool
	previewLines    int
	syntaxType      string
	showMetadata    bool
}

// NewEnhancedEditorModel creates a new enhanced editor model
func NewEnhancedEditorModel(filePath, fileName string) EnhancedEditorModel {
	baseModel := NewEditorModel(filePath, fileName)

	return EnhancedEditorModel{
		EditorModel:     baseModel,
		showLineNumbers: true,
		previewLines:    15,
		syntaxType:      detectSyntaxType(fileName),
		showMetadata:    true,
	}
}

// detectSyntaxType detects the syntax type based on file extension
func detectSyntaxType(fileName string) string {
	ext := strings.ToLower(filepath.Ext(fileName))

	switch ext {
	case ".sh", ".bash", ".zsh":
		return "shell"
	case ".js", ".ts":
		return "javascript"
	case ".py":
		return "python"
	case ".go":
		return "go"
	case ".json":
		return "json"
	case ".yaml", ".yml":
		return "yaml"
	case ".toml":
		return "toml"
	case ".conf", ".config":
		return "config"
	case ".md":
		return "markdown"
	default:
		return "text"
	}
}

// Update handles messages for the enhanced editor
func (m EnhancedEditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+l":
			// Toggle line numbers
			m.showLineNumbers = !m.showLineNumbers
			return m, nil
		case "ctrl+m":
			// Toggle metadata
			m.showMetadata = !m.showMetadata
			return m, nil
		case "ctrl+p":
			// Increase preview lines
			if m.previewLines < 30 {
				m.previewLines += 5
			}
			return m, nil
		case "ctrl+o":
			// Decrease preview lines
			if m.previewLines > 5 {
				m.previewLines -= 5
			}
			return m, nil
		default:
			// Delegate to base editor
			updatedModel, cmd := m.EditorModel.Update(msg)
			if baseModel, ok := updatedModel.(EditorModel); ok {
				m.EditorModel = baseModel
			}
			return m, cmd
		}
	default:
		// Delegate to base editor
		updatedModel, cmd := m.EditorModel.Update(msg)
		if baseModel, ok := updatedModel.(EditorModel); ok {
			m.EditorModel = baseModel
		}
		return m, cmd
	}
}

// View renders the enhanced editor
func (m EnhancedEditorModel) View() string {
	var s strings.Builder

	// Header with file info
	s.WriteString(CreateBanner(fmt.Sprintf("📝 Éditeur Avancé - %s", m.fileName)))
	s.WriteString("\n\n")

	// File metadata if enabled
	if m.showMetadata {
		metadata := m.renderMetadata()
		s.WriteString(CreateCard("📋 Informations du fichier", metadata))
		s.WriteString("\n")
	}

	// Enhanced content preview
	if m.content != "" {
		preview := m.renderEnhancedPreview()
		s.WriteString(CreateCard("📄 Aperçu du contenu", preview))
		s.WriteString("\n")
	}

	// Status message
	if m.message != "" {
		s.WriteString(CreateStatusBadge("info", m.message))
		s.WriteString("\n")
	}

	// Enhanced instructions
	instructions := m.renderInstructions()
	s.WriteString(CreateCard("⌨️ Commandes disponibles", instructions))

	// Footer
	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("E pour éditer • Ctrl+L lignes • Ctrl+M métadonnées • Échap retour"))

	return s.String()
}

// renderMetadata renders file metadata
func (m EnhancedEditorModel) renderMetadata() string {
	var metadata strings.Builder

	// File path
	metadata.WriteString(fmt.Sprintf("📁 Chemin: %s\n", m.filePath))

	// File type
	metadata.WriteString(fmt.Sprintf("🏷️ Type: %s\n", m.syntaxType))

	// File stats if file exists
	if info, err := os.Stat(m.filePath); err == nil {
		metadata.WriteString(fmt.Sprintf("📊 Taille: %d octets\n", info.Size()))
		metadata.WriteString(fmt.Sprintf("🕒 Modifié: %s\n", info.ModTime().Format("2006-01-02 15:04:05")))
		metadata.WriteString(fmt.Sprintf("🔒 Permissions: %s\n", info.Mode().String()))
	} else {
		metadata.WriteString("📝 Nouveau fichier (sera créé lors de la sauvegarde)\n")
	}

	// Content stats
	if m.content != "" {
		lines := strings.Split(m.content, "\n")
		metadata.WriteString(fmt.Sprintf("📏 Lignes: %d\n", len(lines)))
		metadata.WriteString(fmt.Sprintf("🔤 Caractères: %d\n", len(m.content)))
	}

	return metadata.String()
}

// renderEnhancedPreview renders an enhanced content preview
func (m EnhancedEditorModel) renderEnhancedPreview() string {
	if m.content == "" {
		return "Fichier vide ou non chargé"
	}

	lines := strings.Split(m.content, "\n")
	previewLines := lines

	if len(lines) > m.previewLines {
		previewLines = lines[:m.previewLines]
	}

	var preview strings.Builder

	// Add syntax highlighting hints based on file type
	preview.WriteString(m.renderSyntaxHeader())
	preview.WriteString("\n")

	for i, line := range previewLines {
		lineNum := i + 1

		// Truncate long lines
		if len(line) > 100 {
			line = line[:97] + "..."
		}

		// Render line with optional line numbers
		if m.showLineNumbers {
			lineNumStyle := lipgloss.NewStyle().
				Foreground(ColorTextMuted).
				Width(4).
				Align(lipgloss.Right)

			contentStyle := lipgloss.NewStyle().
				Foreground(m.getLineColor(line))

			preview.WriteString(fmt.Sprintf("%s │ %s\n",
				lineNumStyle.Render(fmt.Sprintf("%d", lineNum)),
				contentStyle.Render(line)))
		} else {
			contentStyle := lipgloss.NewStyle().
				Foreground(m.getLineColor(line))
			preview.WriteString(contentStyle.Render(line))
			preview.WriteString("\n")
		}
	}

	if len(lines) > m.previewLines {
		moreLines := len(lines) - m.previewLines
		preview.WriteString(lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Italic(true).
			Render(fmt.Sprintf("\n... et %d ligne(s) supplémentaire(s)", moreLines)))
	}

	return preview.String()
}

// renderSyntaxHeader renders a header indicating the syntax type
func (m EnhancedEditorModel) renderSyntaxHeader() string {
	var icon string
	switch m.syntaxType {
	case "shell":
		icon = "🐚"
	case "javascript":
		icon = "🟨"
	case "python":
		icon = "🐍"
	case "go":
		icon = "🐹"
	case "json":
		icon = "📋"
	case "yaml":
		icon = "📄"
	case "toml":
		icon = "⚙️"
	case "markdown":
		icon = "📝"
	default:
		icon = "📄"
	}

	return lipgloss.NewStyle().
		Foreground(ColorInfo).
		Bold(true).
		Render(fmt.Sprintf("%s %s", icon, strings.ToUpper(m.syntaxType)))
}

// getLineColor returns appropriate color for a line based on content
func (m EnhancedEditorModel) getLineColor(line string) lipgloss.Color {
	trimmed := strings.TrimSpace(line)

	// Comments
	if strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "//") {
		return ColorTextMuted
	}

	// Keywords based on syntax type
	switch m.syntaxType {
	case "shell":
		if strings.HasPrefix(trimmed, "export") || strings.HasPrefix(trimmed, "alias") {
			return ColorPrimary
		}
	case "yaml", "toml":
		if strings.Contains(trimmed, ":") && !strings.HasPrefix(trimmed, " ") {
			return ColorPrimary
		}
	case "config":
		if strings.Contains(trimmed, "=") {
			return ColorPrimary
		}
	}

	// Empty lines
	if trimmed == "" {
		return ColorTextMuted
	}

	// Default
	return ColorTextPrimary
}

// renderInstructions renders enhanced instructions
func (m EnhancedEditorModel) renderInstructions() string {
	instructions := `Édition:
• E/Entrée    - Ouvrir dans l'éditeur externe
• Ctrl+S      - Sauvegarder (si modifié)

Affichage:
• Ctrl+L      - Basculer les numéros de ligne
• Ctrl+M      - Basculer les métadonnées
• Ctrl+P      - Plus de lignes d'aperçu
• Ctrl+O      - Moins de lignes d'aperçu

Navigation:
• Échap       - Retour au menu
• Ctrl+C      - Quitter`

	return instructions
}

// GetPreviewLines returns the current number of preview lines
func (m EnhancedEditorModel) GetPreviewLines() int {
	return m.previewLines
}

// GetSyntaxType returns the detected syntax type
func (m EnhancedEditorModel) GetSyntaxType() string {
	return m.syntaxType
}

// IsShowingLineNumbers returns whether line numbers are shown
func (m EnhancedEditorModel) IsShowingLineNumbers() bool {
	return m.showLineNumbers
}

// IsShowingMetadata returns whether metadata is shown
func (m EnhancedEditorModel) IsShowingMetadata() bool {
	return m.showMetadata
}
