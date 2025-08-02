# CRUSH.md - Dotfiles Configuration Guide

## Build/Test/Lint Commands
- **Apply changes**: `chezmoi apply`
- **Test configuration**: `chezmoi diff` (preview changes)
- **Validate templates**: `chezmoi execute-template < file.tmpl`
- **Check status**: `chezmoi status`
- **Update from remote**: `chezmoi update`
- **Test script syntax**: `bash -n script.sh.tmpl`
- **Lint shell scripts**: `shellcheck script.sh`
- **Format Lua code**: `stylua .`
- **Validate Neovim config**: `nvim --headless -c "Lazy! sync" -c "qa"`
- **Check Neovim health**: `nvim --headless -c "checkhealth" -c "qa"`

## Code Style Guidelines

### File Structure
- Use chezmoi naming conventions: `dot_filename` for dotfiles, `executable_script` for executables
- Template files end with `.tmpl` extension
- Scripts in `run_once_*` pattern for one-time setup
- Group related configurations in directories

### Shell Scripts (Bash/Zsh)
- Use `#!/usr/bin/env bash` or `#!/usr/bin/env zsh` shebang
- Set strict mode: `set -euo pipefail`
- Use lowercase with underscores for variables: `project_path`
- Quote variables: `"$variable"` not `$variable`
- Use `[[ ]]` for conditionals, not `[ ]`
- Add `# shellcheck disable=SC2207` for arrays with command substitution
- Use `readonly` for constants

### Lua (Neovim config)
- Use 4 spaces for indentation
- Local variables: `local opt = vim.opt`
- Comment sections with `-- ===== SECTION =====`
- Group related settings together
- Use descriptive variable names
- Follow snake_case naming convention
- Use single quotes for strings when possible
- Add comments for complex configurations

### Templates
- Use chezmoi template syntax: `{{ .variable }}`
- Conditional blocks: `{{- if condition }}...{{- end }}`
- OS-specific: `{{ if eq .chezmoi.os "darwin" }}`
- Personal/work context: `{{ if (index . "personal" | default false) }}`
- Use consistent spacing around template variables

### Error Handling
- Always check command existence: `command -v tool &> /dev/null`
- Provide fallbacks for missing tools
- Use descriptive error messages
- Use safe template variable access: `{{ if (index . "variable" | default false) }}
- Exit with appropriate codes: `exit 1` for errors`