-- Ultra-minimal Neovim configuration - no plugins
-- Basic settings
vim.opt.number = true
vim.opt.relativenumber = true
vim.opt.expandtab = true
vim.opt.shiftwidth = 2
vim.opt.tabstop = 2
vim.opt.termguicolors = true
vim.opt.scrolloff = 8
vim.opt.ignorecase = true
vim.opt.smartcase = true
vim.opt.hlsearch = true
vim.opt.incsearch = true

-- Set leader key  
vim.g.mapleader = " "

-- Basic keymaps
vim.keymap.set("n", "<leader>w", ":w<CR>")
vim.keymap.set("n", "<leader>q", ":q<CR>")
vim.keymap.set("n", "<leader>e", ":Ex<CR>")

-- Alt+hjkl navigation (matches tmux)
vim.keymap.set("n", "<M-h>", "<C-w>h")
vim.keymap.set("n", "<M-j>", "<C-w>j")  
vim.keymap.set("n", "<M-k>", "<C-w>k")
vim.keymap.set("n", "<M-l>", "<C-w>l")

-- Basic colorscheme (no plugins needed)
vim.cmd("colorscheme desert")

print("✨ Minimal nvim ready - no plugins!")