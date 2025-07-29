-- Formatting and Linting - PHP/Symfony focused
-- LazyVim-inspired configuration with conform.nvim

return {
  -- Formatting
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
          local ok, util = pcall(require, "util")
          if ok then
            util.format.format = function(opts)
              require("conform").format(opts)
            end
          end
        end,
      })
    end,
    opts = {
      -- Map of filetype to formatters
      formatters_by_ft = {
        -- PHP
        php = { "php_cs_fixer" },
        
        -- JavaScript/TypeScript
        javascript = { { "prettierd", "prettier" } },
        typescript = { { "prettierd", "prettier" } },
        javascriptreact = { { "prettierd", "prettier" } },
        typescriptreact = { { "prettierd", "prettier" } },
        
        -- Web
        html = { { "prettierd", "prettier" } },
        css = { { "prettierd", "prettier" } },
        scss = { { "prettierd", "prettier" } },
        
        -- Config files
        json = { { "prettierd", "prettier" } },
        jsonc = { { "prettierd", "prettier" } },
        yaml = { { "prettierd", "prettier" } },
        
        -- Lua
        lua = { "stylua" },
        
        -- Shell
        sh = { "shfmt" },
        bash = { "shfmt" },
        zsh = { "shfmt" },
        
        -- Markdown
        markdown = { { "prettierd", "prettier" } },
        ["markdown.mdx"] = { { "prettierd", "prettier" } },
        
        -- Twig (use prettier for now, could be improved)
        twig = { { "prettierd", "prettier" } },
      },
      
      -- Custom formatters
      formatters = {
        php_cs_fixer = {
          command = "php-cs-fixer",
          args = {
            "fix",
            "--rules=@PSR12,@Symfony",
            "--using-cache=no",
            "--show-progress=none",
            "$FILENAME",
          },
          stdin = false,
        },
        injected = {
          options = {
            ignore_errors = true,
            lang_to_formatters = {
              php = { "php_cs_fixer" },
              javascript = { "prettier" },
              typescript = { "prettier" },
              css = { "prettier" },
              html = { "prettier" },
            },
          },
        },
      },
      
      -- Format on save
      format_on_save = function(bufnr)
        -- Disable with a global or buffer-local variable
        if vim.g.disable_autoformat or vim.b[bufnr].disable_autoformat then
          return
        end
        return { timeout_ms = 500, lsp_fallback = true }
      end,
      
      -- Format after save for async formatters
      format_after_save = {
        lsp_fallback = true,
      },
      
      -- Log level
      log_level = vim.log.levels.ERROR,
      
      -- Notify on error
      notify_on_error = true,
    },
    config = function(_, opts)
      require("conform").setup(opts)
      
      -- Commands to toggle formatting
      vim.api.nvim_create_user_command("FormatDisable", function(args)
        if args.bang then
          -- FormatDisable! will disable formatting just for this buffer
          vim.b.disable_autoformat = true
        else
          vim.g.disable_autoformat = true
        end
      end, {
        desc = "Disable autoformat-on-save",
        bang = true,
      })
      
      vim.api.nvim_create_user_command("FormatEnable", function()
        vim.b.disable_autoformat = false
        vim.g.disable_autoformat = false
      end, {
        desc = "Re-enable autoformat-on-save",
      })
      
      -- PHP specific commands
      vim.api.nvim_create_user_command("PhpCsFix", function()
        local conform = require("conform")
        conform.format({
          formatters = { "php_cs_fixer" },
          timeout_ms = 10000,
        })
      end, {
        desc = "Format PHP with CS Fixer",
      })
    end,
  },

  -- Linting
  {
    "mfussenegger/nvim-lint",
    event = { "BufReadPre", "BufNewFile" },
    opts = {
      -- Event to trigger linters
      events = { "BufWritePost", "BufReadPost", "InsertLeave" },
      linters_by_ft = {
        -- PHP
        php = { "phpstan", "psalm" },
        
        -- JavaScript/TypeScript
        javascript = { "eslint_d" },
        typescript = { "eslint_d" },
        javascriptreact = { "eslint_d" },
        typescriptreact = { "eslint_d" },
        
        -- Shell
        sh = { "shellcheck" },
        bash = { "shellcheck" },
        zsh = { "shellcheck" },
        
        -- YAML
        yaml = { "yamllint" },
        
        -- JSON
        json = { "jsonlint" },
        
        -- Dockerfile
        dockerfile = { "hadolint" },
      },
      
      -- Custom linters
      linters = {
        phpstan = {
          cmd = "phpstan",
          stdin = false,
          args = {
            "analyse",
            "--error-format=json",
            "--no-progress",
            "--level=5",
          },
          stream = "stdout",
          ignore_exitcode = true,
          parser = function(output, bufnr)
            local diagnostics = {}
            local ok, decoded = pcall(vim.json.decode, output)
            if not ok or not decoded.files then
              return diagnostics
            end
            
            local filename = vim.api.nvim_buf_get_name(bufnr)
            local file_diagnostics = decoded.files[filename]
            if not file_diagnostics then
              return diagnostics
            end
            
            for _, msg in ipairs(file_diagnostics.messages) do
              table.insert(diagnostics, {
                lnum = msg.line - 1,
                col = 0,
                end_lnum = msg.line - 1,
                end_col = -1,
                severity = vim.diagnostic.severity.ERROR,
                message = msg.message,
                source = "phpstan",
              })
            end
            
            return diagnostics
          end,
        },
        
        psalm = {
          cmd = "psalm",
          stdin = false,
          args = {
            "--output-format=json",
            "--no-progress",
          },
          stream = "stdout",
          ignore_exitcode = true,
          parser = function(output, bufnr)
            local diagnostics = {}
            local ok, decoded = pcall(vim.json.decode, output)
            if not ok or not decoded then
              return diagnostics
            end
            
            local filename = vim.api.nvim_buf_get_name(bufnr)
            for _, msg in ipairs(decoded) do
              if msg.file_name == filename then
                local severity = vim.diagnostic.severity.ERROR
                if msg.severity == "info" then
                  severity = vim.diagnostic.severity.INFO
                elseif msg.severity == "warning" then
                  severity = vim.diagnostic.severity.WARN
                end
                
                table.insert(diagnostics, {
                  lnum = msg.line_from - 1,
                  col = msg.column_from - 1,
                  end_lnum = msg.line_to - 1,
                  end_col = msg.column_to,
                  severity = severity,
                  message = msg.message,
                  source = "psalm",
                })
              end
            end
            
            return diagnostics
          end,
        },
      },
    },
    config = function(_, opts)
      local lint = require("lint")
      
      for name, linter in pairs(opts.linters) do
        if type(linter) == "table" and type(lint.linters[name]) == "table" then
          lint.linters[name] = vim.tbl_deep_extend("force", lint.linters[name], linter)
        else
          lint.linters[name] = linter
        end
      end
      
      lint.linters_by_ft = opts.linters_by_ft
      
      local function debounce(ms, fn)
        local timer = vim.loop.new_timer()
        return function(...)
          local argv = { ... }
          timer:start(ms, 0, function()
            timer:stop()
            vim.schedule_wrap(fn)(unpack(argv))
          end)
        end
      end
      
      local function lint_fn()
        local names = lint._resolve_linter_by_ft(vim.bo.filetype)
        
        -- Filter out linters that don't exist or don't have executable
        names = vim.tbl_filter(function(name)
          local linter = lint.linters[name]
          if not linter then
            vim.notify("Linter not found: " .. name, vim.log.levels.WARN)
            return false
          end
          
          return true
        end, names)
        
        if #names > 0 then
          lint.try_lint(names)
        end
      end
      
      vim.api.nvim_create_autocmd(opts.events, {
        group = vim.api.nvim_create_augroup("nvim-lint", { clear = true }),
        callback = debounce(100, lint_fn),
      })
    end,
  },
}