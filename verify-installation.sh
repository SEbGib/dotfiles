#!/bin/bash

# Script de vérification de l'installation des dotfiles
# Vérifie que tous les composants sont correctement installés

set -euo pipefail

# Couleurs
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
CYAN='\033[0;36m'
NC='\033[0m'

# Compteurs
CHECKS_PASSED=0
CHECKS_FAILED=0
TOTAL_CHECKS=0

print_header() {
    echo -e "${BLUE}╭─────────────────────────────────────────────────────────────╮${NC}"
    echo -e "${BLUE}│${NC}  🔍  ${CYAN}Vérification de l'installation des dotfiles${NC}  🔍   ${BLUE}│${NC}"
    echo -e "${BLUE}╰─────────────────────────────────────────────────────────────╯${NC}"
    echo ""
}

check_command() {
    local cmd="$1"
    local description="$2"
    local optional="${3:-false}"
    
    ((TOTAL_CHECKS++))
    
    if command -v "$cmd" &> /dev/null; then
        local version=""
        case "$cmd" in
            "starship") version="$(starship --version | head -1)" ;;
            "nvim") version="$(nvim --version | head -1)" ;;
            "tmux") version="$(tmux -V)" ;;
            "git") version="$(git --version)" ;;
            "zsh") version="$(zsh --version)" ;;
            *) version="$(command -v "$cmd")" ;;
        esac
        echo -e "${GREEN}✅${NC} $description - $version"
        ((CHECKS_PASSED++))
        return 0
    else
        if [[ "$optional" == "true" ]]; then
            echo -e "${YELLOW}⚠️${NC} $description - Optionnel, non installé"
        else
            echo -e "${RED}❌${NC} $description - Non trouvé"
            ((CHECKS_FAILED++))
        fi
        return 1
    fi
}

check_file() {
    local file="$1"
    local description="$2"
    local optional="${3:-false}"
    
    ((TOTAL_CHECKS++))
    
    if [[ -f "$file" ]]; then
        echo -e "${GREEN}✅${NC} $description - $file"
        ((CHECKS_PASSED++))
        return 0
    else
        if [[ "$optional" == "true" ]]; then
            echo -e "${YELLOW}⚠️${NC} $description - Optionnel, non trouvé"
        else
            echo -e "${RED}❌${NC} $description - $file non trouvé"
            ((CHECKS_FAILED++))
        fi
        return 1
    fi
}

check_directory() {
    local dir="$1"
    local description="$2"
    local optional="${3:-false}"
    
    ((TOTAL_CHECKS++))
    
    if [[ -d "$dir" ]]; then
        local count=$(find "$dir" -type f 2>/dev/null | wc -l | tr -d ' ')
        echo -e "${GREEN}✅${NC} $description - $dir ($count fichiers)"
        ((CHECKS_PASSED++))
        return 0
    else
        if [[ "$optional" == "true" ]]; then
            echo -e "${YELLOW}⚠️${NC} $description - Optionnel, non trouvé"
        else
            echo -e "${RED}❌${NC} $description - $dir non trouvé"
            ((CHECKS_FAILED++))
        fi
        return 1
    fi
}

