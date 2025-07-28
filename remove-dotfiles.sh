#!/bin/bash

# Script de désinstallation des dotfiles modernes
# Supprime proprement toute la configuration installée par Chezmoi

set -euo pipefail

# Couleurs pour l'affichage
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
PURPLE='\033[0;35m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Variables globales
DRY_RUN=false
INTERACTIVE=true
BACKUP_DIR="$HOME/.dotfiles-removal-backup-$(date +%Y%m%d_%H%M%S)"
VERBOSE=false

# Fonction d'affichage avec couleurs
print_header() {
    echo -e "${PURPLE}╭─────────────────────────────────────────────────────────────╮${NC}"
    echo -e "${PURPLE}│${NC}  🗑️  ${CYAN}Désinstallation des Dotfiles Modernes${NC}  🗑️   ${PURPLE}│${NC}"
    echo -e "${PURPLE}╰─────────────────────────────────────────────────────────────╯${NC}"
    echo ""
}

print_step() {
    echo -e "${BLUE}▶${NC} $1"
}

print_success() {
    echo -e "${GREEN}✅${NC} $1"
}

print_warning() {
    echo -e "${YELLOW}⚠️${NC} $1"
}

print_error() {
    echo -e "${RED}❌${NC} $1"
}

print_info() {
    echo -e "${CYAN}ℹ️${NC} $1"
}

# Fonction d'aide
show_help() {
    cat << EOF
Usage: $0 [OPTIONS]

Options:
    -h, --help          Afficher cette aide
    -d, --dry-run       Mode simulation (ne supprime rien)
    -y, --yes           Mode non-interactif (accepte tout)
    -v, --verbose       Mode verbeux
    -b, --backup-only   Créer seulement une sauvegarde
    --keep-tools        Garder les outils installés (starship, etc.)
    --keep-shell        Ne pas restaurer le shell précédent
    --keep-configs      Garder les fichiers de configuration
    --nuclear           Suppression complète (attention!)

Exemples:
    $0                  # Désinstallation interactive
    $0 --dry-run        # Voir ce qui serait supprimé
    $0 --yes --verbose  # Désinstallation automatique avec détails
    $0 --backup-only    # Créer seulement une sauvegarde
EOF
}

# Fonction de confirmation
confirm() {
    if [[ "$INTERACTIVE" == false ]]; then
        return 0
    fi
    
    local message="$1"
    echo -e "${YELLOW}❓${NC} $message"
    read -p "Continuer? [y/N] " -n 1 -r
    echo
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_info "Opération annulée par l'utilisateur"
        exit 0
    fi
}

# Fonction de sauvegarde
create_backup() {
    print_step "Création d'une sauvegarde dans $BACKUP_DIR"
    
    if [[ "$DRY_RUN" == true ]]; then
        print_info "[DRY RUN] Sauvegarde qui serait créée: $BACKUP_DIR"
        return 0
    fi
    
    mkdir -p "$BACKUP_DIR"
    
    # Sauvegarder les fichiers de configuration importants
    local files_to_backup=(
        "$HOME/.zshrc"
        "$HOME/.gitconfig"
        "$HOME/.tmux.conf"
        "$HOME/.aliases"
        "$HOME/.env"
        "$HOME/.config/starship.toml"
        "$HOME/.config/nvim"
        "$HOME/.config/tmux"
    )
    
    for file in "${files_to_backup[@]}"; do
        if [[ -e "$file" ]]; then
            if [[ "$VERBOSE" == true ]]; then
                print_info "Sauvegarde: $file"
            fi
            cp -r "$file" "$BACKUP_DIR/" 2>/dev/null || true
        fi
    done
    
    # Sauvegarder la liste des packages installés
    if command -v brew &> /dev/null; then
        brew list > "$BACKUP_DIR/brew_packages.txt" 2>/dev/null || true
    fi
    
    if command -v apt &> /dev/null; then
        apt list --installed > "$BACKUP_DIR/apt_packages.txt" 2>/dev/null || true
    fi
    
    print_success "Sauvegarde créée dans $BACKUP_DIR"
}

# Fonction de suppression des fichiers Chezmoi
remove_chezmoi() {
    print_step "Suppression de Chezmoi et des fichiers gérés"
    
    local chezmoi_dirs=(
        "$HOME/.local/share/chezmoi"
        "$HOME/.config/chezmoi"
        "$HOME/.cache/chezmoi"
    )
    
    for dir in "${chezmoi_dirs[@]}"; do
        if [[ -d "$dir" ]]; then
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $dir"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Suppression: $dir"
                fi
                rm -rf "$dir"
            fi
        fi
    done
    
    # Supprimer le binaire Chezmoi
    local chezmoi_bins=(
        "$HOME/.local/bin/chezmoi"
        "$HOME/bin/chezmoi"
        "/usr/local/bin/chezmoi"
    )
    
    for bin in "${chezmoi_bins[@]}"; do
        if [[ -f "$bin" ]]; then
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $bin"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Suppression: $bin"
                fi
                rm -f "$bin"
            fi
        fi
    done
    
    print_success "Chezmoi supprimé"
}

