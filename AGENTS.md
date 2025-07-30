# AGENTS.md - Dotfiles Configuration Guide

## Build/Test/Lint Commands
- **Apply changes**: `chezmoi apply`
- **Test configuration**: `chezmoi diff` (preview changes)
- **Validate templates**: `chezmoi execute-template < file.tmpl`
- **Check status**: `chezmoi status`
- **Update from remote**: `chezmoi update`
- **Test script syntax**: `bash -n script.sh.tmpl`

## Installation & Setup
- **Full setup**: Run `chezmoi init --apply` (executes all run_once scripts)
- **Manual tool install**: `./run_once_00-install-tools.sh` (installs all development tools)
- **Tools installed**: Ghostty terminal, Docker, Starship, Node.js, PHP, Python, modern CLI tools (fzf, ripgrep, bat, eza, etc.)
- **Post-install**: Restart terminal, run `chezmoi apply` for configuration

## Code Style Guidelines

### File Structure
- Use chezmoi naming conventions: `dot_filename` for dotfiles, `executable_script` for executables
- Template files end with `.tmpl` extension
- Scripts in `run_once_*` pattern for one-time setup

### Shell Scripts (Bash/Zsh)
- Use `#!/usr/bin/env bash` or `#!/usr/bin/env zsh` shebang
- Set strict mode: `set -euo pipefail`
- Use lowercase with underscores for variables: `project_path`
- Quote variables: `"$variable"` not `$variable`
- Use `[[ ]]` for conditionals, not `[ ]`

### Lua (Neovim config)
- Use 4 spaces for indentation
- Local variables: `local opt = vim.opt`
- Comment sections with `-- ===== SECTION =====`
- Group related settings together
- Use descriptive variable names

### Templates
- Use chezmoi template syntax: `{{ .variable }}`
- Conditional blocks: `{{- if condition }}...{{- end }}`
- OS-specific: `{{ if eq .chezmoi.os "darwin" }}`
- Personal/work context: `{{ if (index . "personal" | default false) }}`

### Error Handling
- Always check command existence: `command -v tool &> /dev/null`
- Provide fallbacks for missing tools
- Use descriptive error messages with emojis for user feedback
- Use safe template variable access: `{{ if (index . "variable" | default false) }}`

## Adding Neovim Plugins Workflow

### Process for Adding New Plugins
1. **Create plugin file**: `dot_config/nvim/lua/plugins/plugin-name.lua`
2. **Follow Lazy.nvim format**: Return table with plugin spec
3. **Add dependencies**: List required plugins in `dependencies` array
4. **Configure keymaps**: Either in plugin spec or update `keymaps.lua`
5. **Update theme integration**: Modify `colorscheme.lua` integrations if needed
6. **Copy to chezmoi source**: `cp file ~/.local/share/chezmoi/path`
7. **Apply changes**: `chezmoi apply`
8. **Test configuration**: `nvim --headless -c "lua require('lazy').setup('plugins')" -c "qa"`

### Neo-tree Example Implementation
- **Plugin file**: Created `neo-tree.lua` with full configuration
- **Dependencies**: plenary.nvim, nvim-web-devicons, nui.nvim
- **Keymaps added**: `<leader>e`, `<leader>E`, `<leader>ge`, `<leader>be`
- **Theme updated**: Changed `nvimtree = false, neotree = true` in colorscheme
- **Utility module**: Created `util.lua` for LSP integration helpers