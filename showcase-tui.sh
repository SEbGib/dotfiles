#!/bin/bash

# Simple showcase for the Dotfiles TUI
# Clean and focused presentation

set -euo pipefail

# Colors
BLUE='\033[0;34m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
PURPLE='\033[0;35m'
NC='\033[0m'

clear

echo -e "${PURPLE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${PURPLE}║${NC}     ${CYAN}Dotfiles TUI - Interface Élégante et Moderne${NC}     ${PURPLE}║${NC}"
echo -e "${PURPLE}╚══════════════════════════════════════════════════════════════╝${NC}"
echo ""

echo -e "${BLUE}🎨 Votre TUI a été améliorée avec un design moderne!${NC}"
echo ""

echo -e "${CYAN}✨ Améliorations visuelles:${NC}"
echo ""
echo -e "  ${GREEN}🎨 Thème Catppuccin${NC}             - Couleurs harmonieuses"
echo -e "  ${GREEN}📦 Interface Organisée${NC}          - Contenu structuré"
echo -e "  ${GREEN}🌈 Indicateurs Colorés${NC}          - Status visuels clairs"
echo -e "  ${GREEN}📊 Barres de Progression${NC}        - Suivi en temps réel"
echo -e "  ${GREEN}💫 Navigation Fluide${NC}            - Expérience utilisateur optimisée"
echo ""

echo -e "${YELLOW}🎯 Fonctionnalités par écran:${NC}"
echo ""
echo -e "  ${GREEN}🏠 Menu Principal${NC}        - Interface épurée et claire"
echo -e "  ${GREEN}🚀 Installation${NC}          - Progression visuelle détaillée"
echo -e "  ${GREEN}✅ Vérification${NC}          - Rapport de santé complet"
echo -e "  ${GREEN}⚙️ Configuration${NC}         - Gestion simplifiée"
echo ""

echo -e "${CYAN}🎮 Commandes de lancement:${NC}"
echo ""
echo -e "  ${YELLOW}./demo-tui.sh${NC}           - Demo simple et rapide"
echo -e "  ${YELLOW}./launch-tui.sh${NC}         - Lancement avec vérifications"
echo -e "  ${YELLOW}./dotfiles-tui${NC}          - Lancement direct"
echo -e "  ${YELLOW}make run${NC}                - Construction et lancement"
echo ""

echo -e "${BLUE}🚀 Découvrez votre interface moderne maintenant!${NC}"
echo ""

# Simple prompt
read -p "Appuyez sur Entrée pour lancer l'interface TUI..."

echo ""
echo -e "${GREEN}✨ Lancement de votre interface moderne...${NC}"
echo ""

# Launch
exec ./dotfiles-tui