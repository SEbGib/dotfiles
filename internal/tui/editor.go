package tui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type EditorModel struct {
	filePath string
	fileName string
	content  string
	message  string
	editing  bool
	saved    bool
}

func NewEditorModel(filePath, fileName string) EditorModel {
	// Resolve environment variables in file path
	resolvedPath := os.ExpandEnv(filePath)

	return EditorModel{
		filePath: resolvedPath,
		fileName: fileName,
		message:  "Chargement du fichier...",
	}
}

func (m EditorModel) Init() tea.Cmd {
	return m.loadFile()
}

func (m EditorModel) loadFile() tea.Cmd {
	return func() tea.Msg {
		// Try to read the file
		if content, err := os.ReadFile(m.filePath); err == nil {
			return fileLoadedMsg{content: string(content)}
		} else {
			// If file doesn't exist, create an empty one
			if os.IsNotExist(err) {
				return fileLoadedMsg{content: "# Nouveau fichier de configuration\n# Ajoutez votre configuration ici\n", error: nil}
			}
			return fileLoadedMsg{content: "", error: err}
		}
	}
}

type fileLoadedMsg struct {
	content string
	error   error
}

type fileSavedMsg struct {
	success bool
	error   error
}

func (m EditorModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewMainModel(), nil
		case "ctrl+s":
			return m, m.saveFile()
		case "e", "enter":
			if !m.editing {
				return m, m.openInEditor()
			}
		}

	case fileLoadedMsg:
		if msg.error != nil {
			m.message = fmt.Sprintf("Erreur lors du chargement: %v", msg.error)
		} else {
			m.content = msg.content
			m.message = "Fichier chargé avec succès"
		}

	case fileSavedMsg:
		if msg.success {
			m.message = "Fichier sauvegardé avec succès!"
			m.saved = true
		} else {
			m.message = fmt.Sprintf("Erreur lors de la sauvegarde: %v", msg.error)
		}
	}

	return m, nil
}

func (m EditorModel) saveFile() tea.Cmd {
	return func() tea.Msg {
		err := os.WriteFile(m.filePath, []byte(m.content), 0644)
		return fileSavedMsg{success: err == nil, error: err}
	}
}

func (m EditorModel) openInEditor() tea.Cmd {
	return func() tea.Msg {
		// Try different editors in order of preference
		editors := []string{"nvim", "vim", "nano", "code", "subl"}

		for _, editor := range editors {
			if _, err := exec.LookPath(editor); err == nil {
				// Found an available editor
				cmd := exec.Command(editor, m.filePath)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				if err := cmd.Run(); err == nil {
					return fileSavedMsg{success: true, error: nil}
				}
			}
		}

		return fileSavedMsg{success: false, error: fmt.Errorf("aucun éditeur disponible")}
	}
}

func (m EditorModel) View() string {
	var s strings.Builder

	// Header
	s.WriteString(CreateBanner(fmt.Sprintf("📝 Éditeur - %s", m.fileName)))
	s.WriteString("\n\n")

	// File info
	fileInfo := fmt.Sprintf("📁 Fichier: %s", m.filePath)
	s.WriteString(CreateCard("📋 Informations", fileInfo))
	s.WriteString("\n")

	// Content preview (first 10 lines)
	if m.content != "" {
		lines := strings.Split(m.content, "\n")
		previewLines := lines
		if len(lines) > 10 {
			previewLines = lines[:10]
		}

		var preview strings.Builder
		for i, line := range previewLines {
			if len(line) > 80 {
				line = line[:77] + "..."
			}
			preview.WriteString(fmt.Sprintf("%3d │ %s\n", i+1, line))
		}

		if len(lines) > 10 {
			preview.WriteString(fmt.Sprintf("... et %d lignes supplémentaires", len(lines)-10))
		}

		s.WriteString(CreateCard("📄 Aperçu du contenu", preview.String()))
		s.WriteString("\n")
	}

	// Status message
	if m.message != "" {
		s.WriteString(CreateStatusBadge("info", m.message))
		s.WriteString("\n")
	}

	// Instructions
	s.WriteString("\n")
	instructions := `Commandes disponibles:
• E/Entrée  - Ouvrir dans l'éditeur externe
• Ctrl+S    - Sauvegarder (si modifié)
• Échap     - Retour au menu
• Ctrl+C    - Quitter`

	s.WriteString(CreateCard("⌨️ Instructions", instructions))

	// Footer
	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("E pour éditer • Échap pour retour • Ctrl+C pour quitter"))

	return s.String()
}
