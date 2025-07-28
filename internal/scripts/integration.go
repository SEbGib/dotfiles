package scripts

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// ScriptRunner handles execution of shell scripts
type ScriptRunner struct {
	baseDir string
}

// NewScriptRunner creates a new script runner
func NewScriptRunner() *ScriptRunner {
	// Get the directory where the binary is located
	baseDir, _ := os.Getwd()
	return &ScriptRunner{baseDir: baseDir}
}

// RunInstallScript executes the installation script
func (sr *ScriptRunner) RunInstallScript() error {
	scriptPath := filepath.Join(sr.baseDir, "run_once_00-install-tools.sh.tmpl")

	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script d'installation non trouvé: %s", scriptPath)
	}

	// Execute the script
	cmd := exec.Command("bash", scriptPath)
	cmd.Dir = sr.baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// RunVerificationScript executes the verification script
func (sr *ScriptRunner) RunVerificationScript() (string, error) {
	scriptPath := filepath.Join(sr.baseDir, "verify-installation.sh")

	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return "", fmt.Errorf("script de vérification non trouvé: %s", scriptPath)
	}

	// Execute the script and capture output
	cmd := exec.Command("bash", scriptPath)
	cmd.Dir = sr.baseDir

	output, err := cmd.CombinedOutput()
	return string(output), err
}

// RunBackupScript creates a backup of current configurations
func (sr *ScriptRunner) RunBackupScript() error {
	scriptPath := filepath.Join(sr.baseDir, "run_once_backup-existing-configs.sh.tmpl")

	// Check if script exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		return fmt.Errorf("script de sauvegarde non trouvé: %s", scriptPath)
	}

	// Execute the script
	cmd := exec.Command("bash", scriptPath)
	cmd.Dir = sr.baseDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// CheckCommand verifies if a command exists
func (sr *ScriptRunner) CheckCommand(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

// GetSystemInfo returns system information
func (sr *ScriptRunner) GetSystemInfo() map[string]string {
	info := make(map[string]string)

	info["OS"] = runtime.GOOS
	info["Architecture"] = runtime.GOARCH
	info["Go Version"] = runtime.Version()

	// Get shell
	if shell := os.Getenv("SHELL"); shell != "" {
		info["Shell"] = shell
	} else {
		info["Shell"] = "Unknown"
	}

	// Get home directory
	if home := os.Getenv("HOME"); home != "" {
		info["Home"] = home
	} else {
		info["Home"] = "Unknown"
	}

	// Get user
	if user := os.Getenv("USER"); user != "" {
		info["User"] = user
	} else {
		info["User"] = "Unknown"
	}

	return info
}

// ListBackups returns a list of available backups
func (sr *ScriptRunner) ListBackups() ([]string, error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return nil, fmt.Errorf("impossible de déterminer le répertoire home")
	}

	pattern := filepath.Join(homeDir, ".dotfiles-backup-*")
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	// Extract just the directory names
	var backups []string
	for _, match := range matches {
		backups = append(backups, filepath.Base(match))
	}

	return backups, nil
}

// GetInstalledTools returns a list of installed tools
func (sr *ScriptRunner) GetInstalledTools() map[string]bool {
	tools := map[string]bool{
		"chezmoi":  sr.CheckCommand("chezmoi"),
		"starship": sr.CheckCommand("starship"),
		"zsh":      sr.CheckCommand("zsh"),
		"nvim":     sr.CheckCommand("nvim"),
		"tmux":     sr.CheckCommand("tmux"),
		"git":      sr.CheckCommand("git"),
		"fzf":      sr.CheckCommand("fzf"),
		"rg":       sr.CheckCommand("rg"),
		"fd":       sr.CheckCommand("fd"),
		"bat":      sr.CheckCommand("bat"),
		"eza":      sr.CheckCommand("eza"),
		"lazygit":  sr.CheckCommand("lazygit"),
		"zoxide":   sr.CheckCommand("zoxide"),
		"node":     sr.CheckCommand("node"),
		"npm":      sr.CheckCommand("npm"),
		"php":      sr.CheckCommand("php"),
		"composer": sr.CheckCommand("composer"),
		"symfony":  sr.CheckCommand("symfony"),
		"bw":       sr.CheckCommand("bw"),
		"age":      sr.CheckCommand("age"),
		"gpg":      sr.CheckCommand("gpg"),
	}

	return tools
}

// CheckConfigFiles verifies if configuration files exist
func (sr *ScriptRunner) CheckConfigFiles() map[string]bool {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		return nil
	}

	configs := map[string]bool{
		"~/.zshrc":                sr.fileExists(filepath.Join(homeDir, ".zshrc")),
		"~/.gitconfig":            sr.fileExists(filepath.Join(homeDir, ".gitconfig")),
		"~/.aliases":              sr.fileExists(filepath.Join(homeDir, ".aliases")),
		"~/.config/starship.toml": sr.fileExists(filepath.Join(homeDir, ".config", "starship.toml")),
		"~/.config/nvim/":         sr.dirExists(filepath.Join(homeDir, ".config", "nvim")),
		"~/.config/tmux/":         sr.dirExists(filepath.Join(homeDir, ".config", "tmux")),
		"~/.oh-my-zsh/":           sr.dirExists(filepath.Join(homeDir, ".oh-my-zsh")),
		"~/.local/share/chezmoi/": sr.dirExists(filepath.Join(homeDir, ".local", "share", "chezmoi")),
	}

	return configs
}

// fileExists checks if a file exists
func (sr *ScriptRunner) fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

// dirExists checks if a directory exists
func (sr *ScriptRunner) dirExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && info.IsDir()
}

// ExecuteCommand runs a command and returns its output
func (sr *ScriptRunner) ExecuteCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	cmd.Dir = sr.baseDir

	output, err := cmd.CombinedOutput()
	return strings.TrimSpace(string(output)), err
}

// GetPackageManager detects the system package manager
func (sr *ScriptRunner) GetPackageManager() string {
	switch runtime.GOOS {
	case "darwin":
		if sr.CheckCommand("brew") {
			return "homebrew"
		}
		return "none"
	case "linux":
		if sr.CheckCommand("apt") {
			return "apt"
		} else if sr.CheckCommand("pacman") {
			return "pacman"
		} else if sr.CheckCommand("yum") {
			return "yum"
		} else if sr.CheckCommand("dnf") {
			return "dnf"
		}
		return "unknown"
	default:
		return "unsupported"
	}
}
