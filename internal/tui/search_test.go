package tui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewSearchableMainModel(t *testing.T) {
	model := NewSearchableMainModel()

	if model.searchMode {
		t.Error("Expected search mode to be false initially")
	}

	if len(model.originalItems) == 0 {
		t.Error("Expected original items to be populated")
	}

	if len(model.filteredItems) != len(model.originalItems) {
		t.Error("Expected filtered items to match original items initially")
	}
}

func TestSearchFunctionality(t *testing.T) {
	model := NewSearchableMainModel()

	// Test entering search mode
	updatedModel, _ := model.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("/")})
	searchModel := updatedModel.(SearchableMainModel)

	if !searchModel.searchMode {
		t.Error("Expected search mode to be enabled after pressing '/'")
	}

	// Test search input
	searchModel.searchInput.SetValue("install")
	searchModel.filterItems("install")

	if len(searchModel.filteredItems) == 0 {
		t.Error("Expected to find items matching 'install'")
	}

	// Verify that filtered items contain the search term
	found := false
	for _, item := range searchModel.filteredItems {
		if strings.Contains(strings.ToLower(item.title), "install") ||
			strings.Contains(strings.ToLower(item.description), "install") {
			found = true
			break
		}
	}

	if !found {
		t.Error("Expected filtered items to contain search term")
	}
}

func TestFuzzyMatching(t *testing.T) {
	model := NewSearchableMainModel()

	// Test fuzzy matching
	testCases := []struct {
		text     string
		pattern  string
		expected bool
	}{
		{"Installation Interactive", "inst", true},
		{"Configuration Management", "config", true},
		{"System Verification", "sys", true},
		{"Installation Interactive", "xyz", false},
		{"Backup & Restore", "bkp", true}, // fuzzy match
	}

	for _, tc := range testCases {
		result := model.fuzzyMatch(strings.ToLower(tc.text), strings.ToLower(tc.pattern))
		if result != tc.expected {
			t.Errorf("fuzzyMatch(%s, %s) = %v, expected %v", tc.text, tc.pattern, result, tc.expected)
		}
	}
}

func TestSearchModeExit(t *testing.T) {
	model := NewSearchableMainModel()

	// Enter search mode
	model.searchMode = true
	model.searchInput.SetValue("test")

	// Exit search mode with Esc
	updatedModel, _ := model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	searchModel := updatedModel.(SearchableMainModel)

	if searchModel.searchMode {
		t.Error("Expected search mode to be disabled after pressing Esc")
	}

	if searchModel.searchInput.Value() != "" {
		t.Error("Expected search input to be cleared after exiting search mode")
	}

	if len(searchModel.filteredItems) != len(searchModel.originalItems) {
		t.Error("Expected filtered items to be reset to original items")
	}
}

func TestSearchView(t *testing.T) {
	model := NewSearchableMainModel()

	// Test normal view
	view := model.View()
	if !strings.Contains(view, "Dotfiles Manager") {
		t.Error("Expected view to contain title")
	}

	if !strings.Contains(view, "Rechercher") {
		t.Error("Expected view to contain search instruction")
	}

	// Test search mode view
	model.searchMode = true
	model.searchInput.SetValue("test")
	model.filterItems("test")

	searchView := model.View()
	if !strings.Contains(searchView, "Mode Recherche") {
		t.Error("Expected search view to contain search mode indicator")
	}

	if !strings.Contains(searchView, "résultat") {
		t.Error("Expected search view to show result count")
	}
}

func BenchmarkFuzzyMatch(b *testing.B) {
	model := NewSearchableMainModel()
	text := "Installation Interactive Guide"
	pattern := "inst"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.fuzzyMatch(text, pattern)
	}
}

func BenchmarkFilterItems(b *testing.B) {
	model := NewSearchableMainModel()
	query := "config"

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		model.filterItems(query)
	}
}
