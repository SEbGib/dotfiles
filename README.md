# 🚀 Dotfiles modernes - Configuration développement 2024-2025

Configuration complète et moderne pour un environnement de développement PHP/Symfony et TypeScript optimisé avec Neovim, tmux, Zsh, et Starship.

## ✨ Fonctionnalités

- **🔥 Setup en une commande** : Installation complète automatisée
- **🎨 Interface moderne** : Thème Catppuccin coordonné (Neovim + tmux + Starship)
- **🔒 Gestion des secrets** : Intégration Bitwarden + chiffrement AGE
- **⚡ Performance optimisée** : Lazy loading, configurations tuned
- **🌐 Multi-plateforme** : macOS, Linux, WSL
- **🛠️ Workflows intégrés** : Scripts Symfony/TypeScript pour tmux
- **🧠 Intelligence** : Auto-complétion, suggestions, navigation améliorée

## 🚀 Installation rapide

```bash
# Remplacez 'votreusername' par votre nom d'utilisateur GitHub
sh -c "$(curl -fsLS get.chezmoi.io)" -- init --apply https://github.com/votreusername/dotfiles.git
```

## 📦 Ce qui est installé

### Outils essentiels
- **Starship** : Prompt moderne et rapide
- **Zsh + Oh My Zsh** : Shell avancé avec plugins
- **Neovim** : Éditeur moderne avec LSP complet
- **tmux** : Multiplexeur terminal avec scripts automatisés

### Plugins Zsh
- **zsh-autosuggestions** : Suggestions intelligentes
- **fast-syntax-highlighting** : Coloration syntaxique rapide
- **zsh-completions** : Complétion avancée
- **zsh-history-substring-search** : Recherche dans l'historique
- **zsh-z** : Navigation rapide avec `z`

### Outils modernes
- **fzf** : Recherche floue interactive
- **ripgrep** : Recherche rapide dans fichiers
- **fd** : Alternative moderne à `find`
- **bat** : `cat` avec coloration syntaxique
- **eza** : `ls` moderne avec icônes
- **lazygit** : Interface Git élégante

### Développement
- **Node.js + npm** : Runtime JavaScript
- **PHP + Composer** : Développement PHP
- **Symfony CLI** : Outils Symfony
- **GitHub CLI** : Intégration GitHub



## ⌨️ Raccourcis essentiels

### Général
- `↑/↓` : Recherche dans l'historique
- `Ctrl+R` : Recherche fuzzy avec fzf

### Tmux (Prefix: `Ctrl+Space`)
- `|` : Split vertical
- `-` : Split horizontal
- `Alt+hjkl` : Navigation entre panes (sans prefix)
- `f` : Sessionizer avec fzf
- `Ctrl+S` : Workflow Symfony
- `Ctrl+T` : Workflow TypeScript

### Navigation
- `z <dossier>` : Navigation rapide
- `ts` : Tmux sessionizer
- `dev-symfony` : Démarrage session Symfony
- `dev-ts` : Démarrage session TypeScript

### Git
- `lg` : Lazygit
- `gst` : Git status
- `ga` : Git add
- `gc` : Git commit
- `gp` : Git push

## 🔧 Personnalisation

### Modifier les informations personnelles

Éditez `.chezmoi.toml.tmpl` :
```toml
[data]
    email = "votre@email.com"
    name = "Votre Nom"
    github_username = "votreusername"
```

### Ajouter des secrets Bitwarden

1. Créez les items dans votre vault Bitwarden
2. Utilisez dans les templates : `{{ (bitwarden "item" "Nom Item").login.password }}`
3. Appliquez : `chezmoi apply`

### Configuration locale

Créez des fichiers de surcharge :
- `~/.zshrc.local` : Configuration Zsh locale
- `~/.gitconfig.local` : Configuration Git locale
- `~/.config/tmux/tmux.conf.local` : Configuration tmux locale

## 🛠️ Workflows de développement

### Projet Symfony
```bash
dev-symfony mon-projet
# → Crée session tmux avec :
#   - Éditeur Neovim
#   - Serveur Symfony
#   - Console/DB
#   - Tests + logs
#   - Git (lazygit)
```

### Projet TypeScript
```bash
dev-ts mon-projet
# → Crée session tmux avec :
#   - Éditeur Neovim
#   - Dev server + TypeScript watch
#   - Tests + linting
#   - Build
#   - Git (lazygit)
```

## 📱 Multi-machines

Configuration intelligente selon :
- **OS** : macOS vs Linux
- **Contexte** : Personnel vs Entreprise
- **Hostname** : Configuration spécifique par machine

## 🔄 Maintenance

```bash
# Voir les changements
chezmoi diff

# Appliquer les mises à jour
chezmoi apply

# Synchroniser depuis GitHub
chezmoi update

# Mise à jour complète du système
update_all  # Alias configuré
```

## 🎨 Thèmes

Thème **Catppuccin Mocha** coordonné :
- Starship prompt
- Neovim interface
- tmux statusline
- Zsh syntax highlighting
- FZF interface

## 🐛 Dépannage

### Zsh lent au démarrage
```bash
# Profiler le démarrage
zmodload zsh/zprof
# Ajouter à .zshrc, puis à la fin :
zprof
```

### Plugins Zsh manquants
```bash
# Ré-installer les plugins
~/.local/share/chezmoi/run_once_01-setup-zsh-plugins.sh
```

### Starship ne s'affiche pas
```bash
# Vérifier installation
starship --version

# Réinstaller si nécessaire
curl -sS https://starship.rs/install.sh | sh
```

## 📚 Documentation

- [Chezmoi](https://www.chezmoi.io/) : Gestionnaire de dotfiles
- [Starship](https://starship.rs/) : Prompt cross-shell
- [Neovim](https://neovim.io/) : Éditeur moderne
- [tmux](https://github.com/tmux/tmux) : Multiplexeur terminal

## 🤝 Contribution

1. Fork le projet
2. Adaptez selon vos besoins
3. Partagez vos améliorations !

## 📄 License

MIT License - Utilisez librement !

---

**Profitez de votre environnement de développement moderne ! 🎉**
