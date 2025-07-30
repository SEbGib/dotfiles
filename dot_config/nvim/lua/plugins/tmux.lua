-- Tmux integration for Neovim
-- Seamless navigation between Neovim and tmux panes

return {
  -- Tmux navigation
  {
    "christoomey/vim-tmux-navigator",
    cmd = {
      "TmuxNavigateLeft",
      "TmuxNavigateDown",
      "TmuxNavigateUp",
      "TmuxNavigateRight",
      "TmuxNavigatePrevious",
    },
    keys = {
      { "<M-h>", "<cmd>TmuxNavigateLeft<cr>", desc = "Aller à gauche (tmux)" },
      { "<M-j>", "<cmd>TmuxNavigateDown<cr>", desc = "Aller en bas (tmux)" },
      { "<M-k>", "<cmd>TmuxNavigateUp<cr>", desc = "Aller en haut (tmux)" },
      { "<M-l>", "<cmd>TmuxNavigateRight<cr>", desc = "Aller à droite (tmux)" },
      { "<M-\\>", "<cmd>TmuxNavigatePrevious<cr>", desc = "Pane précédent (tmux)" },
    },
    init = function()
      -- Disable tmux navigator when zooming the Vim pane
      vim.g.tmux_navigator_disable_when_zoomed = 1
      
      -- Save on switch
      vim.g.tmux_navigator_save_on_switch = 2
      
      -- Preserve zoom
      vim.g.tmux_navigator_preserve_zoom = 1
    end,
  },

  -- Tmux integration utilities
  {
    "preservim/vimux",
    cond = function()
      return vim.env.TMUX ~= nil
    end,
    keys = {
      { "<leader>vp", "<cmd>VimuxPromptCommand<cr>", desc = "Tmux: Prompt Command" },
      { "<leader>vl", "<cmd>VimuxRunLastCommand<cr>", desc = "Tmux: Run Last Command" },
      { "<leader>vi", "<cmd>VimuxInspectRunner<cr>", desc = "Tmux: Inspect Runner" },
      { "<leader>vq", "<cmd>VimuxCloseRunner<cr>", desc = "Tmux: Close Runner" },
      { "<leader>vx", "<cmd>VimuxInterruptRunner<cr>", desc = "Tmux: Interrupt Runner" },
      { "<leader>vz", "<cmd>VimuxZoomRunner<cr>", desc = "Tmux: Zoom Runner" },
      
      -- PHP/Symfony specific commands
      { "<leader>pt", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('vendor/bin/phpunit')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Run PHPUnit" },
      { "<leader>pcc", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('symfony console cache:clear')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Clear Symfony Cache" },
      { "<leader>pcs", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('symfony console')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Symfony Console" },
      { "<leader>pss", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('symfony serve')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Symfony Serve" },
      
      -- Node.js/TypeScript commands
      { "<leader>nt", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('npm test')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Run npm test" },
      { "<leader>nd", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('npm run dev')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Run npm dev" },
      { "<leader>nb", function()
          if vim.env.TMUX then
            vim.cmd("VimuxRunCommand('npm run build')")
          else
            vim.notify("Not in tmux session", vim.log.levels.WARN)
          end
        end, desc = "Tmux: Run npm build" },
    },
    init = function()
      -- Vimux configuration
      vim.g.VimuxHeight = "30"
      vim.g.VimuxOrientation = "v"
      vim.g.VimuxUseNearest = 0
      vim.g.VimuxResetSequence = ""
    end,
  },

  -- Tmux clipboard integration
  {
    "roxma/vim-tmux-clipboard",
    event = { "BufReadPost", "BufNewFile" },
    cond = function()
      return vim.env.TMUX ~= nil
    end,
  },

  -- Enhanced tmux statusline integration
  {
    "edkolev/tmuxline.vim",
    cond = function()
      return vim.env.TMUX ~= nil
    end,
    cmd = { "Tmuxline", "TmuxlineSnapshot" },
    keys = {
      { "<leader>ut", "<cmd>Tmuxline<cr>", desc = "Update Tmux statusline" },
    },
    init = function()
      -- Tmuxline configuration to match our Catppuccin theme
      vim.g.tmuxline_preset = {
        a = "#S",
        b = "#W",
        c = "#H",
        win = "#I #W",
        cwin = "#I #W",
        x = "%d/%m",
        y = "%H:%M",
        z = "#h"
      }
      
      -- Disable tmuxline by default (we have our custom config)
      vim.g.tmuxline_powerline_separators = 0
    end,
  },
}