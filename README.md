# 🚀 Dotfiles modernes - Configuration développement 2024-2025

Configuration complète et moderne pour un environnement de développement PHP/Symfony et TypeScript optimisé avec Neovim, tmux, Zsh, et Starship.

## ✨ Fonctionnalités

- **🎮 Interface TUI moderne** : Interface graphique terminal avec Bubble Tea
- **🔥 Setup en une commande** : Installation complète automatisée
- **🎨 Interface moderne** : Thème Catppuccin coordonné (Neovim + tmux + Starship)
- **🔒 Gestion des secrets** : Intégration Bitwarden + chiffrement AGE
- **⚡ Performance optimisée** : Lazy loading, configurations tuned
- **🌐 Multi-plateforme** : macOS, Linux, WSL
- **🛠️ Workflows intégrés** : Scripts Symfony/TypeScript pour tmux
- **🧠 Intelligence** : Auto-complétion, suggestions, navigation améliorée

## 🚀 Installation rapide

### 🎮 Interface TUI (Recommandée)
```bash
# Cloner le repository
git clone https://github.com/votreusername/dotfiles.git
cd dotfiles

# Lancer l'interface graphique moderne
./demo-tui.sh
```

### 📟 Installation en ligne de commande
```bash
# Remplacez 'votreusername' par votre nom d'utilisateur GitHub
sh -c "$(curl -fsLS get.chezmoi.io)" -- init --apply https://github.com/votreusername/dotfiles.git
```

### 💾 Sauvegarde automatique

**Lors de la première installation, vos configurations existantes sont automatiquement sauvegardées !**

Le script détecte et sauvegarde :
- Configurations shell (`.zshrc`, `.bashrc`, etc.)
- Configuration Git (`.gitconfig`)
- Configurations d'éditeurs (`.vimrc`, `.config/nvim`)
- Configuration tmux (`.tmux.conf`)
- Oh My Zsh et plugins existants
- Configuration SSH
- Et bien plus...

La sauvegarde est créée dans `~/.dotfiles-backup-YYYYMMDD_HHMMSS/` avec :
- 📄 Tous vos fichiers de configuration
- 📋 Liste des packages installés
- ℹ️ Informations système
- 🔧 Script de restauration

### 🔄 Restauration des anciennes configurations

Si vous voulez revenir à vos anciennes configurations :

```bash
# Trouver votre sauvegarde
ls ~/.dotfiles-backup-*

# Restaurer (remplacez par le bon chemin)
cp -r ~/.dotfiles-backup-YYYYMMDD_HHMMSS/* ~/

# Ou utiliser l'alias créé
source ~/.dotfiles-backup-YYYYMMDD_HHMMSS/restore_alias.sh
restore-dotfiles-backup
```

## 📋 Processus d'installation détaillé

L'installation suit ces étapes automatiquement :

1. **💾 Sauvegarde** - Vos configurations existantes sont sauvegardées
2. **🔧 Installation des outils** - Starship, Zsh, outils modernes
3. **🐚 Configuration Zsh** - Oh My Zsh + plugins modernes
4. **📁 Création des dossiers** - Structure de développement
5. **⚙️ Application des configs** - Neovim, tmux, Git, etc.
6. **🎨 Thème Catppuccin** - Interface coordonnée
7. **🔐 Secrets Bitwarden** - Si configuré

### ⚠️ Prérequis

- **Git** installé
- **Curl** disponible
- Connexion Internet
- Permissions d'écriture dans `$HOME`

### 🎯 Installation personnalisée

Pour une installation avec options spécifiques :

```bash
# Installation non-interactive (accepte tout)
chezmoi init --apply --force https://github.com/votreusername/dotfiles.git

# Installation avec sauvegarde forcée
chezmoi init https://github.com/votreusername/dotfiles.git
chezmoi apply

# Installation en mode debug
chezmoi init --verbose https://github.com/votreusername/dotfiles.git
chezmoi apply --verbose
```

### ✅ Vérification de l'installation

Après l'installation, vérifiez que tout fonctionne correctement :

#### 🎮 Via l'interface TUI (Recommandé)
```bash
# Lancer l'interface et sélectionner "Vérification du Système"
./dotfiles-tui
```

