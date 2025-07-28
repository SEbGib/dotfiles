package scripts

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestNewScriptRunner(t *testing.T) {
	runner := NewScriptRunner()

	if runner == nil {
		t.Error("Expected NewScriptRunner to return non-nil runner")
	}

	if runner.baseDir == "" {
		t.Error("Expected baseDir to be set")
	}
}

func TestCheckCommand(t *testing.T) {
	runner := NewScriptRunner()

	// Test with a command that should exist on most systems
	if !runner.CheckCommand("ls") && runtime.GOOS != "windows" {
		t.Error("Expected 'ls' command to exist on Unix systems")
	}

	// Test with a command that definitely doesn't exist
	if runner.CheckCommand("definitely_not_a_real_command_12345") {
		t.Error("Expected non-existent command to return false")
	}
}

func TestGetSystemInfo(t *testing.T) {
	runner := NewScriptRunner()
	info := runner.GetSystemInfo()

	expectedKeys := []string{"OS", "Architecture", "Go Version", "Shell", "Home", "User"}

	for _, key := range expectedKeys {
		if _, exists := info[key]; !exists {
			t.Errorf("Expected system info to contain key '%s'", key)
		}
	}

	// Test specific values
	if info["OS"] != runtime.GOOS {
		t.Errorf("Expected OS to be '%s', got '%s'", runtime.GOOS, info["OS"])
	}

	if info["Architecture"] != runtime.GOARCH {
		t.Errorf("Expected Architecture to be '%s', got '%s'", runtime.GOARCH, info["Architecture"])
	}

	if !strings.Contains(info["Go Version"], "go") {
		t.Errorf("Expected Go Version to contain 'go', got '%s'", info["Go Version"])
	}
}

func TestGetInstalledTools(t *testing.T) {
	runner := NewScriptRunner()
	tools := runner.GetInstalledTools()

	expectedTools := []string{
		"git", "zsh", "nvim", "tmux", "fzf", "rg", "fd", "bat", "eza",
		"lazygit", "zoxide", "node", "npm", "php", "composer", "symfony",
		"bw", "age", "gpg", "chezmoi", "starship",
	}

	for _, tool := range expectedTools {
		if _, exists := tools[tool]; !exists {
			t.Errorf("Expected tools map to contain '%s'", tool)
		}
	}

	// At least git should be available in most environments
	if !tools["git"] {
		t.Log("Warning: git not found - this might be expected in some test environments")
	}
}

func TestCheckConfigFiles(t *testing.T) {
	runner := NewScriptRunner()
	configs := runner.CheckConfigFiles()

	if configs == nil {
		// This might happen if HOME is not set
		if os.Getenv("HOME") == "" {
			t.Skip("Skipping config files test - HOME not set")
		}
		t.Error("Expected configs map to be non-nil")
		return
	}

	expectedConfigs := []string{
		"~/.zshrc", "~/.gitconfig", "~/.aliases", "~/.config/starship.toml",
		"~/.config/nvim/", "~/.config/tmux/", "~/.oh-my-zsh/", "~/.local/share/chezmoi/",
	}

	for _, config := range expectedConfigs {
		if _, exists := configs[config]; !exists {
			t.Errorf("Expected configs map to contain '%s'", config)
		}
	}
}

