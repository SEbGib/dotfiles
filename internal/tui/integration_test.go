package tui

import (
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/sebastiengiband/dotfiles/internal/testutil"
)

// TestMainMenuNavigation tests the complete navigation flow
func TestMainMenuNavigation(t *testing.T) {
	tester := testutil.NewModelTester(t)
	model := NewTwoColumnMainModel()

	// Test initial state
	tester.AssertViewContains(model, "Dotfiles Manager")
	tester.AssertViewContains(model, "Navigation")

	// Test navigation down
	updatedModel := tester.SendKey(model, tea.KeyDown)
	if mainModel, ok := updatedModel.(TwoColumnMainModel); ok {
		model = mainModel
	}
	view := model.View()
	if view == "" {
		t.Error("Expected view after navigation")
	}

	// Test entering configuration menu
	updatedModel = tester.SendKey(model, tea.KeyEnter)

	// Should transition to config model or stay in main (depending on selection)
	// This tests the model switching logic
	if updatedModel == nil {
		t.Error("Expected model to remain valid after key press")
	}
}

// TestEditorIntegration tests the complete editor workflow
func TestEditorIntegration(t *testing.T) {
	fileHelper := testutil.NewTempFileHelper(t)
	tester := testutil.NewModelTester(t)
	timeout := testutil.NewTestTimeout(t, 5*time.Second)

	// Create a test configuration file
	testContent := "# Test configuration\ntest_setting=value\n"
	testFile := fileHelper.CreateFile("test.conf", testContent)

	// Create editor model
	model := NewEditorModel(testFile, "test.conf")

	// Test initial view
	tester.AssertViewContains(model, "test.conf")
	tester.AssertViewContains(model, "Instructions")

	// Test file loading with timeout
	timeout.RunWithTimeout(func() {
		// Initialize the model (should trigger file loading)
		cmd := model.Init()
		if cmd != nil {
			// Execute the load command
			msg := cmd()
			if loadMsg, ok := msg.(fileLoadedMsg); ok {
				// Update model with loaded content
				updatedModel, _ := model.Update(loadMsg)
				model = updatedModel.(EditorModel)
			}
		}
	})

	// Verify content was loaded
	if model.content != testContent {
		t.Errorf("Expected content to be loaded, got: %s", model.content)
	}

	// Test view after loading
	tester.AssertViewContains(model, "Aperçu")
	tester.AssertViewContains(model, "test_setting")

	// Test navigation back to main menu
	updatedModel := tester.SendKey(model, tea.KeyEsc)
	if _, ok := updatedModel.(TwoColumnMainModel); !ok {
		t.Error("Expected Esc to return to main menu")
	}
}

// TestConfigurationFlow tests the complete configuration management flow
func TestConfigurationFlow(t *testing.T) {
	fileHelper := testutil.NewTempFileHelper(t)
	tester := testutil.NewModelTester(t)
	mockEnv := testutil.NewMockEnvironment(t)
	defer mockEnv.Restore()

	// Mock HOME environment variable
	mockEnv.SetEnv("HOME", fileHelper.GetTempDir())

	// Create mock configuration files
	fileHelper.CreateFile(".zshrc", "# Zsh configuration\nexport PATH=$PATH:/usr/local/bin\n")
	fileHelper.CreateFile(".gitconfig", "[user]\n\tname = Test User\n\temail = test@example.com\n")

	// Start with main menu
	mainModel := NewTwoColumnMainModel()
	tester.AssertViewContains(mainModel, "Dotfiles Manager")

	// Navigate to configuration menu (assuming it's the second item)
	configModel := NewConfigModel()
	tester.AssertViewContains(configModel, "Gestion de Configuration")
	tester.AssertViewContains(configModel, ".zshrc")
	tester.AssertViewContains(configModel, ".gitconfig")

	// Test navigation within config menu
	configModel = tester.SendKey(configModel, tea.KeyDown).(ConfigModel)
	configModel = tester.SendKey(configModel, tea.KeyUp).(ConfigModel)

	// Test returning to main menu
	mainModel = tester.SendKey(configModel, tea.KeyEsc).(TwoColumnMainModel)
	tester.AssertViewContains(mainModel, "Dotfiles Manager")
}