# Fonction de suppression des fichiers de configuration
remove_config_files() {
    if [[ "${KEEP_CONFIGS:-false}" == true ]]; then
        print_info "Conservation des fichiers de configuration (--keep-configs)"
        return 0
    fi
    
    print_step "Suppression des fichiers de configuration"
    
    local config_files=(
        "$HOME/.zshrc"
        "$HOME/.aliases"
        "$HOME/.env"
        "$HOME/.gitconfig"
        "$HOME/.gitignore_global"
        "$HOME/.tmux.conf"
        "$HOME/.config/starship.toml"
    )
    
    for file in "${config_files[@]}"; do
        if [[ -f "$file" ]]; then
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $file"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Suppression: $file"
                fi
                rm -f "$file"
            fi
        fi
    done
    
    # Supprimer les dossiers de configuration
    local config_dirs=(
        "$HOME/.config/nvim"
        "$HOME/.config/tmux"
    )
    
    for dir in "${config_dirs[@]}"; do
        if [[ -d "$dir" ]]; then
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $dir"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Suppression: $dir"
                fi
                rm -rf "$dir"
            fi
        fi
    done
    
    print_success "Fichiers de configuration supprimés"
}

# Fonction de suppression d'Oh My Zsh
remove_oh_my_zsh() {
    print_step "Suppression d'Oh My Zsh et des plugins"
    
    if [[ -d "$HOME/.oh-my-zsh" ]]; then
        if [[ "$DRY_RUN" == true ]]; then
            print_info "[DRY RUN] Supprimerait: $HOME/.oh-my-zsh"
        else
            if [[ "$VERBOSE" == true ]]; then
                print_info "Suppression: $HOME/.oh-my-zsh"
            fi
            rm -rf "$HOME/.oh-my-zsh"
        fi
        print_success "Oh My Zsh supprimé"
    else
        print_info "Oh My Zsh non trouvé"
    fi
}

# Fonction de suppression des outils installés
remove_installed_tools() {
    if [[ "${KEEP_TOOLS:-false}" == true ]]; then
        print_info "Conservation des outils installés (--keep-tools)"
        return 0
    fi
    
    print_step "Suppression des outils installés"
    
    # Starship
    if command -v starship &> /dev/null; then
        local starship_path=$(which starship)
        if [[ "$DRY_RUN" == true ]]; then
            print_info "[DRY RUN] Supprimerait Starship: $starship_path"
        else
            confirm "Supprimer Starship prompt?"
            if [[ "$VERBOSE" == true ]]; then
                print_info "Suppression: $starship_path"
            fi
            rm -f "$starship_path"
            print_success "Starship supprimé"
        fi
    fi
    
    # tmux plugins (TPM)
    if [[ -d "$HOME/.tmux/plugins" ]]; then
        if [[ "$DRY_RUN" == true ]]; then
            print_info "[DRY RUN] Supprimerait: $HOME/.tmux/plugins"
        else
            if [[ "$VERBOSE" == true ]]; then
                print_info "Suppression: $HOME/.tmux/plugins"
            fi
            rm -rf "$HOME/.tmux/plugins"
        fi
    fi
    
    print_success "Outils supprimés"
}

# Fonction de restauration du shell
restore_shell() {
    if [[ "${KEEP_SHELL:-false}" == true ]]; then
        print_info "Conservation du shell actuel (--keep-shell)"
        return 0
    fi
    
    print_step "Restauration du shell par défaut"
    
    # Détecter le shell par défaut du système
    local default_shell="/bin/bash"
    if [[ -f "/bin/bash" ]]; then
        default_shell="/bin/bash"
    elif [[ -f "/usr/bin/bash" ]]; then
        default_shell="/usr/bin/bash"
    fi
    
    if [[ "$SHELL" == *"zsh"* ]]; then
        if [[ "$DRY_RUN" == true ]]; then
            print_info "[DRY RUN] Restaurerait le shell vers: $default_shell"
        else
            confirm "Restaurer le shell par défaut ($default_shell)?"
            chsh -s "$default_shell"
            print_success "Shell restauré vers $default_shell"
        fi
    else
        print_info "Shell déjà configuré sur $SHELL"
    fi
}

# Fonction de nettoyage des caches
clean_caches() {
    print_step "Nettoyage des caches"
    
    local cache_dirs=(
        "$HOME/.cache/zsh"
        "$HOME/.cache/nvim"
        "$HOME/.cache/tmux"
        "$HOME/.cache/starship"
        "$HOME/.zcompdump*"
    )
    
    for cache in "${cache_dirs[@]}"; do
        if [[ -e "$cache" ]]; then
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $cache"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Suppression: $cache"
                fi
                rm -rf "$cache"
            fi
        fi
    done
    
    print_success "Caches nettoyés"
}

