package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type ConfigModel struct {
	list list.Model
}

func NewConfigModel() ConfigModel {
	items := []list.Item{
		MenuItem{title: "📝 Éditer .zshrc", description: "Configuration du shell Zsh", action: "edit_zshrc"},
		MenuItem{title: "⚙️ Éditer .gitconfig", description: "Configuration Git", action: "edit_gitconfig"},
		MenuItem{title: "🎨 Éditer starship.toml", description: "Configuration du prompt", action: "edit_starship"},
		MenuItem{title: "📁 Éditer init.lua (Neovim)", description: "Configuration de l'éditeur", action: "edit_nvim"},
		MenuItem{title: "🖥️ Éditer tmux.conf", description: "Configuration du multiplexeur", action: "edit_tmux"},
		MenuItem{title: "🔧 Éditer .aliases", description: "Aliases personnalisés", action: "edit_aliases"},
		MenuItem{title: "🔙 Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = "⚙️ Gestion de Configuration"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ConfigModel{list: l}
}

func (m ConfigModel) Init() tea.Cmd {
	return nil
}

func (m ConfigModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewMainModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				switch i.action {
				case "back":
					return NewMainModel(), nil
				case "edit_zshrc":
					return NewEditorModel("$HOME/.zshrc", ".zshrc"), nil
				case "edit_gitconfig":
					return NewEditorModel("$HOME/.gitconfig", ".gitconfig"), nil
				case "edit_starship":
					return NewEditorModel("$HOME/.config/starship.toml", "starship.toml"), nil
				case "edit_nvim":
					return NewEditorModel("$HOME/.config/nvim/init.lua", "init.lua"), nil
				case "edit_tmux":
					return NewEditorModel("$HOME/.config/tmux/tmux.conf", "tmux.conf"), nil
				case "edit_aliases":
					return NewEditorModel("$HOME/.aliases", ".aliases"), nil
				}
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ConfigModel) View() string {
	var s strings.Builder

	// Beautiful header
	s.WriteString(CreateBanner("⚙️ Gestion de Configuration"))
	s.WriteString("\n\n")

	// Subtitle
	s.WriteString(SubtitleStyle.Render("Éditez et gérez vos fichiers de configuration"))
	s.WriteString("\n\n")

	// Main content in a card
	listContent := m.list.View()
	s.WriteString(CardStyle.Render(listContent))

	// Footer
	footerText := "• Entrée Sélectionner • Échap Retour • Ctrl+C Quitter"
	s.WriteString(FooterStyle.Render(footerText))

	return AppStyle.Render(s.String())
}
