-- Options Neovim modernes
-- Configuration optimisée pour le développement

local opt = vim.opt

-- ===== INTERFACE =====
opt.number = true
opt.relativenumber = true
opt.signcolumn = "yes"
opt.cursorline = true
opt.colorcolumn = "80,120"
opt.wrap = false
opt.scrolloff = 8
opt.sidescrolloff = 8
opt.showmode = false
opt.showcmd = true
opt.cmdheight = 1
opt.laststatus = 3 -- Statusline globale
opt.winminwidth = 5
opt.pumheight = 10
opt.pumblend = 10
opt.winblend = 0

-- ===== ÉDITION =====
opt.expandtab = true
opt.shiftwidth = 4
opt.tabstop = 4
opt.softtabstop = 4
opt.smartindent = true
opt.autoindent = true
opt.breakindent = true
opt.linebreak = true
opt.backspace = "indent,eol,start"

-- ===== RECHERCHE =====
opt.ignorecase = true
opt.smartcase = true
opt.hlsearch = true
opt.incsearch = true
opt.grepprg = "rg --vimgrep"
opt.grepformat = "%f:%l:%c:%m"

-- ===== FICHIERS =====
opt.backup = false
opt.writebackup = false
opt.swapfile = false
opt.undofile = true
opt.undolevels = 10000
opt.autoread = true
opt.autowrite = true
opt.confirm = true
opt.hidden = true

-- ===== PERFORMANCE =====
opt.updatetime = 200
opt.timeout = true
opt.timeoutlen = 300
opt.ttimeoutlen = 0
opt.redrawtime = 10000
opt.maxmempattern = 20000
opt.synmaxcol = 240
opt.lazyredraw = false

-- ===== COMPLÉTION =====
opt.completeopt = "menu,menuone,noselect"
opt.shortmess:append({ W = true, I = true, c = true, C = true })
opt.iskeyword:append("-")

-- ===== FORMATAGE =====
opt.formatoptions = "jcroqlnt"
opt.textwidth = 0

-- ===== SPLITS =====
opt.splitbelow = true
opt.splitright = true
opt.splitkeep = "screen"

-- ===== SOURIS ET CLIPBOARD =====
opt.mouse = "a"
opt.clipboard = "unnamedplus"

-- ===== CARACTÈRES INVISIBLES =====
opt.list = true
opt.listchars = {
  tab = "→ ",
  eol = "↲",
  nbsp = "␣",
  trail = "•",
  extends = "⟩",
  precedes = "⟨",
}
opt.fillchars = {
  foldopen = "",
  foldclose = "",
  fold = " ",
  foldsep = " ",
  diff = "╱",
  eob = " ",
}

-- ===== FOLDING =====
opt.foldcolumn = "1"
opt.foldlevel = 99
opt.foldlevelstart = 99
opt.foldenable = true

-- ===== WILDMENU =====
opt.wildmode = "longest:full,full"
opt.wildoptions = "pum"

-- ===== SESSIONS =====
opt.sessionoptions = { "buffers", "curdir", "tabpages", "winsize", "help", "globals", "skiprtp", "folds" }

-- ===== CONFIGURATION SPÉCIFIQUE =====
-- Désactiver les providers non utilisés
vim.g.loaded_ruby_provider = 0
vim.g.loaded_perl_provider = 0
vim.g.loaded_python_provider = 0

-- Python3 provider
if vim.fn.executable("python3") == 1 then
  vim.g.python3_host_prog = vim.fn.exepath("python3")
end

-- Node.js provider
if vim.fn.executable("node") == 1 then
  vim.g.node_host_prog = vim.fn.exepath("node")
end

-- ===== VARIABLES GLOBALES =====
-- Configuration pour les plugins
vim.g.markdown_recommended_style = 0