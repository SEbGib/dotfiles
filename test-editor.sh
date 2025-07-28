#!/bin/bash

# Test script for TUI editor functionality

echo "🧪 Test de la fonctionnalité d'édition du TUI"
echo ""

# Check if common editors are available
echo "🔍 Vérification des éditeurs disponibles:"
for editor in nvim vim nano code subl; do
    if command -v "$editor" &> /dev/null; then
        echo "  ✅ $editor - disponible"
    else
        echo "  ❌ $editor - non disponible"
    fi
done

echo ""
echo "📝 Test des fichiers de configuration:"

# Test file paths
test_files=(
    "$HOME/.zshrc"
    "$HOME/.gitconfig"
    "$HOME/.config/starship.toml"
    "$HOME/.config/nvim/init.lua"
    "$HOME/.aliases"
)

for file in "${test_files[@]}"; do
    if [[ -f "$file" ]]; then
        echo "  ✅ $(basename "$file") - existe"
    else
        echo "  ⚠️ $(basename "$file") - sera créé si nécessaire"
    fi
done

echo ""
echo "🚀 Lancement du TUI pour test..."
echo "   → Allez dans 'Gestion de Configuration'"
echo "   → Sélectionnez un fichier à éditer"
echo "   → Appuyez sur 'E' ou 'Entrée' pour ouvrir l'éditeur"
echo ""

./dotfiles-tui