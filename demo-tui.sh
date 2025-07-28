#!/bin/bash

# Clean demo script for the Dotfiles TUI
# Simple and elegant presentation

set -euo pipefail

# Simple colors
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

clear

echo -e "${PURPLE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║${NC}  🚀  ${CYAN}Dotfiles TUI - Interface Moderne de Gestion${NC}  🚀   ${PURPLE}║${NC}"
echo -e "${PURPLE}╚══════════════════════════════════════════════════════════════╝${NC}"
echo ""

echo -e "${BLUE}🎉 Bienvenue dans votre interface TUI moderne!${NC}"
echo ""

echo -e "${CYAN}✨ Fonctionnalités principales:${NC}"
echo ""
echo -e "  ${GREEN}🚀 Installation Interactive${NC}     - Guide étape par étape"
echo -e "  ${GREEN}✅ Vérification du Système${NC}      - Contrôle de santé complet"
echo -e "  ${GREEN}⚙️ Gestion de Configuration${NC}     - Édition des dotfiles"
echo -e "  ${GREEN}💾 Sauvegarde & Restauration${NC}    - Gestion des backups"
echo -e "  ${GREEN}🔧 Gestion des Outils${NC}           - Installation/mise à jour"
echo -e "  ${GREEN}🔐 Configuration des Secrets${NC}    - Intégration Bitwarden"
echo -e "  ${GREEN}📊 Informations Système${NC}         - Vue d'ensemble"
echo ""

echo -e "${YELLOW}🎨 Interface moderne avec:${NC}"
echo ""
echo -e "  • Thème ${PURPLE}Catppuccin${NC} coordonné"
echo -e "  • Navigation intuitive au clavier"
echo -e "  • Indicateurs de progression en temps réel"
echo -e "  • Messages d'état colorés et clairs"
echo ""

echo -e "${CYAN}⌨️ Navigation:${NC}"
echo ""
echo -e "  ${GREEN}↑/↓${NC}      - Naviguer dans les menus"
echo -e "  ${GREEN}Entrée${NC}   - Sélectionner une option"
echo -e "  ${GREEN}Échap${NC}    - Retour au menu précédent"
echo -e "  ${GREEN}Ctrl+C${NC}   - Quitter l'application"
echo ""

echo -e "${BLUE}🚀 Prêt à commencer?${NC}"
echo ""

# Simple choice
read -p "Appuyez sur Entrée pour lancer l'interface TUI..."

echo ""
echo -e "${GREEN}🚀 Lancement de l'interface...${NC}"
echo ""

# Launch the TUI
exec ./dotfiles-tui