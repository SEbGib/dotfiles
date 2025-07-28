-- Thème Catppuccin Mocha pour Neovim
-- Coordonné avec tmux et Starship

return {
  "catppuccin/nvim",
  name = "catppuccin",
  priority = 1000,
  opts = {
    flavour = "mocha",
    background = {
      light = "latte",
      dark = "mocha",
    },
    transparent_background = false,
    show_end_of_buffer = false,
    term_colors = true,
    dim_inactive = {
      enabled = false,
      shade = "dark",
      percentage = 0.15,
    },
    no_italic = false,
    no_bold = false,
    no_underline = false,
    styles = {
      comments = { "italic" },
      conditionals = { "italic" },
      loops = {},
      functions = {},
      keywords = {},
      strings = {},
      variables = {},
      numbers = {},
      booleans = {},
      properties = {},
      types = {},
      operators = {},
    },
    color_overrides = {},
    custom_highlights = function(colors)
      return {
        -- Personnalisations pour le développement
        ["@variable"] = { fg = colors.text },
        ["@variable.builtin"] = { fg = colors.red, style = { "italic" } },
        ["@constant"] = { fg = colors.peach },
        ["@constant.builtin"] = { fg = colors.peach, style = { "italic" } },
        
        -- PHP spécifique
        ["@variable.php"] = { fg = colors.text },
        ["@function.php"] = { fg = colors.blue },
        ["@method.php"] = { fg = colors.blue },
        ["@property.php"] = { fg = colors.teal },
        
        -- TypeScript/JavaScript spécifique
        ["@variable.typescript"] = { fg = colors.text },
        ["@function.typescript"] = { fg = colors.blue },
        ["@method.typescript"] = { fg = colors.blue },
        ["@property.typescript"] = { fg = colors.teal },
        
        -- Interface améliorée
        CursorLine = { bg = colors.surface0 },
        ColorColumn = { bg = colors.surface0 },
        SignColumn = { bg = colors.base },
        FoldColumn = { bg = colors.base, fg = colors.overlay0 },
        
        -- Diagnostics
        DiagnosticError = { fg = colors.red },
        DiagnosticWarn = { fg = colors.yellow },
        DiagnosticInfo = { fg = colors.sky },
        DiagnosticHint = { fg = colors.teal },
        
        -- Git
        GitSignsAdd = { fg = colors.green },
        GitSignsChange = { fg = colors.yellow },
        GitSignsDelete = { fg = colors.red },
      }
    end,
    integrations = {
      cmp = true,
      gitsigns = true,
      nvimtree = true,
      treesitter = true,
      notify = true,
      mini = true,
      telescope = {
        enabled = true,
      },
      lsp_trouble = true,
      which_key = true,
      indent_blankline = {
        enabled = true,
        colored_indent_levels = false,
      },
      native_lsp = {
        enabled = true,
        virtual_text = {
          errors = { "italic" },
          hints = { "italic" },
          warnings = { "italic" },
          information = { "italic" },
        },
        underlines = {
          errors = { "underline" },
          hints = { "underline" },
          warnings = { "underline" },
          information = { "underline" },
        },
        inlay_hints = {
          background = true,
        },
      },
    },
  },
  config = function(_, opts)
    require("catppuccin").setup(opts)
    vim.cmd.colorscheme("catppuccin")
  end,
}