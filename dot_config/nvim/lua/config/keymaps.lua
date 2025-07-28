-- Keymaps Neovim modernes
-- Raccourcis optimisés pour le développement

local map = vim.keymap.set

-- ===== NAVIGATION =====
-- Meilleure navigation dans les wraps
map("n", "k", "v:count == 0 ? 'gk' : 'k'", { expr = true, silent = true })
map("n", "j", "v:count == 0 ? 'gj' : 'j'", { expr = true, silent = true })

-- Navigation entre les fenêtres
map("n", "<C-h>", "<C-w>h", { desc = "Aller à la fenêtre de gauche" })
map("n", "<C-j>", "<C-w>j", { desc = "Aller à la fenêtre du bas" })
map("n", "<C-k>", "<C-w>k", { desc = "Aller à la fenêtre du haut" })
map("n", "<C-l>", "<C-w>l", { desc = "Aller à la fenêtre de droite" })

-- Redimensionnement des fenêtres
map("n", "<C-Up>", "<cmd>resize +2<cr>", { desc = "Augmenter la hauteur" })
map("n", "<C-Down>", "<cmd>resize -2<cr>", { desc = "Diminuer la hauteur" })
map("n", "<C-Left>", "<cmd>vertical resize -2<cr>", { desc = "Diminuer la largeur" })
map("n", "<C-Right>", "<cmd>vertical resize +2<cr>", { desc = "Augmenter la largeur" })

-- ===== ÉDITION =====
-- Meilleur indentation
map("v", "<", "<gv")
map("v", ">", ">gv")

-- Déplacer les lignes
map("n", "<A-j>", "<cmd>m .+1<cr>==", { desc = "Déplacer la ligne vers le bas" })
map("n", "<A-k>", "<cmd>m .-2<cr>==", { desc = "Déplacer la ligne vers le haut" })
map("i", "<A-j>", "<esc><cmd>m .+1<cr>==gi", { desc = "Déplacer la ligne vers le bas" })
map("i", "<A-k>", "<esc><cmd>m .-2<cr>==gi", { desc = "Déplacer la ligne vers le haut" })
map("v", "<A-j>", ":m '>+1<cr>gv=gv", { desc = "Déplacer la sélection vers le bas" })
map("v", "<A-k>", ":m '<-2<cr>gv=gv", { desc = "Déplacer la sélection vers le haut" })

-- Duplication de ligne
map("n", "<leader>d", "<cmd>t.<cr>", { desc = "Dupliquer la ligne" })

-- ===== RECHERCHE =====
-- Désactiver la surbrillance
map("n", "<Esc>", "<cmd>nohlsearch<cr>")

-- Recherche et remplacement
map("n", "<leader>sr", ":%s/\\<<C-r><C-w>\\>/<C-r><C-w>/gI<Left><Left><Left>", { desc = "Remplacer le mot sous le curseur" })

-- ===== BUFFERS =====
map("n", "<S-h>", "<cmd>bprevious<cr>", { desc = "Buffer précédent" })
map("n", "<S-l>", "<cmd>bnext<cr>", { desc = "Buffer suivant" })
map("n", "<leader>bb", "<cmd>e #<cr>", { desc = "Basculer vers l'autre buffer" })
map("n", "<leader>bd", "<cmd>bdelete<cr>", { desc = "Supprimer le buffer" })

-- ===== ONGLETS =====
map("n", "<leader>tn", "<cmd>tabnew<cr>", { desc = "Nouvel onglet" })
map("n", "<leader>tc", "<cmd>tabclose<cr>", { desc = "Fermer l'onglet" })
map("n", "<leader>to", "<cmd>tabonly<cr>", { desc = "Fermer les autres onglets" })

-- ===== SPLITS =====
map("n", "<leader>-", "<C-W>s", { desc = "Split horizontal" })
map("n", "<leader>|", "<C-W>v", { desc = "Split vertical" })

-- ===== FICHIERS =====
map("n", "<leader>w", "<cmd>w<cr>", { desc = "Sauvegarder" })
map("n", "<leader>W", "<cmd>wa<cr>", { desc = "Sauvegarder tout" })
map("n", "<leader>q", "<cmd>q<cr>", { desc = "Quitter" })
map("n", "<leader>Q", "<cmd>qa<cr>", { desc = "Quitter tout" })

