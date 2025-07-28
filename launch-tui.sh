#!/bin/bash

# Dotfiles TUI Launcher
# Construit et lance l'interface TUI pour la gestion des dotfiles

set -euo pipefail

# Couleurs pour l'affichage
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

echo -e "${BLUE}🚀 Dotfiles TUI - Interface de Gestion${NC}"
echo -e "${BLUE}======================================${NC}"
echo ""

# Vérifier que Go est installé
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ Go n'est pas installé. Veuillez installer Go 1.21 ou plus récent.${NC}"
    exit 1
fi

# Vérifier la version de Go
GO_VERSION=$(go version | grep -oE 'go[0-9]+\.[0-9]+' | sed 's/go//')
REQUIRED_VERSION="1.21"

if ! printf '%s\n%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V -C; then
    echo -e "${YELLOW}⚠️ Version Go détectée: $GO_VERSION (minimum requis: $REQUIRED_VERSION)${NC}"
fi

# Construire l'application si nécessaire
if [[ ! -f "./dotfiles-tui" ]] || [[ "./cmd/dotfiles-tui/main.go" -nt "./dotfiles-tui" ]]; then
    echo -e "${YELLOW}🔨 Construction de l'application...${NC}"
    
    if go build -o dotfiles-tui ./cmd/dotfiles-tui; then
        echo -e "${GREEN}✅ Application construite avec succès!${NC}"
    else
        echo -e "${RED}❌ Erreur lors de la construction${NC}"
        exit 1
    fi
else
    echo -e "${GREEN}✅ Application déjà construite${NC}"
fi

echo ""
echo -e "${BLUE}🎮 Lancement de l'interface TUI...${NC}"
echo -e "${YELLOW}💡 Utilisez Ctrl+C pour quitter à tout moment${NC}"
echo ""

# Lancer l'application
exec ./dotfiles-tui