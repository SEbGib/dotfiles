#!/bin/bash

# Demo script for the Dotfiles TUI
# Shows the capabilities of the new Bubble Tea interface

set -euo pipefail

# Couleurs
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

clear

echo -e "${PURPLE}╭─────────────────────────────────────────────────────────────╮${NC}"
echo -e "${PURPLE}│${NC}  🚀  ${CYAN}Dotfiles TUI - Interface Moderne de Gestion${NC}  🚀   ${PURPLE}│${NC}"
echo -e "${PURPLE}╰─────────────────────────────────────────────────────────────╯${NC}"
echo ""

echo -e "${BLUE}🎉 Félicitations! Votre nouvelle interface TUI est prête!${NC}"
echo ""

echo -e "${CYAN}✨ Fonctionnalités disponibles:${NC}"
echo ""
echo -e "  ${GREEN}🚀 Installation Interactive${NC}     - Guide étape par étape"
echo -e "  ${GREEN}✅ Vérification du Système${NC}      - Contrôle de santé complet"
echo -e "  ${GREEN}⚙️ Gestion de Configuration${NC}     - Édition des dotfiles"
echo -e "  ${GREEN}💾 Sauvegarde & Restauration${NC}    - Gestion des backups"
echo -e "  ${GREEN}🔧 Gestion des Outils${NC}           - Installation/mise à jour"
echo -e "  ${GREEN}🔐 Configuration des Secrets${NC}    - Intégration Bitwarden"
echo -e "  ${GREEN}📊 Informations Système${NC}         - Vue d'ensemble"
echo ""

echo -e "${YELLOW}🎮 Comment utiliser:${NC}"
echo ""
echo -e "  ${CYAN}1. Lancement rapide:${NC}"
echo -e "     ${GREEN}./launch-tui.sh${NC}     # Script de lancement automatique"
echo -e "     ${GREEN}make run${NC}            # Ou via Makefile"
echo ""
echo -e "  ${CYAN}2. Construction manuelle:${NC}"
echo -e "     ${GREEN}make build${NC}          # Construire l'application"
echo -e "     ${GREEN}./dotfiles-tui${NC}      # Lancer directement"
echo ""
echo -e "  ${CYAN}3. Installation système:${NC}"
echo -e "     ${GREEN}make install${NC}        # Installer dans /usr/local/bin"
echo -e "     ${GREEN}dotfiles-tui${NC}        # Lancer depuis n'importe où"
echo ""

echo -e "${YELLOW}⌨️ Navigation:${NC}"
echo ""
echo -e "  ${GREEN}↑/↓${NC}      - Naviguer dans les menus"
echo -e "  ${GREEN}Entrée${NC}   - Sélectionner une option"
echo -e "  ${GREEN}Échap${NC}    - Retour au menu précédent"
echo -e "  ${GREEN}Ctrl+C${NC}   - Quitter l'application"
echo ""

echo -e "${CYAN}🎨 Interface moderne avec:${NC}"
echo ""
echo -e "  • Thème ${PURPLE}Catppuccin${NC} coordonné"
echo -e "  • Indicateurs de progression en temps réel"
echo -e "  • Logs interactifs pendant l'installation"
echo -e "  • Vérifications système détaillées"
echo -e "  • Navigation intuitive au clavier"
echo ""

echo -e "${BLUE}🚀 Prêt à commencer?${NC}"
echo ""
echo -e "${YELLOW}Choisissez votre méthode de lancement:${NC}"
echo ""

PS3="Votre choix: "
options=("🎮 Lancer l'interface TUI maintenant" "📚 Voir la documentation TUI" "🔧 Construire et installer" "❌ Quitter")

select opt in "${options[@]}"; do
    case $opt in
        "🎮 Lancer l'interface TUI maintenant")
            echo ""
            echo -e "${GREEN}🚀 Lancement de l'interface TUI...${NC}"
            echo ""
            exec ./launch-tui.sh
            break
            ;;
        "📚 Voir la documentation TUI")
            echo ""
            echo -e "${CYAN}📖 Ouverture de la documentation...${NC}"
            if command -v bat &> /dev/null; then
                bat TUI_README.md
            elif command -v less &> /dev/null; then
                less TUI_README.md
            else
                cat TUI_README.md
            fi
            break
            ;;
        "🔧 Construire et installer")
            echo ""
            echo -e "${YELLOW}🔨 Construction et installation...${NC}"
            make build
            echo ""
            echo -e "${GREEN}✅ Application construite!${NC}"
            echo -e "${CYAN}💡 Lancez avec: ./dotfiles-tui${NC}"
            echo -e "${CYAN}💡 Ou installez avec: make install${NC}"
            break
            ;;
        "❌ Quitter")
            echo ""
            echo -e "${BLUE}👋 À bientôt! Votre TUI vous attend.${NC}"
            break
            ;;
        *) 
            echo -e "${YELLOW}Option invalide. Veuillez choisir 1-4.${NC}"
            ;;
    esac
done

echo ""
echo -e "${PURPLE}╭─────────────────────────────────────────────────────────────╮${NC}"
echo -e "${PURPLE}│${NC}     ${CYAN}Profitez de votre nouvelle interface moderne!${NC}     ${PURPLE}│${NC}"
echo -e "${PURPLE}╰─────────────────────────────────────────────────────────────╯${NC}"