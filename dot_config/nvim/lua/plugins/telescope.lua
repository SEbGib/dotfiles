-- Telescope - Fuzzy finder with PHP/Symfony optimizations
-- LazyVim-inspired configuration

return {
  -- Telescope fuzzy finder
  {
    "nvim-telescope/telescope.nvim",
    cmd = "Telescope",
    version = false,
    dependencies = {
      "nvim-lua/plenary.nvim",
      {
        "nvim-telescope/telescope-fzf-native.nvim",
        build = "make",
        enabled = vim.fn.executable("make") == 1,
        config = function()
          require("telescope").load_extension("fzf")
        end,
      },
    },
    keys = {
      -- Files
      { "<leader>ff", "<cmd>Telescope find_files<cr>", desc = "Trouver fichiers" },
      { "<leader>fr", "<cmd>Telescope oldfiles<cr>", desc = "Fichiers récents" },
      { "<leader>fg", "<cmd>Telescope live_grep<cr>", desc = "Recherche dans fichiers" },
      { "<leader>fw", "<cmd>Telescope grep_string<cr>", desc = "Rechercher mot" },
      
      -- Buffers
      { "<leader>fb", "<cmd>Telescope buffers<cr>", desc = "Buffers" },
      
      -- Git
      { "<leader>gc", "<cmd>Telescope git_commits<cr>", desc = "Git commits" },
      { "<leader>gs", "<cmd>Telescope git_status<cr>", desc = "Git status" },
      { "<leader>gb", "<cmd>Telescope git_branches<cr>", desc = "Git branches" },
      
      -- LSP
      { "<leader>fs", "<cmd>Telescope lsp_document_symbols<cr>", desc = "Symboles document" },
      { "<leader>fS", "<cmd>Telescope lsp_workspace_symbols<cr>", desc = "Symboles workspace" },
      { "<leader>fd", "<cmd>Telescope diagnostics<cr>", desc = "Diagnostics" },
      
      -- Help
      { "<leader>fh", "<cmd>Telescope help_tags<cr>", desc = "Aide" },
      { "<leader>fk", "<cmd>Telescope keymaps<cr>", desc = "Keymaps" },
      { "<leader>fc", "<cmd>Telescope commands<cr>", desc = "Commandes" },
      
      -- PHP/Symfony specific
      { "<leader>pt", function()
          require("telescope.builtin").find_files({
            prompt_title = "Templates Twig",
            cwd = vim.fn.getcwd() .. "/templates",
            find_command = { "find", ".", "-name", "*.twig", "-type", "f" },
          })
        end, desc = "Templates Twig" },
      { "<leader>pe", function()
          require("telescope.builtin").find_files({
            prompt_title = "Entités",
            cwd = vim.fn.getcwd() .. "/src/Entity",
            find_command = { "find", ".", "-name", "*.php", "-type", "f" },
          })
        end, desc = "Entités" },
      { "<leader>pc", function()
          require("telescope.builtin").find_files({
            prompt_title = "Contrôleurs",
            cwd = vim.fn.getcwd() .. "/src/Controller",
            find_command = { "find", ".", "-name", "*.php", "-type", "f" },
          })
        end, desc = "Contrôleurs" },
      { "<leader>pr", function()
          require("telescope.builtin").find_files({
            prompt_title = "Repositories",
            cwd = vim.fn.getcwd() .. "/src/Repository",
            find_command = { "find", ".", "-name", "*.php", "-type", "f" },
          })
        end, desc = "Repositories" },
      { "<leader>pf", function()
          require("telescope.builtin").find_files({
            prompt_title = "Forms",
            cwd = vim.fn.getcwd() .. "/src/Form",
            find_command = { "find", ".", "-name", "*.php", "-type", "f" },
          })
        end, desc = "Forms" },
      { "<leader>ps", function()
          require("telescope.builtin").find_files({
            prompt_title = "Services",
            cwd = vim.fn.getcwd() .. "/src/Service",
            find_command = { "find", ".", "-name", "*.php", "-type", "f" },
          })
        end, desc = "Services" },
    },
    opts = function()
      local actions = require("telescope.actions")
      
      return {
        defaults = {
          prompt_prefix = " ",
          selection_caret = " ",
          multi_icon = " ",
          
          -- Layout
          layout_strategy = "horizontal",
          layout_config = {
            horizontal = {
              prompt_position = "top",
              preview_width = 0.55,
              results_width = 0.8,
            },
            vertical = {
              mirror = false,
            },
            width = 0.87,
            height = 0.80,
            preview_cutoff = 120,
          },
          
          -- Sorting
          sorting_strategy = "ascending",
          
          -- Files
          file_ignore_patterns = {
            "^.git/",
            "^node_modules/",
            "^vendor/",
            "^var/cache/",
            "^var/log/",
            "%.lock$",
            "%.min%.js$",
            "%.min%.css$",
          },
          
          -- Mappings
          mappings = {
            i = {
              ["<C-n>"] = actions.cycle_history_next,
              ["<C-p>"] = actions.cycle_history_prev,
              ["<C-j>"] = actions.move_selection_next,
              ["<C-k>"] = actions.move_selection_previous,
              ["<C-c>"] = actions.close,
              ["<Down>"] = actions.move_selection_next,
              ["<Up>"] = actions.move_selection_previous,
              ["<CR>"] = actions.select_default,
              ["<C-x>"] = actions.select_horizontal,
              ["<C-v>"] = actions.select_vertical,
              ["<C-t>"] = actions.select_tab,
              ["<C-u>"] = actions.preview_scrolling_up,
              ["<C-d>"] = actions.preview_scrolling_down,
              ["<PageUp>"] = actions.results_scrolling_up,
              ["<PageDown>"] = actions.results_scrolling_down,
              ["<Tab>"] = actions.toggle_selection + actions.move_selection_worse,
              ["<S-Tab>"] = actions.toggle_selection + actions.move_selection_better,
              ["<C-q>"] = actions.send_to_qflist + actions.open_qflist,
              ["<M-q>"] = actions.send_selected_to_qflist + actions.open_qflist,
              ["<C-l>"] = actions.complete_tag,
              ["<C-_>"] = actions.which_key,
            },
            n = {
              ["<esc>"] = actions.close,
              ["<CR>"] = actions.select_default,
              ["<C-x>"] = actions.select_horizontal,
              ["<C-v>"] = actions.select_vertical,
              ["<C-t>"] = actions.select_tab,
              ["<Tab>"] = actions.toggle_selection + actions.move_selection_worse,
              ["<S-Tab>"] = actions.toggle_selection + actions.move_selection_better,
              ["<C-q>"] = actions.send_to_qflist + actions.open_qflist,
              ["<M-q>"] = actions.send_selected_to_qflist + actions.open_qflist,
              ["j"] = actions.move_selection_next,
              ["k"] = actions.move_selection_previous,
              ["H"] = actions.move_to_top,
              ["M"] = actions.move_to_middle,
              ["L"] = actions.move_to_bottom,
              ["<Down>"] = actions.move_selection_next,
              ["<Up>"] = actions.move_selection_previous,
              ["gg"] = actions.move_to_top,
              ["G"] = actions.move_to_bottom,
              ["<C-u>"] = actions.preview_scrolling_up,
              ["<C-d>"] = actions.preview_scrolling_down,
              ["<PageUp>"] = actions.results_scrolling_up,
              ["<PageDown>"] = actions.results_scrolling_down,
              ["?"] = actions.which_key,
            },
          },
        },
        
        pickers = {
          find_files = {
            find_command = { "rg", "--files", "--hidden", "--glob", "!**/.git/*" },
            theme = "dropdown",
            previewer = false,
          },
          live_grep = {
            additional_args = function()
              return { "--hidden" }
            end,
          },
          grep_string = {
            additional_args = function()
              return { "--hidden" }
            end,
          },
          buffers = {
            theme = "dropdown",
            previewer = false,
            initial_mode = "normal",
            mappings = {
              i = {
                ["<C-d>"] = actions.delete_buffer,
              },
              n = {
                ["dd"] = actions.delete_buffer,
              },
            },
          },
          oldfiles = {
            theme = "dropdown",
            previewer = false,
          },
          lsp_references = {
            theme = "dropdown",
            initial_mode = "normal",
          },
          lsp_definitions = {
            theme = "dropdown",
            initial_mode = "normal",
          },
          lsp_document_symbols = {
            theme = "dropdown",
            initial_mode = "normal",
          },
          lsp_workspace_symbols = {
            theme = "dropdown",
            initial_mode = "normal",
          },
          diagnostics = {
            theme = "ivy",
            initial_mode = "normal",
            layout_config = {
              preview_cutoff = 9999,
            },
          },
        },
        
        extensions = {
          fzf = {
            fuzzy = true,
            override_generic_sorter = true,
            override_file_sorter = true,
            case_mode = "smart_case",
          },
        },
      }
    end,
    config = function(_, opts)
      require("telescope").setup(opts)
      
      -- Custom PHP/Symfony functions
      local builtin = require("telescope.builtin")
      local utils = require("telescope.utils")
      
      -- Find Symfony config files
      vim.keymap.set("n", "<leader>pC", function()
        builtin.find_files({
          prompt_title = "Config Symfony",
          cwd = vim.fn.getcwd() .. "/config",
          find_command = { "find", ".", "-name", "*.yaml", "-o", "-name", "*.yml", "-o", "-name", "*.php", "-type", "f" },
        })
      end, { desc = "Config Symfony" })
      
      -- Find migrations
      vim.keymap.set("n", "<leader>pm", function()
        builtin.find_files({
          prompt_title = "Migrations",
          cwd = vim.fn.getcwd() .. "/migrations",
          find_command = { "find", ".", "-name", "*.php", "-type", "f" },
        })
      end, { desc = "Migrations" })
      
      -- Search in Symfony logs
      vim.keymap.set("n", "<leader>pl", function()
        builtin.live_grep({
          prompt_title = "Logs Symfony",
          cwd = vim.fn.getcwd() .. "/var/log",
          additional_args = function()
            return { "--type", "f" }
          end,
        })
      end, { desc = "Logs Symfony" })
    end,
  },
}