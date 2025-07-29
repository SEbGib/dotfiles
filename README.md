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

### Méthode recommandée (production)
```bash
# Installation directe depuis GitHub - chezmoi gère tout automatiquement
sh -c "$(curl -fsLS get.chezmoi.io)" -- init --apply https://github.com/SEbGib/dotfiles.git
```

### Développement et modifications
```bash
# 1. Initialiser chezmoi depuis GitHub
chezmoi init https://github.com/SEbGib/dotfiles.git

# 2. Modifier les fichiers dans l'éditeur chezmoi
chezmoi edit ~/.zshrc

# 3. Appliquer les changements
chezmoi apply

# 4. Pousser les modifications vers GitHub
chezmoi cd
git add .
git commit -m "feat: update configuration"
git push
```

### Vérification de l'installation
```bash
# Vérifier que les scripts ont bien été exécutés
ls -la ~/.oh-my-zsh/  # Oh My Zsh installé
starship --version    # Starship disponible
which fzf ripgrep    # Outils modernes installés
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

Éditez `.chezmoi.toml` :
```toml
[data]
    email = "sebastien.giband@gmail.com"
    name = "Sebastien Giband"
    github_username = "SEbGib"
    personal = true
    work = false
    gpg_key_id = ""
```

**Important** : Modifiez directement via chezmoi :
```bash
# Éditer la configuration chezmoi
chezmoi edit ~/.config/chezmoi/chezmoi.toml

# Ou éditer le fichier source
chezmoi cd
vim .chezmoi.toml
chezmoi apply
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

## 🔄 Maintenance et workflow

### Workflow quotidien
```bash
# Voir les changements en attente
chezmoi diff

# Appliquer les changements depuis le repo
chezmoi apply

# Synchroniser depuis GitHub (pull + apply)
chezmoi update
```

### Modifier la configuration
```bash
# Éditer un fichier géré par chezmoi
chezmoi edit ~/.zshrc

# Voir les différences avant d'appliquer
chezmoi diff

# Appliquer les changements
chezmoi apply

# Pousser vers GitHub
chezmoi cd
git add . && git commit -m "update config" && git push
```

### Commandes utiles
```bash
# Aller dans le répertoire source chezmoi
chezmoi cd

# Voir le statut des fichiers gérés
chezmoi status

# Réinitialiser depuis GitHub (attention: perd les modifs locales)
chezmoi init --apply --force https://github.com/SEbGib/dotfiles.git

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
