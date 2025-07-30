-- LSP Configuration - PHP/Symfony focused
-- Inspired by LazyVim but optimized for PHP development

return {
  -- Mason for LSP management
  {
    "williamboman/mason.nvim",
    cmd = "Mason",
    build = ":MasonUpdate",
    opts = {
      ensure_installed = {
        "stylua",
        "shfmt",
        "php-cs-fixer",
        "phpstan",
        "prettier",
        "eslint_d",
        "typescript-language-server",
        "intelephense",
        "json-lsp",
        "yaml-language-server",
        "dockerfile-language-server",
        "docker-compose-language-service",
      },
    },
    config = function(_, opts)
      require("mason").setup(opts)
      local mr = require("mason-registry")
      mr:on("package:install:success", function()
        vim.defer_fn(function()
          -- trigger FileType event to possibly load this newly installed LSP server
          require("lazy.core.handler.event").trigger({
            event = "FileType",
            buf = vim.api.nvim_get_current_buf(),
          })
        end, 100)
      end)
      local function ensure_installed()
        for _, tool in ipairs(opts.ensure_installed) do
          local p = mr.get_package(tool)
          if not p:is_installed() then
            p:install()
          end
        end
      end
      if mr.refresh then
        mr.refresh(ensure_installed)
      else
        ensure_installed()
      end
    end,
  },

  -- LSP Configuration
  {
    "neovim/nvim-lspconfig",
    version = false, -- Use latest for modern API support
    event = { "BufReadPre", "BufNewFile" },
    dependencies = {
      "mason.nvim",
      "williamboman/mason-lspconfig.nvim",
      "hrsh7th/cmp-nvim-lsp",
    },
    opts = {
      -- Global diagnostics configuration
      diagnostics = {
        underline = true,
        update_in_insert = false,
        virtual_text = {
          spacing = 4,
          source = "if_many",
          prefix = "●",
        },
        severity_sort = true,
        signs = {
          text = {
            [vim.diagnostic.severity.ERROR] = "✘",
            [vim.diagnostic.severity.WARN] = "▲",
            [vim.diagnostic.severity.HINT] = "⚑",
            [vim.diagnostic.severity.INFO] = "»",
          },
        },
      },
      -- Inlay hints
      inlay_hints = {
        enabled = true,
      },
      -- Codelens
      codelens = {
        enabled = false,
      },
      -- Document highlighting
      document_highlight = {
        enabled = true,
      },
      -- Capabilities
      capabilities = {},
      -- Format options
      format = {
        formatting_options = nil,
        timeout_ms = nil,
      },
      -- Server configurations
      servers = {
        -- PHP Language Server (Intelephense)
        intelephense = {
          settings = {
            intelephense = {
              environment = {
                includePaths = {
                  vim.fn.getcwd() .. "/vendor/symfony",
                  vim.fn.getcwd() .. "/vendor",
                },
              },
              files = {
                maxSize = 5000000,
                associations = {
                  "*.php",
                  "*.phtml",
                  "*.inc",
                  "*.module",
                  "*.install",
                  "*.theme",
                },
                exclude = {
                  "**/node_modules/**",
                  "**/vendor/**/Tests/**",
                  "**/vendor/**/tests/**",
                  "**/var/cache/**",
                },
              },
              stubs = {
                "apache", "bcmath", "bz2", "calendar", "com_dotnet", "Core",
                "ctype", "curl", "date", "dba", "dom", "enchant", "exif",
                "FFI", "fileinfo", "filter", "fpm", "ftp", "gd", "gettext",
                "gmp", "hash", "iconv", "imap", "intl", "json", "ldap",
                "libxml", "mbstring", "meta", "mcrypt", "mysqli", "oci8",
                "odbc", "openssl", "pcntl", "pcre", "PDO", "pdo_ibm",
                "pdo_mysql", "pdo_pgsql", "pdo_sqlite", "pgsql", "Phar",
                "posix", "pspell", "readline", "Reflection", "session",
                "shmop", "SimpleXML", "snmp", "soap", "sockets", "sodium",
                "SPL", "sqlite3", "standard", "superglobals", "sysvmsg",
                "sysvsem", "sysvshm", "tidy", "tokenizer", "xml", "xmlreader",
                "xmlrpc", "xmlwriter", "xsl", "Zend OPcache", "zip", "zlib",
                "symfony", "doctrine", "phpunit",
              },
              diagnostics = {
                enable = true,
                run = "onType",
                embeddedLanguages = true,
              },
              completion = {
                insertUseDeclaration = true,
                fullyQualifyGlobalConstantsAndFunctions = false,
                suggestObjectOperatorStaticMethods = true,
                maxItems = 100,
              },
              format = {
                enable = true,
                braces = "psr12",
              },
            },
          },
        },
        
        -- TypeScript/JavaScript
        tsserver = {
          settings = {
            typescript = {
              inlayHints = {
                includeInlayParameterNameHints = "literal",
                includeInlayParameterNameHintsWhenArgumentMatchesName = false,
                includeInlayFunctionParameterTypeHints = true,
                includeInlayVariableTypeHints = false,
                includeInlayPropertyDeclarationTypeHints = true,
                includeInlayFunctionLikeReturnTypeHints = true,
                includeInlayEnumMemberValueHints = true,
              },
            },
            javascript = {
              inlayHints = {
                includeInlayParameterNameHints = "all",
                includeInlayParameterNameHintsWhenArgumentMatchesName = false,
                includeInlayFunctionParameterTypeHints = true,
                includeInlayVariableTypeHints = true,
                includeInlayPropertyDeclarationTypeHints = true,
                includeInlayFunctionLikeReturnTypeHints = true,
                includeInlayEnumMemberValueHints = true,
              },
            },
          },
        },
        
        -- JSON
        jsonls = {
          settings = {
            json = {
              schemas = {},
              validate = { enable = true },
            },
          },
        },
        
        -- YAML
        yamlls = {
          settings = {
            yaml = {
              schemaStore = {
                enable = true,
                url = "",
              },
              schemas = {},
            },
          },
        },
        
        -- Lua
        lua_ls = {
          settings = {
            Lua = {
              workspace = {
                checkThirdParty = false,
              },
              codeLens = {
                enable = true,
              },
              completion = {
                callSnippet = "Replace",
              },
              doc = {
                privateName = { "^_" },
              },
              hint = {
                enable = true,
                setType = false,
                paramType = true,
                paramName = "Disable",
                semicolon = "Disable",
                arrayIndex = "Disable",
              },
            },
          },
        },
      },
      -- Server setup hooks
      setup = {},
    },
    config = function(_, opts)
      local util = require("util")
      
      -- Setup diagnostics
      vim.diagnostic.config(vim.deepcopy(opts.diagnostics))
      
      -- Setup LSP handlers
      local servers = opts.servers
      local has_cmp, cmp_nvim_lsp = pcall(require, "cmp_nvim_lsp")
      local capabilities = vim.tbl_deep_extend(
        "force",
        {},
        vim.lsp.protocol.make_client_capabilities(),
        has_cmp and cmp_nvim_lsp.default_capabilities() or {},
        opts.capabilities or {}
      )
      
      local function setup(server)
        local server_opts = vim.tbl_deep_extend("force", {
          capabilities = vim.deepcopy(capabilities),
        }, servers[server] or {})
        
        -- Add schemastore integration for JSON/YAML servers
        if server == "jsonls" then
          local ok, schemastore = pcall(require, "schemastore")
          if ok then
            server_opts.settings.json.schemas = schemastore.json.schemas()
          end
        elseif server == "yamlls" then
          local ok, schemastore = pcall(require, "schemastore")
          if ok then
            server_opts.settings.yaml.schemas = schemastore.yaml.schemas()
          end
        end
        
        if opts.setup[server] then
          if opts.setup[server](server, server_opts) then
            return
          end
        elseif opts.setup["*"] then
          if opts.setup["*"](server, server_opts) then
            return
          end
        end
        require("lspconfig")[server].setup(server_opts)
      end
      
      -- Get all available servers
      local have_mason, mlsp = pcall(require, "mason-lspconfig")
      local all_mslp_servers = {}
      if have_mason then
        all_mslp_servers = vim.tbl_keys(require("mason-lspconfig.mappings.server").lspconfig_to_package)
      end
      
      local ensure_installed = {}
      for server, server_opts in pairs(servers) do
        if server_opts then
          server_opts = server_opts == true and {} or server_opts
          if server_opts.mason == false or not vim.tbl_contains(all_mslp_servers, server) then
            setup(server)
          else
            ensure_installed[#ensure_installed + 1] = server
          end
        end
      end
      
      if have_mason then
        mlsp.setup({ ensure_installed = ensure_installed, handlers = { setup } })
      end
      
      -- Setup keymaps and autocommands
      vim.api.nvim_create_autocmd("LspAttach", {
        callback = function(args)
          local buffer = args.buf
          local client = vim.lsp.get_client_by_id(args.data.client_id)
          
          -- Keymaps
          local function map(mode, lhs, rhs, desc)
            vim.keymap.set(mode, lhs, rhs, { buffer = buffer, desc = desc })
          end
          
          map("n", "gd", vim.lsp.buf.definition, "Aller à la définition")
          map("n", "gr", vim.lsp.buf.references, "Références")
          map("n", "gI", vim.lsp.buf.implementation, "Aller à l'implémentation")
          map("n", "gy", vim.lsp.buf.type_definition, "Définition de type")
          map("n", "gD", vim.lsp.buf.declaration, "Aller à la déclaration")
          map("n", "K", vim.lsp.buf.hover, "Documentation")
          map("n", "gK", vim.lsp.buf.signature_help, "Aide signature")
          map("i", "<C-k>", vim.lsp.buf.signature_help, "Aide signature")
          map("n", "<leader>cr", vim.lsp.buf.rename, "Renommer")
          map({ "n", "v" }, "<leader>ca", vim.lsp.buf.code_action, "Actions de code")
          map("n", "<leader>cf", function()
            util.format.format({ force = true })
          end, "Formater")
          
          -- PHP/Symfony specific keymaps
          if client and client.name == "intelephense" then
            map("n", "<leader>pi", "<cmd>!composer install<cr>", "Composer install")
            map("n", "<leader>pu", "<cmd>!composer update<cr>", "Composer update")
            map("n", "<leader>pc", "<cmd>!symfony console cache:clear<cr>", "Clear cache")
            map("n", "<leader>pm", "<cmd>!symfony console make:<cr>", "Symfony make")
          end
          
          -- Document highlighting
          if client and client.server_capabilities.documentHighlightProvider then
            vim.api.nvim_create_autocmd({ "CursorHold", "CursorHoldI" }, {
              buffer = buffer,
              callback = vim.lsp.buf.document_highlight,
            })
            vim.api.nvim_create_autocmd({ "CursorMoved", "CursorMovedI" }, {
              buffer = buffer,
              callback = vim.lsp.buf.clear_references,
            })
          end
          
          -- Inlay hints
          if client and client.server_capabilities.inlayHintProvider and vim.lsp.inlay_hint then
            map("n", "<leader>uh", function()
              vim.lsp.inlay_hint.enable(not vim.lsp.inlay_hint.is_enabled())
            end, "Toggle Inlay Hints")
          end
        end,
      })
    end,
  },

  -- Mason - Package manager for LSP servers
  {
    "williamboman/mason.nvim",
    cmd = "Mason",
    keys = { { "<leader>cm", "<cmd>Mason<cr>", desc = "Mason" } },
    build = ":MasonUpdate",
    opts = {
      ensure_installed = {
        -- PHP
        "intelephense",
        "php-cs-fixer",
        "phpstan",
        "psalm",
        
        -- TypeScript/JavaScript
        "typescript-language-server",
        "eslint-lsp",
        "prettier",
        
        -- Web
        "html-lsp",
        "css-lsp",
        "tailwindcss-language-server",
        
        -- Config files
        "json-lsp",
        "yaml-language-server",
        
        -- Lua
        "lua-language-server",
        "stylua",
        
        -- Shell
        "bash-language-server",
        "shellcheck",
        "shfmt",
      },
    },
    config = function(_, opts)
      require("mason").setup(opts)
      local mr = require("mason-registry")
      mr:on("package:install:success", function()
        vim.defer_fn(function()
          require("lazy.core.handler.event").trigger({
            event = "FileType",
            buf = vim.api.nvim_get_current_buf(),
          })
        end, 100)
      end)
      local function ensure_installed()
        for _, tool in ipairs(opts.ensure_installed) do
          local p = mr.get_package(tool)
          if not p:is_installed() then
            p:install()
          end
        end
      end
      if mr.refresh then
        mr.refresh(ensure_installed)
      else
        ensure_installed()
      end
    end,
  },

  -- Schema store for JSON/YAML
  {
    "b0o/schemastore.nvim",
    lazy = true,
    version = false,
  },
}