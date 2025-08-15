# 🎯 Raccourcis clavier - Configuration Neovim

Documentation complète des raccourcis clavier configurés dans Neovim avec LazyVim et plugins personnalisés.

## 📁 Explorateur de fichiers

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>e` | Explorateur natif (netrw) | Normal |
| `<leader>E` | Explorateur dossier du fichier | Normal |

> **Note** : snacks.nvim picker est désactivé pour éviter les conflits multiples

## 📐 Navigation

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<C-h>` | Aller à la fenêtre de gauche | Normal |
| `<C-j>` | Aller à la fenêtre du bas | Normal |
| `<C-k>` | Aller à la fenêtre du haut | Normal |
| `<C-l>` | Aller à la fenêtre de droite | Normal |
| `<C-Up>` | Augmenter la hauteur | Normal |
| `<C-Down>` | Diminuer la hauteur | Normal |
| `<C-Left>` | Diminuer la largeur | Normal |
| `<C-Right>` | Augmenter la largeur | Normal |
| `j/k` | Navigation intelligente avec wraps | Normal |

## ✂️ Édition

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<` | Indenter à gauche (garder sélection) | Visuel |
| `>` | Indenter à droite (garder sélection) | Visuel |
| `<A-j>` | Déplacer ligne vers le bas | Normal/Insertion |
| `<A-k>` | Déplacer ligne vers le haut | Normal/Insertion |
| `<A-j>` | Déplacer sélection vers le bas | Visuel |
| `<A-k>` | Déplacer sélection vers le haut | Visuel |
| `<C-s>` | Sauvegarder fichier | Tous modes |
| `p` | Coller sans yank | Visuel |
| `<leader>d` | Dupliquer la ligne | Normal |

## 📋 Buffers et Onglets

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<S-h>` | Buffer précédent | Normal |
| `<S-l>` | Buffer suivant | Normal |
| `<leader>bb` | Basculer vers l'autre buffer | Normal |
| `<leader>tn` | Nouvel onglet | Normal |
| `<leader>tc` | Fermer l'onglet | Normal |
| `<leader>to` | Fermer les autres onglets | Normal |

## 🪟 Splits

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>-` | Split horizontal | Normal |
| `<leader>|` | Split vertical | Normal |

## 💾 Fichiers

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>w` | Sauvegarder | Normal |
| `<leader>W` | Sauvegarder tout | Normal |
| `<leader>q` | Quitter | Normal |
| `<leader>Q` | Quitter tout | Normal |

## 🖥️ Terminal

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<Esc><Esc>` | Sortir du mode terminal | Terminal |
| `<C-h/j/k/l>` | Navigation entre fenêtres | Terminal |

## 🔍 Recherche

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<Esc>` | Désactiver surbrillance | Normal |
| `<leader>sr` | Remplacer le mot sous le curseur | Normal |

## 🧰 Développement

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>cd` | Diagnostic de ligne | Normal |
| `[d` | Diagnostic précédent | Normal |
| `]d` | Diagnostic suivant | Normal |
| `gd` | Aller à la définition | Normal |
| `gr` | Références | Normal |
| `gI` | Aller à l'implémentation | Normal |
| `gy` | Définition de type | Normal |
| `gD` | Aller à la déclaration | Normal |
| `K` | Documentation | Normal |
| `gK` | Aide signature | Normal |

## 🚀 PHP/Symfony

| Raccourci | Description | Mode | Source |
|-----------|-------------|------|--------|
| `<leader>ps` | Console Symfony | Normal | keymaps.lua |
| `<leader>pt` | PHPUnit (tmux) / Templates Twig (telescope) | Normal | tmux.lua / telescope.lua |
| `<leader>pcc` | Clear Symfony Cache (tmux) | Normal | tmux.lua |
| `<leader>pcs` | Symfony Console (tmux) | Normal | tmux.lua |
| `<leader>pss` | Symfony Serve (tmux) | Normal | tmux.lua |

## 🟨 TypeScript/Node.js

| Raccourci | Description | Mode | Source |
|-----------|-------------|------|--------|
| `<leader>nr` | Scripts npm | Normal | keymaps.lua |
| `<leader>nt` | Tests npm (local) / Tests npm (tmux) | Normal | keymaps.lua / tmux.lua |
| `<leader>nd` | npm run dev (tmux) | Normal | tmux.lua |
| `<leader>nb` | npm run build (tmux) | Normal | tmux.lua |

## 🔧 Tmux Integration

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<M-h/j/k/l>` | Navigation tmux panes | Normal |
| `<M-\>` | Pane précédent (tmux) | Normal |
| `<leader>vp` | Tmux: Prompt Command | Normal |
| `<leader>vl` | Tmux: Run Last Command | Normal |
| `<leader>vi` | Tmux: Inspect Runner | Normal |
| `<leader>vq` | Tmux: Close Runner | Normal |
| `<leader>vx` | Tmux: Interrupt Runner | Normal |
| `<leader>vz` | Tmux: Zoom Runner | Normal |

## 🔭 Telescope

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>ff` | Trouver fichiers | Normal |
| `<leader>fr` | Fichiers récents | Normal |
| `<leader>fg` | Recherche dans fichiers | Normal |
| `<leader>fw` | Rechercher mot | Normal |
| `<leader>fb` | Buffers | Normal |
| `<leader>gc` | Git commits | Normal |
| `<leader>gs` | Git status | Normal |
| `<leader>gb` | Git branches | Normal |
| `<leader>fs` | Symboles document | Normal |
| `<leader>fS` | Symboles workspace | Normal |
| `<leader>fd` | Diagnostics | Normal |
| `<leader>fh` | Aide | Normal |
| `<leader>fk` | Keymaps | Normal |
| `<leader>fc` | Commandes | Normal |
| `<leader>pt` | Templates Twig | Normal |
| `<leader>pe` | Entités | Normal |

## 🛠️ Utilitaires

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>l` | Lazy | Normal |
| `<leader>ui` | Inspecter la position | Normal |
| `<leader>ut` | Update Tmux statusline | Normal |

## 🔧 Quickfix

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>xl` | Liste de localisation | Normal |
| `<leader>xq` | Quickfix | Normal |
| `[q` | Quickfix précédent | Normal |
| `]q` | Quickfix suivant | Normal |

## 🌐 Git

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<leader>gg` | Git status | Normal |

## 🔤 Divers

| Raccourci | Description | Mode |
|-----------|-------------|------|
| `<C-a>` | Sélectionner tout | Normal |
| `<leader>o` | Nouvelle ligne en dessous | Normal |
| `<leader>O` | Nouvelle ligne au dessus | Normal |

## 📝 Notes importantes

### Conflits résolus
- **`<leader>pt`** : Contexte intelligent entre PHPUnit (tmux), Templates Twig (telescope)
- **Mouvements de lignes** : Doublons supprimés, version unique conservée
- **Points d'annulation** : Doublons supprimés
- **Explorateur** : neo-tree.lua supprimé, utilise netrw natif

### Fichiers sources
- **keymaps.lua** : Raccourcis de base et développement
- **tmux.lua** : Intégration tmux et commandes contextuelles
- **telescope.lua** : Recherche floue et navigation
- **snacks-disable.lua** : Désactivation du picker automatique

### Conventions
- `<leader>` = espace (par défaut LazyVim)
- `<M-*>` = Alt + touche
- `<C-*>` = Ctrl + touche
- `<S-*>` = Shift + touche