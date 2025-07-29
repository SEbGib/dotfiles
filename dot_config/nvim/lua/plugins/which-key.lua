-- Which-key - Keybinding help with PHP/Symfony focus
-- LazyVim-inspired configuration

return {
  {
    "folke/which-key.nvim",
    event = "VeryLazy",
    opts = {
      plugins = { spelling = true },
      defaults = {
        mode = { "n", "v" },
        ["g"] = { name = "+goto" },
        ["gz"] = { name = "+surround" },
        ["]"] = { name = "+next" },
        ["["] = { name = "+prev" },
        ["<leader><tab>"] = { name = "+tabs" },
        ["<leader>b"] = { name = "+buffer" },
        ["<leader>c"] = { name = "+code" },
        ["<leader>f"] = { name = "+file/find" },
        ["<leader>g"] = { name = "+git" },
        ["<leader>gh"] = { name = "+hunks" },
        ["<leader>q"] = { name = "+quit/session" },
        ["<leader>s"] = { name = "+search" },
        ["<leader>u"] = { name = "+ui" },
        ["<leader>w"] = { name = "+windows" },
        ["<leader>x"] = { name = "+diagnostics/quickfix" },
        -- PHP/Symfony specific groups
        ["<leader>p"] = { name = "+php/symfony" },
        ["<leader>pt"] = { name = "+templates" },
        ["<leader>pe"] = { name = "+entities" },
        ["<leader>pc"] = { name = "+controllers" },
        ["<leader>pr"] = { name = "+repositories" },
        ["<leader>pf"] = { name = "+forms" },
        ["<leader>ps"] = { name = "+services" },
        ["<leader>pm"] = { name = "+migrations" },
        ["<leader>pC"] = { name = "+config" },
        ["<leader>pl"] = { name = "+logs" },
      },
    },
    config = function(_, opts)
      local wk = require("which-key")
      wk.setup(opts)
      wk.register(opts.defaults)
    end,
  },
}