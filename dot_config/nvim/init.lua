-- Neovim Configuration moderne - Générée par Chezmoi
-- Optimisée pour PHP/Symfony et TypeScript
-- Thème: Catppuccin Mocha coordonné

-- ===== CONFIGURATION DE BASE =====
-- Désactiver les plugins par défaut pour améliorer les performances
vim.g.loaded_gzip = 1
vim.g.loaded_zip = 1
vim.g.loaded_zipPlugin = 1
vim.g.loaded_tar = 1
vim.g.loaded_tarPlugin = 1
vim.g.loaded_getscript = 1
vim.g.loaded_getscriptPlugin = 1
vim.g.loaded_vimball = 1
vim.g.loaded_vimballPlugin = 1
vim.g.loaded_2html_plugin = 1
vim.g.loaded_logiPat = 1
vim.g.loaded_rrhelper = 1
vim.g.loaded_netrw = 1
vim.g.loaded_netrwPlugin = 1
vim.g.loaded_netrwSettings = 1
vim.g.loaded_netrwFileHandlers = 1

-- Leader key
vim.g.mapleader = " "
vim.g.maplocalleader = " "

-- ===== CHARGEMENT DES MODULES =====
require("config.options")
require("config.keymaps")
require("config.autocmds")

-- Lazy.nvim - Gestionnaire de plugins moderne
local lazypath = vim.fn.stdpath("data") .. "/lazy/lazy.nvim"
if not vim.loop.fs_stat(lazypath) then
  vim.fn.system({
    "git",
    "clone",
    "--filter=blob:none",
    "https://github.com/folke/lazy.nvim.git",
    "--branch=stable",
    lazypath,
  })
end
vim.opt.rtp:prepend(lazypath)

-- Configuration Lazy.nvim
require("lazy").setup("plugins", {
  defaults = {
    lazy = true,
    version = false,
  },
  install = {
    missing = true,
    colorscheme = { "catppuccin" },
  },
  checker = {
    enabled = true,
    notify = false,
  },
  change_detection = {
    enabled = true,
    notify = false,
  },
  rocks = {
    enabled = false, -- Disable luarocks support
  },
  performance = {
    cache = {
      enabled = true,
    },
    rtp = {
      disabled_plugins = {
        "gzip",
        "matchit",
        "matchparen",
        "netrwPlugin",
        "tarPlugin",
        "tohtml",
        "tutor",
        "zipPlugin",
      },
    },
  },
})

-- Message de bienvenue
vim.api.nvim_create_autocmd("VimEnter", {
  callback = function()
    if vim.fn.argc() == 0 then
      vim.defer_fn(function()
        vim.notify("🚀 Neovim configuré avec Catppuccin Mocha", vim.log.levels.INFO, { title = "Dotfiles" })
      end, 100)
    end
  end,
})