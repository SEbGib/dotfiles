package tui

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestNewEditorModel(t *testing.T) {
	model := NewEditorModel("$HOME/.zshrc", ".zshrc")

	if model.fileName != ".zshrc" {
		t.Errorf("Expected fileName to be '.zshrc', got %s", model.fileName)
	}

	// Should expand environment variables
	if strings.Contains(model.filePath, "$HOME") {
		t.Error("Expected filePath to have $HOME expanded")
	}

	if model.message == "" {
		t.Error("Expected initial message to be set")
	}
}

func TestEditorModelWithTempFile(t *testing.T) {
	// Create a temporary file for testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test.conf")
	testContent := "# Test configuration\ntest_setting=value\n"

	err := os.WriteFile(tempFile, []byte(testContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}

	model := NewEditorModel(tempFile, "test.conf")

	// Test file loading
	cmd := model.loadFile()
	if cmd == nil {
		t.Error("Expected loadFile to return a command")
	}

	// Execute the command to simulate loading
	msg := cmd()
	if loadMsg, ok := msg.(fileLoadedMsg); ok {
		if loadMsg.error != nil {
			t.Errorf("Expected no error loading file, got: %v", loadMsg.error)
		}
		if loadMsg.content != testContent {
			t.Errorf("Expected content to match, got: %s", loadMsg.content)
		}
	} else {
		t.Error("Expected fileLoadedMsg from loadFile command")
	}
}

func TestEditorModelUpdate(t *testing.T) {
	model := NewEditorModel("/tmp/test.conf", "test.conf")

	// Test Ctrl+C quits
	updatedModel, cmd := model.Update(tea.KeyMsg{Type: tea.KeyCtrlC})
	if cmd == nil {
		t.Error("Expected Ctrl+C to return quit command")
	}

	// Test Esc returns to main menu
	updatedModel, _ = model.Update(tea.KeyMsg{Type: tea.KeyEsc})
	if _, ok := updatedModel.(TwoColumnMainModel); !ok {
		t.Error("Expected Esc to return TwoColumnMainModel")
	}

	// Test file loaded message
	testContent := "test content"
	loadMsg := fileLoadedMsg{content: testContent, error: nil}
	updatedModel, _ = model.Update(loadMsg)

	if editorModel, ok := updatedModel.(EditorModel); ok {
		if editorModel.content != testContent {
			t.Errorf("Expected content to be updated to '%s', got '%s'", testContent, editorModel.content)
		}
	} else {
		t.Error("Expected EditorModel after file loaded")
	}
}

func TestEditorModelView(t *testing.T) {
	model := NewEditorModel("/tmp/test.conf", "test.conf")
	model.content = "# Test content\nline1=value1\nline2=value2\n"

	view := model.View()

	if view == "" {
		t.Error("Expected view to return non-empty string")
	}

	// Should contain file name
	if !strings.Contains(view, "test.conf") {
		t.Error("Expected view to contain file name")
	}

	// Should contain instructions
	if !strings.Contains(view, "Instructions") {
		t.Error("Expected view to contain instructions")
	}

	// Should contain content preview
	if !strings.Contains(view, "Aperçu") {
		t.Error("Expected view to contain content preview")
	}
}

func TestEditorModelSaveFile(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_save.conf")

	model := NewEditorModel(tempFile, "test_save.conf")
	model.content = "# Test save content\nsave_test=true\n"

	// Test save command
	cmd := model.saveFile()
	if cmd == nil {
		t.Error("Expected saveFile to return a command")
	}

	// Execute the save command
	msg := cmd()
	if saveMsg, ok := msg.(fileSavedMsg); ok {
		if !saveMsg.success {
			t.Errorf("Expected save to succeed, got error: %v", saveMsg.error)
		}
	} else {
		t.Error("Expected fileSavedMsg from saveFile command")
	}

	// Verify file was actually saved
	savedContent, err := os.ReadFile(tempFile)
	if err != nil {
		t.Errorf("Failed to read saved file: %v", err)
	}

	if string(savedContent) != model.content {
		t.Errorf("Expected saved content to match model content")
	}
}

func TestEditorModelNonExistentFile(t *testing.T) {
	model := NewEditorModel("/tmp/nonexistent.conf", "nonexistent.conf")

	// Test loading non-existent file
	cmd := model.loadFile()
	msg := cmd()

	if loadMsg, ok := msg.(fileLoadedMsg); ok {
		// Should create default content for non-existent files
		if loadMsg.error != nil {
			t.Errorf("Expected no error for non-existent file, got: %v", loadMsg.error)
		}
		if !strings.Contains(loadMsg.content, "Nouveau fichier") {
			t.Error("Expected default content for new file")
		}
	} else {
		t.Error("Expected fileLoadedMsg from loadFile command")
	}
}

// Integration test for the full editor workflow
func TestEditorWorkflow(t *testing.T) {
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "workflow_test.conf")

	// Step 1: Create editor model
	model := NewEditorModel(tempFile, "workflow_test.conf")

	// Step 2: Initialize (should trigger file loading)
	cmd := model.Init()
	if cmd == nil {
		t.Error("Expected Init to return a command")
	}

	// Step 3: Simulate file loading (non-existent file)
	loadCmd := model.loadFile()
	loadMsg := loadCmd()

	// Step 4: Update model with loaded content
	updatedModel, _ := model.Update(loadMsg)
	model = updatedModel.(EditorModel)

	// Step 5: Modify content
	model.content = "# Modified content\ntest=modified\n"

	// Step 6: Save file
	saveCmd := model.saveFile()
	saveMsg := saveCmd()

	// Step 7: Update model with save result
	finalModel, _ := model.Update(saveMsg)
	model = finalModel.(EditorModel)

	// Verify the workflow completed successfully
	if !model.saved {
		t.Error("Expected file to be marked as saved")
	}

	// Verify file exists and has correct content
	if _, err := os.Stat(tempFile); os.IsNotExist(err) {
		t.Error("Expected file to exist after save")
	}
}

// Benchmark tests
func BenchmarkEditorModelView(b *testing.B) {
	model := NewEditorModel("/tmp/bench.conf", "bench.conf")
	model.content = strings.Repeat("line content\n", 100)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = model.View()
	}
}

func BenchmarkEditorModelUpdate(b *testing.B) {
	model := NewEditorModel("/tmp/bench.conf", "bench.conf")
	msg := tea.KeyMsg{Type: tea.KeyDown}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		updatedModel, _ := model.Update(msg)
		if editorModel, ok := updatedModel.(EditorModel); ok {
			model = editorModel
		}
	}
}
