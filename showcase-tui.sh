#!/bin/bash

# Showcase script for the enhanced Dotfiles TUI with Lip Gloss styling
# Demonstrates the beautiful new interface

set -euo pipefail

# Couleurs Catppuccin
MAUVE='\033[38;2;203;166;247m'
BLUE='\033[38;2;137;180;250m'
PINK='\033[38;2;243;139;168m'
GREEN='\033[38;2;166;227;161m'
YELLOW='\033[38;2;249;226;175m'
SKY='\033[38;2;116;199;236m'
TEXT='\033[38;2;205;214;244m'
SUBTEXT='\033[38;2;186;194;222m'
NC='\033[0m'

clear

echo -e "${MAUVE}╭─────────────────────────────────────────────────────────────╮${NC}"
echo -e "${MAUVE}│${NC}  ✨  ${SKY}Dotfiles TUI - Interface Magnifiée avec Lip Gloss${NC}  ✨   ${MAUVE}│${NC}"
echo -e "${MAUVE}╰─────────────────────────────────────────────────────────────╯${NC}"
echo ""

echo -e "${BLUE}🎨 Votre TUI a été magnifiée avec Lip Gloss de Charm!${NC}"
echo ""

echo -e "${SKY}✨ Nouvelles fonctionnalités visuelles:${NC}"
echo ""
echo -e "  ${GREEN}🎨 Thème Catppuccin Complet${NC}      - Couleurs coordonnées et harmonieuses"
echo -e "  ${GREEN}🖼️ Bordures Élégantes${NC}            - Bordures arrondies et doubles"
echo -e "  ${GREEN}📦 Cartes Stylisées${NC}              - Contenu organisé en cartes"
echo -e "  ${GREEN}🌈 Badges de Status${NC}              - Indicateurs colorés et expressifs"
echo -e "  ${GREEN}📊 Barres de Progression${NC}         - Indicateurs visuels améliorés"
echo -e "  ${GREEN}📋 Logs Formatés${NC}                 - Affichage des logs avec timestamps"
echo -e "  ${GREEN}🎯 Séparateurs Décoratifs${NC}        - Organisation visuelle claire"
echo -e "  ${GREEN}💫 Animations Fluides${NC}            - Spinners et transitions"
echo ""

echo -e "${YELLOW}🎨 Palette de couleurs Catppuccin Mocha:${NC}"
echo ""
echo -e "  ${MAUVE}● Mauve${NC}     - Couleur principale et accents"
echo -e "  ${BLUE}● Bleu${NC}      - Éléments secondaires et bordures"
echo -e "  ${PINK}● Rose${NC}      - Accents et éléments actifs"
echo -e "  ${GREEN}● Vert${NC}      - Succès et éléments positifs"
echo -e "  ${YELLOW}● Jaune${NC}     - Avertissements et informations"
echo -e "  ${SKY}● Ciel${NC}      - Informations et éléments neutres"
echo ""

echo -e "${PINK}🎭 Éléments de style ajoutés:${NC}"
echo ""
echo -e "  ${TEXT}📦 ${SUBTEXT}CreateCard()${NC}           - Cartes avec bordures et padding"
echo -e "  ${TEXT}🏷️ ${SUBTEXT}CreateStatusBadge()${NC}    - Badges colorés selon le status"
echo -e "  ${TEXT}📊 ${SUBTEXT}CreateProgressBar()${NC}    - Barres de progression stylisées"
echo -e "  ${TEXT}🎨 ${SUBTEXT}CreateBanner()${NC}         - Bannières décoratives"
echo -e "  ${TEXT}📝 ${SUBTEXT}CreateLogEntry()${NC}       - Entrées de log formatées"
echo -e "  ${TEXT}🔘 ${SUBTEXT}CreateButton()${NC}         - Boutons interactifs"
echo -e "  ${TEXT}📋 ${SUBTEXT}CreateTable()${NC}          - Tableaux stylisés"
echo ""

echo -e "${SKY}🖥️ Améliorations par écran:${NC}"
echo ""
echo -e "  ${GREEN}🏠 Menu Principal${NC}        - Bannière décorative + cartes"
echo -e "  ${GREEN}🚀 Installation${NC}          - Progression visuelle + logs formatés"
echo -e "  ${GREEN}✅ Vérification${NC}          - Sections organisées + badges status"
echo -e "  ${GREEN}⚙️ Configuration${NC}         - Interface épurée + navigation claire"
echo -e "  ${GREEN}💾 Sauvegarde${NC}            - Présentation structurée"
echo -e "  ${GREEN}🔧 Outils${NC}                - Gestion visuelle améliorée"
echo -e "  ${GREEN}🔐 Secrets${NC}               - Interface sécurisée et claire"
echo -e "  ${GREEN}📊 Informations${NC}          - Données organisées en tableaux"
echo ""

echo -e "${MAUVE}🎮 Prêt à découvrir la nouvelle interface?${NC}"
echo ""

# Animation de lancement
echo -e "${YELLOW}🚀 Lancement de l'interface magnifiée...${NC}"
for i in {1..3}; do
    echo -n "."
    sleep 0.5
done
echo ""
echo ""

echo -e "${GREEN}✨ Profitez de votre interface moderne et élégante!${NC}"
echo ""

# Lancer l'application
exec ./dotfiles-tui