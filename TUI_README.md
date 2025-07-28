# 🚀 Dotfiles TUI - Interface Graphique Terminal

Une interface utilisateur moderne et interactive pour gérer vos dotfiles, construite avec [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## ✨ Fonctionnalités

### 🏠 Menu Principal
- **🚀 Installation Interactive** - Guide d'installation étape par étape
- **⚙️ Gestion de Configuration** - Éditer et gérer vos configurations
- **✅ Vérification du Système** - Vérifier l'installation et la santé
- **💾 Sauvegarde & Restauration** - Gérer les sauvegardes
- **🔧 Gestion des Outils** - Installer/mettre à jour des outils
- **🔐 Configuration des Secrets** - Gérer Bitwarden et secrets
- **📊 Informations Système** - Afficher les infos système

### 🚀 Installation Interactive
- Détection automatique du système (macOS/Linux)
- Sauvegarde automatique des configurations existantes
- Installation progressive avec indicateurs de progression
- Logs en temps réel de l'installation
- Gestion d'erreurs avec retry

### ✅ Vérification du Système
- Vérification de tous les outils installés
- Contrôle des fichiers de configuration
- Rapport détaillé avec taux de réussite
- Identification des problèmes potentiels

## 🛠️ Installation

### Prérequis
- Go 1.21 ou plus récent
- Git

### Construction
```bash
# Cloner le repository (si pas déjà fait)
git clone https://github.com/votreusername/dotfiles.git
cd dotfiles

# Installer les dépendances
make deps

# Construire l'application
make build

# Ou construire et lancer directement
make run
```

### Installation système
```bash
# Installer le binaire dans /usr/local/bin
make install

# Puis lancer depuis n'importe où
dotfiles-tui
```

## 🎮 Utilisation

### Lancement
```bash
# Depuis le dossier du projet
./dotfiles-tui

# Ou si installé système
dotfiles-tui
```

### Navigation
- **↑/↓** - Naviguer dans les menus
- **Entrée** - Sélectionner une option
- **Échap** - Retour au menu précédent
- **Ctrl+C** - Quitter l'application

### Raccourcis Globaux
- **q** - Quitter (dans certains écrans)
- **h** - Aide (si disponible)
- **r** - Rafraîchir (si disponible)

## 📁 Structure du Projet

```
dotfiles/
├── cmd/dotfiles-tui/          # Point d'entrée de l'application
│   └── main.go
├── internal/tui/              # Modèles Bubble Tea
│   ├── main.go               # Menu principal
│   ├── install.go            # Installation interactive
│   ├── verify.go             # Vérification système
│   ├── config.go             # Gestion configuration
│   ├── backup.go             # Sauvegarde/restauration
│   ├── tools.go              # Gestion des outils
│   ├── secrets.go            # Configuration secrets
│   └── info.go               # Informations système
├── go.mod                     # Dépendances Go
├── go.sum                     # Checksums des dépendances
├── Makefile                   # Commandes de build
└── TUI_README.md             # Cette documentation
```

## 🎨 Thème et Style

L'interface utilise le thème **Catppuccin** pour une cohérence avec le reste de votre environnement :

- **Couleur principale** : `#7D56F4` (Violet)
- **Succès** : `#04B575` (Vert)
- **Avertissement** : `#FFAA00` (Orange)
- **Erreur** : `#FF5555` (Rouge)
- **Texte secondaire** : `#626262` (Gris)

## 🔧 Développement

### Commandes de développement
```bash
# Formater le code
make fmt

# Linter le code
make lint

# Lancer les tests
make test

# Mode développement avec hot reload (nécessite air)
make dev
```

### Ajouter une nouvelle fonctionnalité

1. **Créer un nouveau modèle** dans `internal/tui/`
2. **Implémenter l'interface Bubble Tea** :
   ```go
   func (m YourModel) Init() tea.Cmd
   func (m YourModel) Update(msg tea.Msg) (tea.Model, tea.Cmd)
   func (m YourModel) View() string
   ```
3. **Ajouter au menu principal** dans `main.go`
4. **Tester et documenter**

### Architecture Bubble Tea

Chaque écran est un modèle Bubble Tea indépendant :
- **Init()** - Initialisation du modèle
- **Update()** - Gestion des messages/événements
- **View()** - Rendu de l'interface

## 🐛 Dépannage

### L'application ne se lance pas
```bash
# Vérifier Go
go version

# Réinstaller les dépendances
make clean
make deps
make build
```

### Erreurs de compilation
```bash
# Nettoyer et reconstruire
make clean
go mod tidy
make build
```

### Interface corrompue
- Redimensionner le terminal
- Relancer l'application
- Vérifier que le terminal supporte les couleurs

## 🤝 Contribution

1. Fork le projet
2. Créer une branche feature (`git checkout -b feature/amazing-feature`)
3. Commit vos changements (`git commit -m 'Add amazing feature'`)
4. Push vers la branche (`git push origin feature/amazing-feature`)
5. Ouvrir une Pull Request

## 📄 License

MIT License - Voir le fichier LICENSE pour plus de détails.

## 🙏 Remerciements

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Framework TUI
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - Composants TUI
- [Catppuccin](https://github.com/catppuccin/catppuccin) - Thème de couleurs

---

**Profitez de votre interface moderne pour gérer vos dotfiles ! 🎉**