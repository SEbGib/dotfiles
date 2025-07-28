package testutil

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// TestModel is a helper interface for testing Bubble Tea models
type TestModel interface {
	tea.Model
}

// ModelTester provides utilities for testing Bubble Tea models
type ModelTester struct {
	t *testing.T
}

// NewModelTester creates a new model tester
func NewModelTester(t *testing.T) *ModelTester {
	return &ModelTester{t: t}
}

// SendKey sends a key message to a model and returns the updated model
func (mt *ModelTester) SendKey(model tea.Model, key tea.KeyType) tea.Model {
	msg := tea.KeyMsg{Type: key}
	updatedModel, _ := model.Update(msg)
	return updatedModel
}

// SendKeyString sends a key string message to a model
func (mt *ModelTester) SendKeyString(model tea.Model, key string) tea.Model {
	msg := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(key)}
	updatedModel, _ := model.Update(msg)
	return updatedModel
}

// AssertViewContains checks if the model's view contains expected text
func (mt *ModelTester) AssertViewContains(model tea.Model, expected string) {
	view := model.View()
	if view == "" {
		mt.t.Error("Expected view to return non-empty string")
		return
	}

	if !containsText(view, expected) {
		mt.t.Errorf("Expected view to contain '%s', but it didn't.\nView content:\n%s", expected, view)
	}
}

// AssertViewNotContains checks if the model's view does not contain text
func (mt *ModelTester) AssertViewNotContains(model tea.Model, unexpected string) {
	view := model.View()
	if containsText(view, unexpected) {
		mt.t.Errorf("Expected view to NOT contain '%s', but it did.\nView content:\n%s", unexpected, view)
	}
}

// containsText is a helper to check if text contains a substring (case-insensitive)
func containsText(text, substr string) bool {
	return len(text) >= len(substr) &&
		(text == substr ||
			findSubstring(text, substr))
}

func findSubstring(text, substr string) bool {
	if len(substr) > len(text) {
		return false
	}

	for i := 0; i <= len(text)-len(substr); i++ {
		if text[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// TempFileHelper provides utilities for creating temporary files in tests
type TempFileHelper struct {
	t       *testing.T
	tempDir string
	files   []string
}

// NewTempFileHelper creates a new temporary file helper
func NewTempFileHelper(t *testing.T) *TempFileHelper {
	tempDir := t.TempDir()
	return &TempFileHelper{
		t:       t,
		tempDir: tempDir,
		files:   make([]string, 0),
	}
}

// CreateFile creates a temporary file with content
func (tfh *TempFileHelper) CreateFile(name, content string) string {
	filePath := filepath.Join(tfh.tempDir, name)

	// Create directory if needed
	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		tfh.t.Fatalf("Failed to create directory %s: %v", dir, err)
	}

	err := os.WriteFile(filePath, []byte(content), 0644)
	if err != nil {
		tfh.t.Fatalf("Failed to create temp file %s: %v", filePath, err)
	}

	tfh.files = append(tfh.files, filePath)
	return filePath
}

// CreateDir creates a temporary directory
func (tfh *TempFileHelper) CreateDir(name string) string {
	dirPath := filepath.Join(tfh.tempDir, name)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		tfh.t.Fatalf("Failed to create temp directory %s: %v", dirPath, err)
	}
	return dirPath
}

// GetTempDir returns the base temporary directory
func (tfh *TempFileHelper) GetTempDir() string {
	return tfh.tempDir
}

// ReadFile reads content from a file in the temp directory
func (tfh *TempFileHelper) ReadFile(name string) string {
	filePath := filepath.Join(tfh.tempDir, name)
	content, err := os.ReadFile(filePath)
	if err != nil {
		tfh.t.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	return string(content)
}

// FileExists checks if a file exists in the temp directory
func (tfh *TempFileHelper) FileExists(name string) bool {
	filePath := filepath.Join(tfh.tempDir, name)
	_, err := os.Stat(filePath)
	return err == nil
}

// MockEnvironment provides utilities for mocking environment variables
type MockEnvironment struct {
	t        *testing.T
	original map[string]string
}

// NewMockEnvironment creates a new mock environment
func NewMockEnvironment(t *testing.T) *MockEnvironment {
	return &MockEnvironment{
		t:        t,
		original: make(map[string]string),
	}
}

// SetEnv sets an environment variable and remembers the original value
func (me *MockEnvironment) SetEnv(key, value string) {
	if _, exists := me.original[key]; !exists {
		me.original[key] = os.Getenv(key)
	}
	os.Setenv(key, value)
}

// Restore restores all environment variables to their original values
func (me *MockEnvironment) Restore() {
	for key, originalValue := range me.original {
		if originalValue == "" {
			os.Unsetenv(key)
		} else {
			os.Setenv(key, originalValue)
		}
	}
}

// TestTimeout provides utilities for testing with timeouts
type TestTimeout struct {
	t       *testing.T
	timeout time.Duration
}

// NewTestTimeout creates a new test timeout helper
func NewTestTimeout(t *testing.T, timeout time.Duration) *TestTimeout {
	return &TestTimeout{
		t:       t,
		timeout: timeout,
	}
}

// RunWithTimeout runs a function with a timeout
func (tt *TestTimeout) RunWithTimeout(fn func()) {
	done := make(chan bool, 1)

	go func() {
		fn()
		done <- true
	}()

	select {
	case <-done:
		// Function completed successfully
	case <-time.After(tt.timeout):
		tt.t.Fatalf("Test timed out after %v", tt.timeout)
	}
}

// AssertEventually checks a condition repeatedly until it's true or timeout
func (tt *TestTimeout) AssertEventually(condition func() bool, message string) {
	start := time.Now()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if condition() {
				return // Success
			}
			if time.Since(start) > tt.timeout {
				tt.t.Fatalf("Condition never became true: %s (timeout: %v)", message, tt.timeout)
			}
		}
	}
}

// CommandMocker provides utilities for mocking command execution
type CommandMocker struct {
	t            *testing.T
	mockCommands map[string]MockCommandResult
}

// MockCommandResult represents the result of a mocked command
type MockCommandResult struct {
	Output string
	Error  error
}

// NewCommandMocker creates a new command mocker
func NewCommandMocker(t *testing.T) *CommandMocker {
	return &CommandMocker{
		t:            t,
		mockCommands: make(map[string]MockCommandResult),
	}
}

// MockCommand sets up a mock for a specific command
func (cm *CommandMocker) MockCommand(command string, result MockCommandResult) {
	cm.mockCommands[command] = result
}

// GetMockResult returns the mock result for a command
func (cm *CommandMocker) GetMockResult(command string) (MockCommandResult, bool) {
	result, exists := cm.mockCommands[command]
	return result, exists
}

// ClearMocks clears all command mocks
func (cm *CommandMocker) ClearMocks() {
	cm.mockCommands = make(map[string]MockCommandResult)
}