# Fonction de suppression nucléaire (tout supprimer)
nuclear_removal() {
    print_warning "MODE NUCLÉAIRE ACTIVÉ - SUPPRESSION COMPLÈTE"
    confirm "ATTENTION: Ceci va supprimer TOUS les outils et configurations. Êtes-vous sûr?"
    
    # Supprimer tous les outils modernes installés
    local tools_to_remove=(
        "starship"
        "eza"
        "bat"
        "fd"
        "rg"
        "dust"
        "duf"
        "procs"
        "lazygit"
        "zoxide"
    )
    
    for tool in "${tools_to_remove[@]}"; do
        if command -v "$tool" &> /dev/null; then
            local tool_path=$(which "$tool")
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $tool_path"
            else
                print_warning "Suppression de $tool"
                rm -f "$tool_path"
            fi
        fi
    done
    
    # Supprimer les dossiers de développement créés
    local dev_dirs=(
        "$HOME/projects"
        "$HOME/.local/bin"
    )
    
    for dir in "${dev_dirs[@]}"; do
        if [[ -d "$dir" ]]; then
            confirm "Supprimer le dossier $dir?"
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Supprimerait: $dir"
            else
                rm -rf "$dir"
            fi
        fi
    done
}

# Fonction de restauration depuis la sauvegarde
restore_from_backup() {
    if [[ ! -d "$BACKUP_DIR" ]]; then
        print_error "Aucune sauvegarde trouvée dans $BACKUP_DIR"
        return 1
    fi
    
    print_step "Restauration depuis la sauvegarde"
    
    # Restaurer les fichiers de configuration
    for file in "$BACKUP_DIR"/*; do
        if [[ -f "$file" ]]; then
            local basename=$(basename "$file")
            local target="$HOME/$basename"
            
            if [[ "$DRY_RUN" == true ]]; then
                print_info "[DRY RUN] Restaurerait: $target"
            else
                if [[ "$VERBOSE" == true ]]; then
                    print_info "Restauration: $target"
                fi
                cp "$file" "$target"
            fi
        fi
    done
    
    print_success "Restauration terminée"
}

# Fonction de rapport final
show_final_report() {
    echo ""
    print_header
    echo -e "${GREEN}🎉 Désinstallation terminée avec succès!${NC}"
    echo ""
    
    if [[ "$DRY_RUN" == true ]]; then
        print_info "Mode simulation - Aucune modification effectuée"
    else
        print_info "Sauvegarde disponible dans: $BACKUP_DIR"
        print_info "Pour restaurer: cp -r $BACKUP_DIR/* ~/"
    fi
    
    echo ""
    print_info "Actions recommandées:"
    echo "  • Redémarrez votre terminal"
    echo "  • Vérifiez votre shell: echo \$SHELL"
    echo "  • Supprimez la sauvegarde si tout fonctionne: rm -rf $BACKUP_DIR"
    echo ""
    
    if [[ "${KEEP_SHELL:-false}" == false ]]; then
        print_warning "N'oubliez pas de redémarrer votre terminal pour activer le nouveau shell"
    fi
}

# Fonction principale
main() {
    # Parse des arguments
    while [[ $# -gt 0 ]]; do
        case $1 in
            -h|--help)
                show_help
                exit 0
                ;;
            -d|--dry-run)
                DRY_RUN=true
                shift
                ;;
            -y|--yes)
                INTERACTIVE=false
                shift
                ;;
            -v|--verbose)
                VERBOSE=true
                shift
                ;;
            -b|--backup-only)
                create_backup
                exit 0
                ;;
            --keep-tools)
                KEEP_TOOLS=true
                shift
                ;;
            --keep-shell)
                KEEP_SHELL=true
                shift
                ;;
            --keep-configs)
                KEEP_CONFIGS=true
                shift
                ;;
            --nuclear)
                NUCLEAR=true
                shift
                ;;
            *)
                print_error "Option inconnue: $1"
                show_help
                exit 1
                ;;
        esac
    done
    
    # Affichage de l'en-tête
    print_header
    
    if [[ "$DRY_RUN" == true ]]; then
        print_warning "MODE SIMULATION ACTIVÉ - Aucune modification ne sera effectuée"
        echo ""
    fi
    
    # Vérifications préliminaires
    if [[ ! -d "$HOME/.local/share/chezmoi" ]] && [[ ! -f "$HOME/.zshrc" ]]; then
        print_warning "Aucune configuration dotfiles détectée"
        confirm "Continuer quand même?"
    fi
    
    # Confirmation finale
    if [[ "${NUCLEAR:-false}" == true ]]; then
        print_warning "MODE NUCLÉAIRE - Suppression complète demandée"
    fi
    
    confirm "Commencer la désinstallation des dotfiles?"
    
    # Exécution des étapes
    create_backup
    
    if [[ "${NUCLEAR:-false}" == true ]]; then
        nuclear_removal
    fi
    
    remove_chezmoi
    remove_config_files
    remove_oh_my_zsh
    remove_installed_tools
    clean_caches
    restore_shell
    
    show_final_report
}

# Gestion des signaux
trap 'print_error "Interruption détectée. Nettoyage..."; exit 130' INT TERM

# Exécution
main "$@"