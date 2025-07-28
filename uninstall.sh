#!/bin/bash

# Script de désinstallation rapide des dotfiles
# Version simplifiée pour une suppression rapide

set -euo pipefail

echo "🗑️ Désinstallation rapide des dotfiles modernes"
echo "=============================================="
echo ""

# Confirmation
read -p "⚠️  Voulez-vous vraiment supprimer la configuration dotfiles? [y/N] " -n 1 -r
echo
if [[ ! $REPLY =~ ^[Yy]$ ]]; then
    echo "❌ Annulé par l'utilisateur"
    exit 0
fi

echo "📦 Création d'une sauvegarde..."
BACKUP_DIR="$HOME/.dotfiles-backup-$(date +%Y%m%d_%H%M%S)"
mkdir -p "$BACKUP_DIR"

# Sauvegarder les fichiers importants
cp "$HOME/.zshrc" "$BACKUP_DIR/" 2>/dev/null || true
cp "$HOME/.gitconfig" "$BACKUP_DIR/" 2>/dev/null || true
cp -r "$HOME/.config/nvim" "$BACKUP_DIR/" 2>/dev/null || true

echo "✅ Sauvegarde créée dans: $BACKUP_DIR"

echo "🧹 Suppression des fichiers..."

# Supprimer Chezmoi
rm -rf "$HOME/.local/share/chezmoi" 2>/dev/null || true
rm -f "$HOME/.local/bin/chezmoi" 2>/dev/null || true

# Supprimer les configurations
rm -f "$HOME/.zshrc" 2>/dev/null || true
rm -f "$HOME/.aliases" 2>/dev/null || true
rm -f "$HOME/.env" 2>/dev/null || true
rm -f "$HOME/.gitconfig" 2>/dev/null || true
rm -f "$HOME/.config/starship.toml" 2>/dev/null || true
rm -rf "$HOME/.config/nvim" 2>/dev/null || true
rm -rf "$HOME/.config/tmux" 2>/dev/null || true

# Supprimer Oh My Zsh
rm -rf "$HOME/.oh-my-zsh" 2>/dev/null || true

# Nettoyer les caches
rm -rf "$HOME/.cache/zsh" 2>/dev/null || true
rm -rf "$HOME/.cache/nvim" 2>/dev/null || true
rm -f "$HOME/.zcompdump"* 2>/dev/null || true

echo "✅ Suppression terminée"

# Restaurer le shell par défaut
if [[ "$SHELL" == *"zsh"* ]]; then
    echo "🐚 Restauration du shell par défaut..."
    chsh -s /bin/bash 2>/dev/null || true
    echo "✅ Shell restauré vers bash"
fi

echo ""
echo "🎉 Désinstallation terminée avec succès!"
echo ""
echo "📋 Actions recommandées:"
echo "   • Redémarrez votre terminal"
echo "   • Sauvegarde disponible: $BACKUP_DIR"
echo "   • Pour restaurer: cp -r $BACKUP_DIR/* ~/"
echo ""
echo "💡 Pour une désinstallation avancée, utilisez: ./remove-dotfiles.sh --help"