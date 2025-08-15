-- Désactiver snacks.nvim picker automatique pour dossiers
-- LazyVim utilise snacks_picker par défaut, ce qui cause l'ouverture de 4 fenêtres

return {
  {
    "folke/snacks.nvim",
    opts = {
      -- Désactiver le picker automatique qui s'ouvre sur les dossiers
      picker = { enabled = false },
      -- Garder les autres fonctionnalités snacks utiles
      bigfile = { enabled = true },
      notifier = { enabled = true },
      quickfile = { enabled = true },
      statuscolumn = { enabled = true },
      words = { enabled = true },
    },
  },
}