-- Autocommands Neovim modernes
-- Automatisations pour améliorer l'expérience de développement

local function augroup(name)
  return vim.api.nvim_create_augroup("dotfiles_" .. name, { clear = true })
end

-- ===== INTERFACE =====
-- Surligner le texte copié
vim.api.nvim_create_autocmd("TextYankPost", {
  group = augroup("highlight_yank"),
  callback = function()
    vim.highlight.on_yank()
  end,
})

-- Redimensionner les splits quand la fenêtre est redimensionnée
vim.api.nvim_create_autocmd({ "VimResized" }, {
  group = augroup("resize_splits"),
  callback = function()
    local current_tab = vim.fn.tabpagenr()
    vim.cmd("tabdo wincmd =")
    vim.cmd("tabnext " .. current_tab)
  end,
})

-- Fermer certains types de fichiers avec 'q'
vim.api.nvim_create_autocmd("FileType", {
  group = augroup("close_with_q"),
  pattern = {
    "PlenaryTestPopup",
    "help",
    "lspinfo",
    "man",
    "notify",
    "qf",
    "query",
    "spectre_panel",
    "startuptime",
    "tsplayground",
    "neotest-output",
    "checkhealth",
    "neotest-summary",
    "neotest-output-panel",
  },
  callback = function(event)
    vim.bo[event.buf].buflisted = false
    vim.keymap.set("n", "q", "<cmd>close<cr>", { buffer = event.buf, silent = true })
  end,
})

-- ===== FICHIERS =====
-- Aller au dernier emplacement connu lors de l'ouverture d'un buffer
vim.api.nvim_create_autocmd("BufReadPost", {
  group = augroup("last_loc"),
  callback = function(event)
    local exclude = { "gitcommit" }
    local buf = event.buf
    if vim.tbl_contains(exclude, vim.bo[buf].filetype) or vim.b[buf].dotfiles_last_loc then
      return
    end
    vim.b[buf].dotfiles_last_loc = true
    local mark = vim.api.nvim_buf_get_mark(buf, '"')
    local lcount = vim.api.nvim_buf_line_count(buf)
    if mark[1] > 0 and mark[1] <= lcount then
      pcall(vim.api.nvim_win_set_cursor, 0, mark)
    end
  end,
})

-- Créer les répertoires manquants lors de la sauvegarde
vim.api.nvim_create_autocmd({ "BufWritePre" }, {
  group = augroup("auto_create_dir"),
  callback = function(event)
    if event.match:match("^%w%w+://") then
      return
    end
    local file = vim.loop.fs_realpath(event.match) or event.match
    vim.fn.mkdir(vim.fn.fnamemodify(file, ":p:h"), "p")
  end,
})

-- ===== DÉVELOPPEMENT =====
-- Configuration spécifique par type de fichier
vim.api.nvim_create_autocmd("FileType", {
  group = augroup("filetype_settings"),
  pattern = { "php", "javascript", "typescript", "json", "yaml", "html", "css", "scss" },
  callback = function()
    vim.opt_local.shiftwidth = 2
    vim.opt_local.tabstop = 2
    vim.opt_local.softtabstop = 2
  end,
})

-- Configuration pour les fichiers Markdown
vim.api.nvim_create_autocmd("FileType", {
  group = augroup("markdown_settings"),
  pattern = { "markdown" },
  callback = function()
    vim.opt_local.wrap = true
    vim.opt_local.spell = true
    vim.opt_local.conceallevel = 2
  end,
})

-- Configuration pour les fichiers de configuration
vim.api.nvim_create_autocmd("FileType", {
  group = augroup("config_files"),
  pattern = { "gitcommit", "gitrebase" },
  callback = function()
    vim.opt_local.spell = true
  end,
})

-- ===== LSP =====
-- Formater automatiquement avant la sauvegarde (optionnel)
vim.api.nvim_create_autocmd("BufWritePre", {
  group = augroup("lsp_format"),
  pattern = { "*.php", "*.ts", "*.tsx", "*.js", "*.jsx" },
  callback = function()
    -- Décommenter pour activer le formatage automatique
    -- vim.lsp.buf.format({ async = false })
  end,
})

-- ===== TERMINAL =====
-- Entrer automatiquement en mode insertion dans le terminal
vim.api.nvim_create_autocmd("TermOpen", {
  group = augroup("terminal_settings"),
  callback = function()
    vim.opt_local.number = false
    vim.opt_local.relativenumber = false
    vim.opt_local.scrolloff = 0
    vim.cmd("startinsert")
  end,
})

-- ===== PERFORMANCE =====
-- Désactiver certaines fonctionnalités pour les gros fichiers
vim.api.nvim_create_autocmd("BufReadPre", {
  group = augroup("big_file"),
  callback = function(event)
    local ok, stats = pcall(vim.loop.fs_stat, vim.api.nvim_buf_get_name(event.buf))
    if ok and stats and stats.size > 1024 * 1024 then -- 1MB
      vim.b[event.buf].big_file = true
      vim.opt_local.syntax = ""
      vim.opt_local.swapfile = false
      vim.opt_local.undofile = false
      vim.opt_local.breakindent = false
      vim.opt_local.colorcolumn = ""
      vim.opt_local.statuscolumn = ""
      vim.opt_local.signcolumn = "no"
      vim.opt_local.foldcolumn = "0"
      vim.opt_local.winbar = ""
    end
  end,
})

-- ===== PROJETS =====
-- Détection automatique du type de projet
vim.api.nvim_create_autocmd("VimEnter", {
  group = augroup("project_detection"),
  callback = function()
    local cwd = vim.fn.getcwd()
    
    -- Projet Symfony
    if vim.fn.filereadable(cwd .. "/bin/console") == 1 then
      vim.g.project_type = "symfony"
      vim.notify("🎼 Projet Symfony détecté", vim.log.levels.INFO)
    -- Projet Laravel
    elseif vim.fn.filereadable(cwd .. "/artisan") == 1 then
      vim.g.project_type = "laravel"
      vim.notify("🔥 Projet Laravel détecté", vim.log.levels.INFO)
    -- Projet Node.js/TypeScript
    elseif vim.fn.filereadable(cwd .. "/package.json") == 1 then
      vim.g.project_type = "nodejs"
      vim.notify("📦 Projet Node.js détecté", vim.log.levels.INFO)
    -- Projet Python
    elseif vim.fn.filereadable(cwd .. "/requirements.txt") == 1 or vim.fn.filereadable(cwd .. "/pyproject.toml") == 1 then
      vim.g.project_type = "python"
      vim.notify("🐍 Projet Python détecté", vim.log.levels.INFO)
    end
  end,
})

-- ===== SÉCURITÉ =====
-- Avertir pour les fichiers avec des permissions dangereuses
vim.api.nvim_create_autocmd("BufRead", {
  group = augroup("security_check"),
  callback = function(event)
    local file = event.match
    if file:match("%.env") or file:match("%.key") or file:match("id_rsa") then
      vim.notify("⚠️ Fichier sensible détecté: " .. vim.fn.fnamemodify(file, ":t"), vim.log.levels.WARN)
    end
  end,
})

-- ===== SESSIONS =====
-- Sauvegarder automatiquement la session (optionnel)
vim.api.nvim_create_autocmd("VimLeavePre", {
  group = augroup("auto_session"),
  callback = function()
    -- Décommenter pour activer la sauvegarde automatique de session
    -- if vim.fn.argc() == 0 then
    --   vim.cmd("mksession! ~/.config/nvim/session.vim")
    -- end
  end,
})