main() {
    print_header
    
    echo -e "${CYAN}🔍 Vérification des outils essentiels...${NC}"
    check_command "chezmoi" "Chezmoi"
    check_command "starship" "Starship"
    check_command "zsh" "Zsh"
    check_command "nvim" "Neovim"
    check_command "tmux" "tmux"
    check_command "git" "Git"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification des outils modernes...${NC}"
    check_command "fzf" "FZF"
    check_command "rg" "Ripgrep"
    check_command "fd" "fd"
    check_command "bat" "bat"
    check_command "eza" "eza"
    check_command "lazygit" "Lazygit" "true"
    check_command "zoxide" "Zoxide" "true"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification des fichiers de configuration...${NC}"
    check_file "$HOME/.zshrc" "Configuration Zsh"
    check_file "$HOME/.gitconfig" "Configuration Git"
    check_file "$HOME/.aliases" "Aliases personnalisés"
    check_file "$HOME/.config/starship.toml" "Configuration Starship"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification des dossiers de configuration...${NC}"
    check_directory "$HOME/.config/nvim" "Configuration Neovim"
    check_directory "$HOME/.config/tmux" "Configuration tmux"
    check_directory "$HOME/.oh-my-zsh" "Oh My Zsh"
    check_directory "$HOME/.local/share/chezmoi" "Repository Chezmoi"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification des plugins Zsh...${NC}"
    check_directory "$HOME/.oh-my-zsh/custom/plugins/zsh-autosuggestions" "zsh-autosuggestions"
    check_directory "$HOME/.oh-my-zsh/custom/plugins/fast-syntax-highlighting" "fast-syntax-highlighting"
    check_directory "$HOME/.oh-my-zsh/custom/plugins/zsh-completions" "zsh-completions"
    check_directory "$HOME/.oh-my-zsh/custom/plugins/zsh-history-substring-search" "zsh-history-substring-search"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification des outils de développement...${NC}"
    check_command "node" "Node.js" "true"
    check_command "npm" "npm" "true"
    check_command "php" "PHP" "true"
    check_command "composer" "Composer" "true"
    check_command "symfony" "Symfony CLI" "true"
    echo ""
    
    echo -e "${CYAN}🔍 Vérification de la sécurité...${NC}"
    check_command "bw" "Bitwarden CLI" "true"
    check_command "age" "AGE encryption" "true"
    check_command "gpg" "GPG" "true"
    echo ""
    
    # Vérification du shell actuel
    echo -e "${CYAN}🔍 Vérification du shell...${NC}"
    ((TOTAL_CHECKS++))
    if [[ "$SHELL" == *"zsh"* ]]; then
        echo -e "${GREEN}✅${NC} Shell actuel - $SHELL"
        ((CHECKS_PASSED++))
    else
        echo -e "${YELLOW}⚠️${NC} Shell actuel - $SHELL (Zsh recommandé)"
    fi
    echo ""
    
    # Vérification des sauvegardes
    echo -e "${CYAN}🔍 Vérification des sauvegardes...${NC}"
    if ls ~/.dotfiles-backup-* 1> /dev/null 2>&1; then
        local backup_count=$(ls -1d ~/.dotfiles-backup-* | wc -l | tr -d ' ')
        echo -e "${GREEN}✅${NC} Sauvegardes trouvées - $backup_count sauvegarde(s)"
        ls -1d ~/.dotfiles-backup-* | while read backup; do
            echo "    📁 $backup"
        done
    else
        echo -e "${YELLOW}⚠️${NC} Aucune sauvegarde trouvée - Installation propre"
    fi
    echo ""
    
    # Rapport final
    echo -e "${BLUE}╭─────────────────────────────────────────────────────────────╮${NC}"
    echo -e "${BLUE}│${NC}                    ${CYAN}RAPPORT FINAL${NC}                        ${BLUE}│${NC}"
    echo -e "${BLUE}╰─────────────────────────────────────────────────────────────╯${NC}"
    echo ""
    
    local success_rate=$((CHECKS_PASSED * 100 / TOTAL_CHECKS))
    
    echo -e "📊 Résultats:"
    echo -e "   • Total des vérifications: $TOTAL_CHECKS"
    echo -e "   • ${GREEN}Réussies: $CHECKS_PASSED${NC}"
    echo -e "   • ${RED}Échouées: $CHECKS_FAILED${NC}"
    echo -e "   • Taux de réussite: $success_rate%"
    echo ""
    
    if [[ $CHECKS_FAILED -eq 0 ]]; then
        echo -e "${GREEN}🎉 Installation parfaite! Tous les composants sont installés.${NC}"
        echo ""
        echo -e "${CYAN}🚀 Prochaines étapes recommandées:${NC}"
        echo "   1. Redémarrez votre terminal"
        echo "   2. Testez les commandes: ts, dev-symfony, dev-ts"
        echo "   3. Configurez Bitwarden si nécessaire: bw login"
        echo "   4. Personnalisez selon vos besoins"
    elif [[ $success_rate -ge 80 ]]; then
        echo -e "${YELLOW}⚠️ Installation majoritairement réussie avec quelques éléments manquants.${NC}"
        echo ""
        echo -e "${CYAN}🔧 Actions recommandées:${NC}"
        echo "   • Vérifiez les éléments marqués ❌"
        echo "   • Ré-exécutez: chezmoi apply"
        echo "   • Consultez la documentation"
    else
        echo -e "${RED}❌ Installation incomplète. Plusieurs composants manquent.${NC}"
        echo ""
        echo -e "${CYAN}🆘 Actions de dépannage:${NC}"
        echo "   • Ré-exécutez l'installation complète"
        echo "   • Vérifiez votre connexion internet"
        echo "   • Consultez les logs d'erreur"
        echo "   • Restaurez depuis la sauvegarde si nécessaire"
    fi
    
    echo ""
    echo -e "${CYAN}📚 Ressources utiles:${NC}"
    echo "   • Documentation: README.md"
    echo "   • Dépannage: ./remove-dotfiles.sh --help"
    echo "   • Support: GitHub Issues"
    
    # Code de sortie basé sur le résultat
    if [[ $CHECKS_FAILED -eq 0 ]]; then
        exit 0
    elif [[ $success_rate -ge 80 ]]; then
        exit 1
    else
        exit 2
    fi
}

main "$@"