-- ===== TERMINAL =====
map("t", "<Esc><Esc>", "<C-\\><C-n>", { desc = "Sortir du mode terminal" })
map("t", "<C-h>", "<cmd>wincmd h<cr>", { desc = "Aller à la fenêtre de gauche" })
map("t", "<C-j>", "<cmd>wincmd j<cr>", { desc = "Aller à la fenêtre du bas" })
map("t", "<C-k>", "<cmd>wincmd k<cr>", { desc = "Aller à la fenêtre du haut" })
map("t", "<C-l>", "<cmd>wincmd l<cr>", { desc = "Aller à la fenêtre de droite" })

-- ===== DÉVELOPPEMENT =====
-- Formatage
map("n", "<leader>cf", function()
  vim.lsp.buf.format({ async = true })
end, { desc = "Formater le code" })

-- Diagnostics
map("n", "<leader>cd", vim.diagnostic.open_float, { desc = "Diagnostic de ligne" })
map("n", "[d", vim.diagnostic.goto_prev, { desc = "Diagnostic précédent" })
map("n", "]d", vim.diagnostic.goto_next, { desc = "Diagnostic suivant" })

-- LSP
map("n", "gd", vim.lsp.buf.definition, { desc = "Aller à la définition" })
map("n", "gr", vim.lsp.buf.references, { desc = "Références" })
map("n", "gI", vim.lsp.buf.implementation, { desc = "Aller à l'implémentation" })
map("n", "gy", vim.lsp.buf.type_definition, { desc = "Définition de type" })
map("n", "gD", vim.lsp.buf.declaration, { desc = "Aller à la déclaration" })
map("n", "K", vim.lsp.buf.hover, { desc = "Documentation" })
map("n", "gK", vim.lsp.buf.signature_help, { desc = "Aide signature" })
map("n", "<leader>cr", vim.lsp.buf.rename, { desc = "Renommer" })
map("n", "<leader>ca", vim.lsp.buf.code_action, { desc = "Actions de code" })

-- ===== UTILITAIRES =====
-- Lazy
map("n", "<leader>l", "<cmd>Lazy<cr>", { desc = "Lazy" })

-- Inspection
map("n", "<leader>ui", vim.show_pos, { desc = "Inspecter la position" })

-- ===== DÉVELOPPEMENT SPÉCIFIQUE =====
-- PHP/Symfony
map("n", "<leader>ps", "<cmd>!symfony console<cr>", { desc = "Console Symfony" })
map("n", "<leader>pt", "<cmd>!vendor/bin/phpunit<cr>", { desc = "Tests PHPUnit" })

-- TypeScript/Node.js
map("n", "<leader>nr", "<cmd>!npm run<cr>", { desc = "Scripts npm" })
map("n", "<leader>nt", "<cmd>!npm test<cr>", { desc = "Tests npm" })

-- Git (si lazygit n'est pas disponible)
map("n", "<leader>gg", "<cmd>!git status<cr>", { desc = "Git status" })

-- ===== QUICKFIX =====
map("n", "<leader>xl", "<cmd>lopen<cr>", { desc = "Liste de localisation" })
map("n", "<leader>xq", "<cmd>copen<cr>", { desc = "Quickfix" })
map("n", "[q", vim.cmd.cprev, { desc = "Quickfix précédent" })
map("n", "]q", vim.cmd.cnext, { desc = "Quickfix suivant" })

-- ===== AUTRES =====
-- Ajouter des points d'annulation
map("i", ",", ",<c-g>u")
map("i", ".", ".<c-g>u")
map("i", ";", ";<c-g>u")

-- Sauvegarder en mode insertion
map("i", "<C-s>", "<cmd>w<cr><esc>", { desc = "Sauvegarder et sortir du mode insertion" })

-- Sélectionner tout
map("n", "<C-a>", "gg<S-v>G")

-- Nouvelle ligne sans entrer en mode insertion
map("n", "<leader>o", "o<Esc>", { desc = "Nouvelle ligne en dessous" })
map("n", "<leader>O", "O<Esc>", { desc = "Nouvelle ligne au dessus" })