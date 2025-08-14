-- Disable treesitter completely and use vim-polyglot instead
-- This overrides LazyVim's default treesitter configuration

return {
  -- Completely disable LazyVim's treesitter setup
  {
    "nvim-treesitter/nvim-treesitter",
    enabled = false,
  },
  {
    "nvim-treesitter/nvim-treesitter-textobjects", 
    enabled = false,
  },
  {
    "nvim-treesitter/nvim-treesitter-context",
    enabled = false,
  },
  {
    "windwp/nvim-ts-autotag",
    enabled = false,
  },

  -- Replace with stable vim-polyglot  
  {
    "sheerun/vim-polyglot",
    init = function()
      -- Disable conflicting languages
      vim.g.polyglot_disabled = {
        "autoindent",
        "sensible", 
      }
      
      -- Enable enhanced syntax for key languages
      vim.g.php_html_load = 1
      vim.g.php_html_in_heredoc = 1
      vim.g.php_html_in_nowdoc = 1
      vim.g.php_sql_query = 1
      vim.g.javascript_plugin_jsdoc = 1
      vim.g.javascript_plugin_ngdoc = 1
      vim.g.css3_syntax_highlight = 1
    end,
  },
}