#### 📟 Via ligne de commande
```bash
# Télécharger et exécuter le script de vérification
curl -fsSL https://raw.githubusercontent.com/votreusername/dotfiles/main/verify-installation.sh | bash

# Ou si vous avez cloné le repository
./verify-installation.sh
```

Le script vérifie :
- ✅ Tous les outils installés (Starship, Zsh, Neovim, etc.)
- ✅ Fichiers de configuration présents
- ✅ Plugins Zsh fonctionnels
- ✅ Sauvegardes créées
- ✅ Shell configuré correctement
- 📊 Rapport détaillé avec taux de réussite

## 🎮 Interface TUI - Gestion Moderne

Une interface utilisateur terminal moderne construite avec [Bubble Tea](https://github.com/charmbracelet/bubbletea) pour une gestion intuitive de vos dotfiles.

### 🚀 Lancement rapide
```bash
# Demo interactif avec toutes les options
./demo-tui.sh

# Lancement direct
./launch-tui.sh

# Ou construction manuelle
make run
```

### ✨ Fonctionnalités TUI
- **🚀 Installation Interactive** : Guide étape par étape avec progression en temps réel
- **✅ Vérification Système** : Contrôle de santé complet avec rapport détaillé
- **⚙️ Gestion Configuration** : Édition des fichiers de configuration
- **💾 Sauvegarde/Restauration** : Gestion des backups avec interface graphique
- **🔧 Gestion des Outils** : Installation et mise à jour des outils
- **🔐 Configuration Secrets** : Interface pour Bitwarden et secrets
- **📊 Informations Système** : Vue d'ensemble de votre environnement

### 🎨 Interface Moderne
- Thème **Catppuccin** coordonné avec le reste de l'environnement
- Navigation intuitive au clavier (↑/↓, Entrée, Échap)
- Indicateurs de progression en temps réel
- Logs interactifs pendant les opérations
- Messages d'erreur clairs avec suggestions

### 📚 Documentation TUI
Consultez [TUI_README.md](TUI_README.md) pour la documentation complète de l'interface.

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

## 🎯 Gestion des secrets

Configuration sécurisée avec Bitwarden :

```bash
# Configuration initiale
bw login
export BW_SESSION="$(bw unlock --raw)"
chezmoi apply
```

Les secrets sont automatiquement injectés dans :
- Configuration SSH (`~/.ssh/config`)
- Variables d'environnement (`~/.env`)
- Configuration Git (`~/.gitconfig`)

## ⌨️ Raccourcis essentiels

### Général
- `Ctrl+Space` : Accepter suggestion Zsh
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

## 🗑️ Désinstallation

Si vous souhaitez supprimer cette configuration :

### Désinstallation rapide
```bash
# Télécharger et exécuter le script de désinstallation simple
curl -fsSL https://raw.githubusercontent.com/votreusername/dotfiles/main/uninstall.sh | bash
```

### Désinstallation avancée
```bash
# Télécharger le script complet avec options
curl -fsSL https://raw.githubusercontent.com/votreusername/dotfiles/main/remove-dotfiles.sh -o remove-dotfiles.sh
chmod +x remove-dotfiles.sh

# Voir les options disponibles
./remove-dotfiles.sh --help

# Simulation (voir ce qui serait supprimé)
./remove-dotfiles.sh --dry-run

# Désinstallation interactive
./remove-dotfiles.sh

# Désinstallation automatique
./remove-dotfiles.sh --yes

# Garder certains éléments
./remove-dotfiles.sh --keep-tools --keep-shell

# Suppression complète (attention!)
./remove-dotfiles.sh --nuclear
```

### Options de désinstallation

- `--dry-run` : Mode simulation (ne supprime rien)
- `--yes` : Mode non-interactif
- `--verbose` : Affichage détaillé
- `--backup-only` : Créer seulement une sauvegarde
- `--keep-tools` : Garder les outils installés (starship, eza, etc.)
- `--keep-shell` : Ne pas restaurer le shell précédent
- `--keep-configs` : Garder les fichiers de configuration
- `--nuclear` : Suppression complète (attention!)

### Sauvegarde automatique

Les scripts de désinstallation créent automatiquement une sauvegarde dans `~/.dotfiles-removal-backup-YYYYMMDD_HHMMSS/` avant toute suppression.

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