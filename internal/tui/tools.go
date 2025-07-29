package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebastiengiband/dotfiles/internal/scripts"
)

type ToolsModel struct {
	list list.Model
}

func NewToolsModel() ToolsModel {
	items := []list.Item{
		MenuItem{title: " Installer des outils", description: "Installer de nouveaux outils", action: "install_tools"},
		MenuItem{title: " Mettre à jour les outils", description: "Mettre à jour tous les outils", action: "update_tools"},
		MenuItem{title: " Lister les outils installés", description: "Voir tous les outils installés", action: "list_tools"},
		MenuItem{title: " Désinstaller des outils", description: "Supprimer des outils", action: "uninstall_tools"},
		MenuItem{title: " Retour au menu principal", description: "", action: "back"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Gestion des Outils"
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
			return NewTwoColumnMainModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				switch i.action {
				case "back":
					return NewTwoColumnMainModel(), nil
				case "install_tools":
					return NewToolsInstallModel(), nil
				case "update_tools":
					return NewToolsUpdateModel(), nil
				case "list_tools":
					return NewToolsListModel(), nil
				case "uninstall_tools":
					return NewToolsUninstallModel(), nil
				}
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
	s.WriteString(FooterStyle.Render("• Entrée Sélectionner • Échap Retour au menu • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// ToolsInstallModel handles tools installation
type ToolsInstallModel struct {
	availableTools []ToolInfo
	list           list.Model
	scriptRunner   *scripts.ScriptRunner
}

type ToolInfo struct {
	name        string
	description string
	installed   bool
	command     string
}

func NewToolsInstallModel() ToolsInstallModel {
	scriptRunner := scripts.NewScriptRunner()

	// Define available tools
	availableTools := []ToolInfo{
		{name: "chezmoi", description: "Gestionnaire de dotfiles", command: "chezmoi"},
		{name: "starship", description: "Prompt moderne", command: "starship"},
		{name: "fzf", description: "Recherche floue", command: "fzf"},
		{name: "ripgrep", description: "Recherche dans fichiers", command: "rg"},
		{name: "fd", description: "Alternative à find", command: "fd"},
		{name: "bat", description: "Alternative à cat", command: "bat"},
		{name: "eza", description: "Alternative à ls", command: "eza"},
		{name: "lazygit", description: "Interface Git", command: "lazygit"},
	}

	// Check which tools are installed
	for i := range availableTools {
		availableTools[i].installed = scriptRunner.CheckCommand(availableTools[i].command)
	}

	// Create list items
	items := make([]list.Item, 0)
	for _, tool := range availableTools {
		if !tool.installed {
			items = append(items, MenuItem{
				title:       fmt.Sprintf(" %s", tool.name),
				description: tool.description,
				action:      tool.name,
			})
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu outils",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Installer des Outils"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ToolsInstallModel{
		availableTools: availableTools,
		list:           l,
		scriptRunner:   scriptRunner,
	}
}

func (m ToolsInstallModel) Init() tea.Cmd {
	return nil
}

func (m ToolsInstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewToolsModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewToolsModel(), nil
				}
				// Create progress model for installation
				progress := NewToolsProgress("install")
				return NewToolsProgressModel(progress, i.action), nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ToolsInstallModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner(" Installer des Outils"))
	s.WriteString("\n\n")

	// Count non-installed tools
	nonInstalled := 0
	for _, tool := range m.availableTools {
		if !tool.installed {
			nonInstalled++
		}
	}

	if nonInstalled == 0 {
		s.WriteString(SubtitleStyle.Render(" Tous les outils sont déjà installés"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render("Félicitations! Votre système est parfaitement configuré."))
	} else {
		s.WriteString(SubtitleStyle.Render(fmt.Sprintf("%d outil(s) disponible(s) à l'installation", nonInstalled)))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Installer • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// ToolsUpdateModel handles tools update
type ToolsUpdateModel struct {
	installedTools []ToolInfo
	list           list.Model
	scriptRunner   *scripts.ScriptRunner
}

func NewToolsUpdateModel() ToolsUpdateModel {
	scriptRunner := scripts.NewScriptRunner()
	installedTools := scriptRunner.GetInstalledTools()

	items := make([]list.Item, 0)
	for tool, installed := range installedTools {
		if installed {
			items = append(items, MenuItem{
				title:       fmt.Sprintf(" %s", tool),
				description: "Mettre à jour cet outil",
				action:      tool,
			})
		}
	}

	// Add update all option
	if len(items) > 0 {
		items = append([]list.Item{MenuItem{
			title:       " Tout mettre à jour",
			description: "Mettre à jour tous les outils installés",
			action:      "update_all",
		}}, items...)
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu outils",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Mettre à Jour les Outils"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ToolsUpdateModel{
		list:         l,
		scriptRunner: scriptRunner,
	}
}

func (m ToolsUpdateModel) Init() tea.Cmd {
	return nil
}

func (m ToolsUpdateModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewToolsModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewToolsModel(), nil
				}
				// Create progress model for update
				progress := NewToolsProgress("update")
				return NewToolsProgressModel(progress, i.action), nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ToolsUpdateModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner(" Mettre à Jour les Outils"))
	s.WriteString("\n\n")
	s.WriteString(SubtitleStyle.Render("Sélectionnez les outils à mettre à jour"))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Mettre à jour • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// ToolsListModel handles listing installed tools
type ToolsListModel struct {
	installedTools map[string]bool
	list           list.Model
	scriptRunner   *scripts.ScriptRunner
}

func NewToolsListModel() ToolsListModel {
	scriptRunner := scripts.NewScriptRunner()
	installedTools := scriptRunner.GetInstalledTools()

	items := make([]list.Item, 0)
	for tool, installed := range installedTools {
		status := " Non installé"
		if installed {
			status = " Installé"
		}

		items = append(items, MenuItem{
			title:       fmt.Sprintf("%s %s", status, tool),
			description: fmt.Sprintf("État: %s", status),
			action:      tool,
		})
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu outils",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Outils Installés"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ToolsListModel{
		installedTools: installedTools,
		list:           l,
		scriptRunner:   scriptRunner,
	}
}

func (m ToolsListModel) Init() tea.Cmd {
	return nil
}

func (m ToolsListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewToolsModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok && i.action == "back" {
				return NewToolsModel(), nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ToolsListModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner(" État des Outils"))
	s.WriteString("\n\n")

	// Count installed tools
	installed := 0
	total := len(m.installedTools)
	for _, isInstalled := range m.installedTools {
		if isInstalled {
			installed++
		}
	}

	s.WriteString(SubtitleStyle.Render(fmt.Sprintf("%d/%d outils installés", installed, total)))
	s.WriteString("\n\n")
	s.WriteString(CardStyle.Render(m.list.View()))
	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Détails • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// ToolsUninstallModel handles tools uninstallation
type ToolsUninstallModel struct {
	installedTools map[string]bool
	list           list.Model
	scriptRunner   *scripts.ScriptRunner
}

func NewToolsUninstallModel() ToolsUninstallModel {
	scriptRunner := scripts.NewScriptRunner()
	installedTools := scriptRunner.GetInstalledTools()

	items := make([]list.Item, 0)
	for tool, installed := range installedTools {
		if installed {
			items = append(items, MenuItem{
				title:       fmt.Sprintf(" %s", tool),
				description: "Cliquez pour désinstaller",
				action:      tool,
			})
		}
	}

	// Add back option
	items = append(items, MenuItem{
		title:       " Retour",
		description: "Retour au menu outils",
		action:      "back",
	})

	l := list.New(items, list.NewDefaultDelegate(), 80, 14)
	l.Title = " Désinstaller des Outils"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)

	return ToolsUninstallModel{
		installedTools: installedTools,
		list:           l,
		scriptRunner:   scriptRunner,
	}
}

func (m ToolsUninstallModel) Init() tea.Cmd {
	return nil
}

func (m ToolsUninstallModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return NewToolsModel(), nil
		case "enter":
			i, ok := m.list.SelectedItem().(MenuItem)
			if ok {
				if i.action == "back" {
					return NewToolsModel(), nil
				}
				// TODO: Implement actual uninstallation with confirmation
				return m, nil
			}
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m ToolsUninstallModel) View() string {
	var s strings.Builder

	s.WriteString(CreateBanner(" Désinstaller des Outils"))
	s.WriteString("\n\n")

	// Count installed tools
	installed := 0
	for _, isInstalled := range m.installedTools {
		if isInstalled {
			installed++
		}
	}

	if installed == 0 {
		s.WriteString(SubtitleStyle.Render("Aucun outil à désinstaller"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(" Aucun outil installé trouvé"))
	} else {
		s.WriteString(SubtitleStyle.Render(" Attention: Désinstallation définitive"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.list.View()))
	}

	s.WriteString("\n")
	s.WriteString(FooterStyle.Render("• Entrée Désinstaller • Échap Retour • Ctrl+C Quitter"))

	return AppStyle.Render(s.String())
}

// ToolsProgressModel handles progress for tools operations
type ToolsProgressModel struct {
	progress UnifiedProgressModel
	complete bool
	toolName string
}

// NewToolsProgressModel creates a new tools progress model
func NewToolsProgressModel(progress UnifiedProgressModel, toolName string) ToolsProgressModel {
	return ToolsProgressModel{
		progress: progress,
		toolName: toolName,
	}
}

func (m ToolsProgressModel) Init() tea.Cmd {
	return tea.Batch(
		m.progress.Init(),
		m.simulateToolsOperation(),
	)
}

func (m ToolsProgressModel) simulateToolsOperation() tea.Cmd {
	return func() tea.Msg {
		// Simulate tools operation with progress updates
		m.progress.AddLog(fmt.Sprintf("Démarrage de l'opération pour %s", m.toolName))

		// Step 1: Detection
		m.progress.AddLog("Détection du système et des dépendances...")
		time.Sleep(time.Millisecond * 500)

		// Step 2: Download
		m.progress.AddLog("Téléchargement des paquets...")
		time.Sleep(time.Millisecond * 800)

		// Step 3: Installation
		m.progress.AddLog(fmt.Sprintf("Installation de %s...", m.toolName))
		time.Sleep(time.Millisecond * 1000)

		// Step 4: Verification
		m.progress.AddLog("Vérification de l'installation...")
		time.Sleep(time.Millisecond * 300)

		return ProgressFinishedMsg{
			Success: true,
			Message: fmt.Sprintf("Installation de %s terminée avec succès!", m.toolName),
		}
	}
}

func (m ToolsProgressModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc", "enter":
			if m.complete {
				return NewToolsModel(), nil
			}
		}
	case ProgressFinishedMsg:
		m.complete = true
	default:
		var cmd tea.Cmd
		m.progress, cmd = m.progress.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m ToolsProgressModel) View() string {
	if m.complete {
		var s strings.Builder
		s.WriteString(CreateBanner(" Gestion des Outils"))
		s.WriteString("\n\n")
		s.WriteString(SubtitleStyle.Render(" Opération terminée avec succès!"))
		s.WriteString("\n\n")
		s.WriteString(CardStyle.Render(m.progress.Message))
		s.WriteString("\n\n")
		s.WriteString(FooterStyle.Render("• Entrée/Échap Retour • Ctrl+C Quitter"))
		return AppStyle.Render(s.String())
	}

	return m.progress.View()
}