// TestVerificationFlow tests the system verification workflow
func TestVerificationFlow(t *testing.T) {
	tester := testutil.NewModelTester(t)
	timeout := testutil.NewTestTimeout(t, 10*time.Second)

	model := NewVerifyModel()

	// Test initial state
	tester.AssertViewContains(model, "Vérification du Système")

	// Initialize the model (should start verification)
	cmd := model.Init()
	if cmd == nil {
		t.Error("Expected Init to return a command")
	}

	// Simulate the verification process with timeout
	timeout.RunWithTimeout(func() {
		// Run a few verification steps
		for i := 0; i < 3 && !model.complete; i++ {
			if model.current < len(model.checks) {
				// Simulate a check completion
				checkMsg := checkCompleteMsg{
					index:   model.current,
					status:  "passed",
					message: "Test check passed",
				}
				updatedModel, _ := model.Update(checkMsg)
				model = updatedModel.(VerifyModel)
			}
		}
	})

	// Test that some checks were processed
	if model.current == 0 {
		t.Error("Expected some verification checks to be processed")
	}

	// Test view during verification
	view := model.View()
	if view == "" {
		t.Error("Expected non-empty view during verification")
	}
}

// TestErrorHandling tests error scenarios across different models
func TestErrorHandling(t *testing.T) {
	tester := testutil.NewModelTester(t)

	// Test editor with non-existent file
	editorModel := NewEditorModel("/nonexistent/path/file.conf", "file.conf")

	// Test that it handles the error gracefully
	cmd := editorModel.loadFile()
	msg := cmd()

	if loadMsg, ok := msg.(fileLoadedMsg); ok {
		updatedModel, _ := editorModel.Update(loadMsg)
		editorModel = updatedModel.(EditorModel)

		// Should have default content for non-existent file
		if editorModel.content == "" {
			t.Error("Expected default content for non-existent file")
		}

		tester.AssertViewContains(editorModel, "file.conf")
	}

	// Test quit functionality across models
	models := []tea.Model{
		NewTwoColumnMainModel(),
		NewConfigModel(),
		NewVerifyModel(),
		editorModel,
	}

	for _, model := range models {
		updatedModel := tester.SendKey(model, tea.KeyCtrlC)
		// All models should handle Ctrl+C gracefully
		if updatedModel == nil {
			t.Error("Expected model to handle Ctrl+C without panicking")
		}
	}
}

// TestModelTransitions tests transitions between different models
func TestModelTransitions(t *testing.T) {
	tester := testutil.NewModelTester(t)

	// Start with main menu
	mainModel := NewTwoColumnMainModel()
	tester.AssertViewContains(mainModel, "Dotfiles Manager")

	// Test transition to config model
	configModel := NewConfigModel()
	tester.AssertViewContains(configModel, "Configuration")

	// Test transition back to main
	backToMain := tester.SendKey(configModel, tea.KeyEsc)
	if _, ok := backToMain.(TwoColumnMainModel); !ok {
		t.Error("Expected transition back to TwoColumnMainModel")
	}

	// Test transition to editor
	editorModel := NewEditorModel("/tmp/test.conf", "test.conf")
	tester.AssertViewContains(editorModel, "test.conf")

	// Test transition back to main from editor
	backFromEditor := tester.SendKey(editorModel, tea.KeyEsc)
	if _, ok := backFromEditor.(TwoColumnMainModel); !ok {
		t.Error("Expected transition back to TwoColumnMainModel from editor")
	}
}

// TestConcurrentOperations tests that models handle concurrent operations safely
func TestConcurrentOperations(t *testing.T) {
	timeout := testutil.NewTestTimeout(t, 5*time.Second)

	// Test that multiple models can be created and used concurrently
	timeout.RunWithTimeout(func() {
		models := make([]tea.Model, 10)

		// Create multiple models concurrently
		done := make(chan bool, 10)

		for i := 0; i < 10; i++ {
			go func(index int) {
				switch index % 4 {
				case 0:
					models[index] = NewTwoColumnMainModel()
				case 1:
					models[index] = NewConfigModel()
				case 2:
					models[index] = NewVerifyModel()
				case 3:
					models[index] = NewEditorModel("/tmp/test.conf", "test.conf")
				}
				done <- true
			}(i)
		}

		// Wait for all models to be created
		for i := 0; i < 10; i++ {
			<-done
		}

		// Verify all models were created successfully
		for i, model := range models {
			if model == nil {
				t.Errorf("Model %d was not created", i)
			}
		}
	})
}

// BenchmarkModelOperations benchmarks common model operations
func BenchmarkModelOperations(b *testing.B) {
	models := []tea.Model{
		NewTwoColumnMainModel(),
		NewConfigModel(),
		NewVerifyModel(),
		NewEditorModel("/tmp/bench.conf", "bench.conf"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for _, model := range models {
			// Benchmark view rendering
			_ = model.View()

			// Benchmark key handling
			_, _ = model.Update(tea.KeyMsg{Type: tea.KeyDown})
		}
	}
}
