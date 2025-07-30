-- Completion and Snippets - PHP/Symfony optimized
-- LazyVim-inspired completion setup

return {
  -- Auto completion
  {
    "hrsh7th/nvim-cmp",
    version = false,
    event = "InsertEnter",
    dependencies = {
      "hrsh7th/cmp-nvim-lsp",
      "hrsh7th/cmp-buffer",
      "hrsh7th/cmp-path",
      "hrsh7th/cmp-cmdline",
      "saadparwaiz1/cmp_luasnip",
    },
    opts = function()
      vim.api.nvim_set_hl(0, "CmpGhostText", { link = "Comment", default = true })
      local cmp = require("cmp")
      local defaults = require("cmp.config.default")()
      
      return {
        completion = {
          completeopt = "menu,menuone,noinsert",
        },
        snippet = {
          expand = function(args)
            require("luasnip").lsp_expand(args.body)
          end,
        },
        mapping = cmp.mapping.preset.insert({
          ["<C-n>"] = cmp.mapping.select_next_item({ behavior = cmp.SelectBehavior.Insert }),
          ["<C-p>"] = cmp.mapping.select_prev_item({ behavior = cmp.SelectBehavior.Insert }),
          ["<C-b>"] = cmp.mapping.scroll_docs(-4),
          ["<C-f>"] = cmp.mapping.scroll_docs(4),
          ["<C-Space>"] = cmp.mapping.complete(),
          ["<C-e>"] = cmp.mapping.abort(),
          ["<CR>"] = cmp.mapping.confirm({ select = true }),
          ["<S-CR>"] = cmp.mapping.confirm({
            behavior = cmp.ConfirmBehavior.Replace,
            select = true,
          }),
          ["<C-CR>"] = function(fallback)
            cmp.abort()
            fallback()
          end,
        }),
        sources = cmp.config.sources({
          { name = "nvim_lsp" },
          { name = "luasnip" },
          { name = "path" },
          { name = "path" },
        }, {
          { name = "buffer" },
        }),
        formatting = {
          format = function(entry, vim_item)
            local icons = {
              Text = "󰉿",
              Method = "󰆧",
              Function = "󰊕",
              Constructor = "",
              Field = "󰜢",
              Variable = "󰀫",
              Class = "󰠱",
              Interface = "",
              Module = "",
              Property = "󰜢",
              Unit = "󰑭",
              Value = "󰎠",
              Enum = "",
              Keyword = "󰌋",
              Snippet = "",
              Color = "󰏘",
              File = "󰈙",
              Reference = "󰈇",
              Folder = "󰉋",
              EnumMember = "",
              Constant = "󰏿",
              Struct = "󰙅",
              Event = "",
              Operator = "󰆕",
              TypeParameter = "",
            }
            
            if icons[vim_item.kind] then
              vim_item.kind = icons[vim_item.kind] .. " " .. vim_item.kind
            end
            
            -- Source indicators
            local source_names = {
              nvim_lsp = "[LSP]",
              luasnip = "[Snip]",
              buffer = "[Buf]",
              path = "[Path]",
            }
            vim_item.menu = source_names[entry.source.name] or "[?]"
            
            return vim_item
          end,
        },
        experimental = {
          ghost_text = {
            hl_group = "CmpGhostText",
          },
        },
        sorting = defaults.sorting,
      }
    end,
    config = function(_, opts)
      for _, source in ipairs(opts.sources) do
        source.group_index = source.group_index or 1
      end
      require("cmp").setup(opts)
      
      -- Command line completion
      local cmp = require("cmp")
      cmp.setup.cmdline({ "/", "?" }, {
        mapping = cmp.mapping.preset.cmdline(),
        sources = {
          { name = "buffer" }
        }
      })
      
      cmp.setup.cmdline(":", {
        mapping = cmp.mapping.preset.cmdline(),
        sources = cmp.config.sources({
          { name = "path" }
        }, {
          { name = "cmdline" }
        })
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
    config = function(_, opts)
      require("luasnip").setup(opts)
      
      -- Custom PHP/Symfony snippets
      local ls = require("luasnip")
      local s = ls.snippet
      local t = ls.text_node
      local i = ls.insert_node
      local f = ls.function_node
      
      ls.add_snippets("php", {
        -- Symfony Controller
        s("controller", {
          t({"<?php", "", "namespace App\\Controller;", "", "use Symfony\\Bundle\\FrameworkBundle\\Controller\\AbstractController;", "use Symfony\\Component\\HttpFoundation\\Response;", "use Symfony\\Component\\Routing\\Annotation\\Route;", "", "class "}),
          i(1, "HomeController"),
          t({" extends AbstractController", "{", "    #[Route('/"}),
          i(2, "home"),
          t({"', name: '"}),
          i(3, "app_home"),
          t({"')]", "    public function "}),
          i(4, "index"),
          t({"(): Response", "    {", "        return $this->render('"}),
          i(5, "home/index.html.twig"),
          t({"', [", "            "}),
          i(6),
          t({"", "        ]);", "    }", "}"}),
        }),
        
        -- Symfony Entity
        s("entity", {
          t({"<?php", "", "namespace App\\Entity;", "", "use Doctrine\\ORM\\Mapping as ORM;", "", "#[ORM\\Entity(repositoryClass: "}),
          f(function(args) return args[1][1] .. "Repository::class" end, {1}),
          t({")]", "class "}),
          i(1, "User"),
          t({"", "{", "    #[ORM\\Id]", "    #[ORM\\GeneratedValue]", "    #[ORM\\Column]", "    private ?int $id = null;", "", "    "}),
          i(2),
          t({"", "", "    public function getId(): ?int", "    {", "        return $this->id;", "    }", "}"}),
        }),
        
        -- PHP Class
        s("class", {
          t({"<?php", "", "namespace "}),
          i(1, "App"),
          t({";", "", "class "}),
          i(2, "ClassName"),
          t({"", "{", "    "}),
          i(3),
          t({"", "}"}),
        }),
        
        -- PHP Method
        s("method", {
          t("public function "),
          i(1, "methodName"),
          t("("),
          i(2),
          t("): "),
          i(3, "void"),
          t({"", "{", "    "}),
          i(4),
          t({"", "}"}),
        }),
        
        -- Twig template
        s("twig", {
          t("{% extends '"),
          i(1, "base.html.twig"),
          t({"' %}", "", "{% block "}),
          i(2, "title"),
          t(" %}"),
          i(3, "Page Title"),
          t({"{% endblock %}", "", "{% block "}),
          i(4, "body"),
          t({" %}", "    "}),
          i(5),
          t({"", "{% endblock %}"}),
        }),
      })
      
      -- TypeScript/JavaScript snippets
      ls.add_snippets("typescript", {
        s("interface", {
          t("interface "),
          i(1, "InterfaceName"),
          t({" {", "  "}),
          i(2),
          t({"", "}"}),
        }),
        
        s("component", {
          t("export interface "),
          i(1, "Component"),
          t({"Props {", "  "}),
          i(2),
          t({"", "}", "", "export function "}),
          f(function(args) return args[1][1] end, {1}),
          t("({ "),
          i(3),
          t(" }: "),
          f(function(args) return args[1][1] end, {1}),
          t({"Props) {", "  return (", "    <div>"}),
          i(4),
          t({"</div>", "  );", "}"}),
        }),
      })
    end,
  },

  -- Auto pairs
  {
    "windwp/nvim-autopairs",
    event = "InsertEnter",
    opts = {},
    config = function(_, opts)
      require("nvim-autopairs").setup(opts)
      
      -- Integration with cmp
      local cmp_autopairs = require("nvim-autopairs.completion.cmp")
      local cmp = require("cmp")
      cmp.event:on("confirm_done", cmp_autopairs.on_confirm_done())
    end,
  },
}