func TestFileExists(t *testing.T) {
	runner := NewScriptRunner()

	// Create a temporary file
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_file.txt")

	// File shouldn't exist initially
	if runner.fileExists(tempFile) {
		t.Error("Expected fileExists to return false for non-existent file")
	}

	// Create the file
	err := os.WriteFile(tempFile, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Now it should exist
	if !runner.fileExists(tempFile) {
		t.Error("Expected fileExists to return true for existing file")
	}

	// Test with directory (should return false)
	if runner.fileExists(tempDir) {
		t.Error("Expected fileExists to return false for directory")
	}
}

func TestDirExists(t *testing.T) {
	runner := NewScriptRunner()

	// Create a temporary directory
	tempDir := t.TempDir()
	tempSubDir := filepath.Join(tempDir, "test_subdir")

	// Directory shouldn't exist initially
	if runner.dirExists(tempSubDir) {
		t.Error("Expected dirExists to return false for non-existent directory")
	}

	// Create the directory
	err := os.Mkdir(tempSubDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create test directory: %v", err)
	}

	// Now it should exist
	if !runner.dirExists(tempSubDir) {
		t.Error("Expected dirExists to return true for existing directory")
	}

	// Test with file (should return false)
	tempFile := filepath.Join(tempDir, "test_file.txt")
	err = os.WriteFile(tempFile, []byte("test"), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	if runner.dirExists(tempFile) {
		t.Error("Expected dirExists to return false for file")
	}
}

func TestExecuteCommand(t *testing.T) {
	runner := NewScriptRunner()

	// Test a simple command that should work on most systems
	var cmd, arg string
	if runtime.GOOS == "windows" {
		cmd, arg = "cmd", "/c echo test"
	} else {
		cmd, arg = "echo", "test"
	}

	output, err := runner.ExecuteCommand(cmd, arg)
	if err != nil {
		t.Errorf("Expected command to succeed, got error: %v", err)
	}

	if !strings.Contains(output, "test") {
		t.Errorf("Expected output to contain 'test', got: %s", output)
	}
}

func TestGetPackageManager(t *testing.T) {
	runner := NewScriptRunner()
	pm := runner.GetPackageManager()

	switch runtime.GOOS {
	case "darwin":
		// On macOS, should detect homebrew if available, otherwise "none"
		if pm != "homebrew" && pm != "none" {
			t.Errorf("Expected package manager to be 'homebrew' or 'none' on macOS, got: %s", pm)
		}
	case "linux":
		// On Linux, should detect one of the common package managers or "unknown"
		validPMs := []string{"apt", "pacman", "yum", "dnf", "unknown"}
		valid := false
		for _, validPM := range validPMs {
			if pm == validPM {
				valid = true
				break
			}
		}
		if !valid {
			t.Errorf("Expected valid Linux package manager, got: %s", pm)
		}
	default:
		if pm != "unsupported" {
			t.Errorf("Expected 'unsupported' for unknown OS, got: %s", pm)
		}
	}
}

func TestListBackups(t *testing.T) {
	runner := NewScriptRunner()

	// This test might not find any backups, which is fine
	backups, err := runner.ListBackups()
	if err != nil {
		// Error might occur if HOME is not set
		if os.Getenv("HOME") == "" {
			t.Skip("Skipping backup test - HOME not set")
		}
		t.Errorf("Expected ListBackups to not error, got: %v", err)
	}

	// backups can be empty, that's fine
	if backups == nil {
		t.Error("Expected backups slice to be non-nil")
	}
}

// Integration test that combines multiple functions
func TestScriptRunnerIntegration(t *testing.T) {
	runner := NewScriptRunner()

	// Get system info
	info := runner.GetSystemInfo()
	if len(info) == 0 {
		t.Error("Expected system info to contain data")
	}

	// Check some tools
	tools := runner.GetInstalledTools()
	if len(tools) == 0 {
		t.Error("Expected tools map to contain data")
	}

	// Get package manager
	pm := runner.GetPackageManager()
	if pm == "" {
		t.Error("Expected package manager to be detected")
	}

	// Check if git is available (common tool)
	gitAvailable := runner.CheckCommand("git")
	gitInTools := tools["git"]

	if gitAvailable != gitInTools {
		t.Error("Expected CheckCommand and GetInstalledTools to agree on git availability")
	}
}

// Benchmark tests
func BenchmarkGetSystemInfo(b *testing.B) {
	runner := NewScriptRunner()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = runner.GetSystemInfo()
	}
}

func BenchmarkGetInstalledTools(b *testing.B) {
	runner := NewScriptRunner()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = runner.GetInstalledTools()
	}
}

func BenchmarkCheckCommand(b *testing.B) {
	runner := NewScriptRunner()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = runner.CheckCommand("git")
	}
}
