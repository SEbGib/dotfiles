-- Coding enhancements - LazyVim inspired

return {
  -- Better formatting
  {
    "stevearc/conform.nvim",
    dependencies = { "mason.nvim" },
    lazy = true,
    cmd = "ConformInfo",
    keys = {
      {
        "<leader>cF",
        function()
          require("conform").format({ formatters = { "injected" } })
        end,
        mode = { "n", "v" },
        desc = "Format Injected Langs",
      },
    },
    init = function()
      -- Install conform on VeryLazy
      vim.api.nvim_create_autocmd("User", {
        pattern = "VeryLazy",
        callback = function()
          require("conform").setup({
            formatters_by_ft = {
              lua = { "stylua" },
              fish = { "fish_indent" },
              sh = { "shfmt" },
              php = { "php_cs_fixer" },
              javascript = { { "prettierd", "prettier" } },
              typescript = { { "prettierd", "prettier" } },
              javascriptreact = { { "prettierd", "prettier" } },
              typescriptreact = { { "prettierd", "prettier" } },
              vue = { { "prettierd", "prettier" } },
              css = { { "prettierd", "prettier" } },
              scss = { { "prettierd", "prettier" } },
              less = { { "prettierd", "prettier" } },
              html = { { "prettierd", "prettier" } },
              json = { { "prettierd", "prettier" } },
              jsonc = { { "prettierd", "prettier" } },
              yaml = { { "prettierd", "prettier" } },
              markdown = { { "prettierd", "prettier" } },
              ["markdown.mdx"] = { { "prettierd", "prettier" } },
              graphql = { { "prettierd", "prettier" } },
            },
            format_on_save = function(bufnr)
              -- Disable with a global or buffer-local variable
              if vim.g.disable_autoformat or vim.b[bufnr].disable_autoformat then
                return
              end
              return { timeout_ms = 500, lsp_fallback = true }
            end,
            formatters = {
              injected = { options = { ignore_errors = true } },
              php_cs_fixer = {
                command = "php-cs-fixer",
                args = {
                  "fix",
                  "--rules=@PSR12",
                  "$FILENAME",
                },
                stdin = false,
              },
            },
          })
        end,
      })
    end,
  },

  -- Linting
  {
    "mfussenegger/nvim-lint",
    event = { "BufReadPre", "BufNewFile" },
    config = function()
      local lint = require("lint")
      lint.linters_by_ft = {
        php = { "phpstan" },
        javascript = { "eslint_d" },
        typescript = { "eslint_d" },
        javascriptreact = { "eslint_d" },
        typescriptreact = { "eslint_d" },
      }

      local lint_augroup = vim.api.nvim_create_augroup("lint", { clear = true })
      vim.api.nvim_create_autocmd({ "BufEnter", "BufWritePost", "InsertLeave" }, {
        group = lint_augroup,
        callback = function()
          lint.try_lint()
        end,
      })
    end,
  },

  -- Snippets
  {
    "L3MON4D3/LuaSnip",
    build = (function()
      if vim.fn.has("win32") == 1 or vim.fn.executable("make") == 0 then
        return
      end
      return "make install_jsregexp"
    end)(),
    dependencies = {
      {
        "rafamadriz/friendly-snippets",
        config = function()
          require("luasnip.loaders.from_vscode").lazy_load()
        end,
      },
    },
    opts = {
      history = true,
      delete_check_events = "TextChanged",
    },
    keys = {
      {
        "<tab>",
        function()
          return require("luasnip").jumpable(1) and "<Plug>luasnip-jump-next" or "<tab>"
        end,
        expr = true,
        silent = true,
        mode = "i",
      },
      { "<tab>", function() require("luasnip").jump(1) end, mode = "s" },
      { "<s-tab>", function() require("luasnip").jump(-1) end, mode = { "i", "s" } },
    },
  },

  -- Better PHP support
  {
    "gbprod/phpactor.nvim",
    ft = "php",
    build = function()
      require("phpactor.handler.update")()
    end,
    dependencies = {
      "nvim-lua/plenary.nvim",
      "neovim/nvim-lspconfig",
    },
    opts = {
      install = {
        path = vim.fn.stdpath("data") .. "/phpactor",
        branch = "master",
        bin = vim.fn.stdpath("data") .. "/phpactor/bin/phpactor",
        php_bin = "php",
        composer_bin = "composer",
        git_bin = "git",
        check_on_startup = "none",
      },
      lspconfig = {
        enabled = false, -- We use intelephense
        options = {},
      },
    },
    keys = {
      { "<leader>pm", ":PhpactorContextMenu<cr>", desc = "Phpactor Menu" },
      { "<leader>pn", ":PhpactorClassNew<cr>", desc = "New Class" },
    },
  },

  -- TypeScript support
  {
    "pmizio/typescript-tools.nvim",
    ft = { "typescript", "typescriptreact", "javascript", "javascriptreact" },
    dependencies = { "nvim-lua/plenary.nvim", "neovim/nvim-lspconfig" },
    opts = {},
  },

  -- Markdown preview
  {
    "iamcco/markdown-preview.nvim",
    cmd = { "MarkdownPreviewToggle", "MarkdownPreview", "MarkdownPreviewStop" },
    build = function() vim.fn["mkdp#util#install"]() end,
    keys = {
      {
        "<leader>cp",
        ft = "markdown",
        "<cmd>MarkdownPreviewToggle<cr>",
        desc = "Markdown Preview",
      },
    },
    config = function()
      vim.cmd([[do FileType]])
    end,
  },
}