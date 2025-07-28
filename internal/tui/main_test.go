package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewMainModel(t *testing.T) {
	model := NewMainModel()

	// Test initial state
	if model.quitting {
		t.Error("Expected quitting to be false initially")
	}

	if model.currentView != "main" {
		t.Errorf("Expected currentView to be 'main', got %s", model.currentView)
	}

	// Test that list has expected items
	if model.list.Items() == nil {
		t.Error("Expected list items to be initialized")
	}

	items := model.list.Items()
	if len(items) == 0 {
		t.Error("Expected list to have menu items")
	}

	// Check for essential menu items
	expectedItems := []string{
		"Installation Interactive",
		"Gestion de Configuration",
		"Vérification du Système",
		"Quitter",
	}

	itemTitles := make([]string, len(items))
	for i, item := range items {
		if menuItem, ok := item.(MenuItem); ok {
			itemTitles[i] = menuItem.title
		}
	}

	for _, expected := range expectedItems {
		found := false
		for _, title := range itemTitles {
			if strings.Contains(title, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find menu item containing '%s'", expected)
		}
	}
}

func TestMainModelUpdate(t *testing.T) {
	model := NewMainModel()

	// Test Ctrl+C quits
	updatedModel, cmd := model.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	mainModel := updatedModel.(MainModel)

	if !mainModel.quitting {
		t.Error("Expected Ctrl+C to set quitting to true")
	}

	if cmd == nil {
		t.Error("Expected Ctrl+C to return quit command")
	}
}

func TestMainModelView(t *testing.T) {
	model := NewMainModel()

	// Test normal view
	view := model.View()
	if view == "" {
		t.Error("Expected view to return non-empty string")
	}

	// Should contain title
	if !strings.Contains(view, "Dotfiles Manager") {
		t.Error("Expected view to contain title")
	}

	// Should contain navigation instructions
	if !strings.Contains(view, "Navigation") {
		t.Error("Expected view to contain navigation instructions")
	}

	// Test quitting view
	model.quitting = true
	quitView := model.View()
	if !strings.Contains(quitView, "Au revoir") {
		t.Error("Expected quitting view to contain goodbye message")
	}
}

func TestMenuItem(t *testing.T) {
	item := MenuItem{
		title:       "Test Item",
		description: "Test Description",
		action:      "test_action",
	}

	if item.Title() != "Test Item" {
		t.Errorf("Expected Title() to return 'Test Item', got %s", item.Title())
	}

	if item.Description() != "Test Description" {
		t.Errorf("Expected Description() to return 'Test Description', got %s", item.Description())
	}

	if item.FilterValue() != "Test Item" {
		t.Errorf("Expected FilterValue() to return 'Test Item', got %s", item.FilterValue())
	}
}

func TestMainModelKeyBindings(t *testing.T) {
	model := NewMainModel()

	// Test help bindings
	shortHelp := model.ShortHelp()
	if len(shortHelp) == 0 {
		t.Error("Expected short help to have key bindings")
	}

	fullHelp := model.FullHelp()
	if len(fullHelp) == 0 {
		t.Error("Expected full help to have key bindings")
	}
}

// Benchmark tests
func BenchmarkMainModelView(b *testing.B) {
	model := NewMainModel()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = model.View()
	}
}

func BenchmarkMainModelUpdate(b *testing.B) {
	model := NewMainModel()
	msg := tea.KeyMsg{Type: tea.KeyDown}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updatedModel, _ := model.Update(msg)
		model = updatedModel.(MainModel)
	}
}
