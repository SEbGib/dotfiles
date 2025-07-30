-- Utilitaires pour la configuration Neovim
-- Module requis par certains plugins

local M = {}

-- Utilitaires LSP
M.lsp = {}

-- Gestion du renommage de fichiers avec LSP
-- Uses modern vim.lsp.get_clients() API (Neovim 0.10+)
function M.lsp.on_rename(from, to)
  local clients = vim.lsp.get_clients()
  for _, client in ipairs(clients) do
    if client.supports_method("workspace/willRenameFiles") then
      local resp = client.request_sync("workspace/willRenameFiles", {
        files = {
          {
            oldUri = vim.uri_from_fname(from),
            newUri = vim.uri_from_fname(to),
          },
        },
      }, 1000, 0)
      if resp and resp.result ~= nil then
        vim.lsp.util.apply_workspace_edit(resp.result, client.offset_encoding)
      end
    end
  end
end

-- Utilitaires de formatage
M.format = {}

function M.format.format(opts)
  local buf = vim.api.nvim_get_current_buf()
  if vim.b.disable_autoformat or vim.g.disable_autoformat then
    return
  end
  
  -- Try conform.nvim first, fallback to LSP
  local ok, conform = pcall(require, "conform")
  if ok then
    conform.format(vim.tbl_deep_extend("force", {
      bufnr = buf,
      lsp_fallback = true,
    }, opts or {}))
  else
    vim.lsp.buf.format(vim.tbl_deep_extend("force", {
      bufnr = buf,
    }, opts or {}))
  end
end

-- Utilitaires de diagnostic
M.diagnostic = {}

function M.diagnostic.get_diagnostics()
  local diagnostics = vim.diagnostic.get(0)
  local count = { errors = 0, warnings = 0, info = 0, hints = 0 }
  
  for _, diagnostic in ipairs(diagnostics) do
    if diagnostic.severity == vim.diagnostic.severity.ERROR then
      count.errors = count.errors + 1
    elseif diagnostic.severity == vim.diagnostic.severity.WARN then
      count.warnings = count.warnings + 1
    elseif diagnostic.severity == vim.diagnostic.severity.INFO then
      count.info = count.info + 1
    elseif diagnostic.severity == vim.diagnostic.severity.HINT then
      count.hints = count.hints + 1
    end
  end
  
  return count
end

-- Utilitaires UI
M.ui = {}

function M.ui.fg(name)
  local hl = vim.api.nvim_get_hl and vim.api.nvim_get_hl(0, { name = name }) or vim.api.nvim_get_hl_by_name(name, true)
  local fg = hl and hl.fg or hl.foreground
  return fg and { fg = string.format("#%06x", fg) }
end

-- Utilitaires de projet
M.project = {}

function M.project.is_symfony()
  return vim.fn.filereadable("symfony.lock") == 1 or vim.fn.filereadable("bin/console") == 1
end

function M.project.is_laravel()
  return vim.fn.filereadable("artisan") == 1
end

function M.project.is_node()
  return vim.fn.filereadable("package.json") == 1
end

function M.project.get_root()
  local patterns = { ".git", "composer.json", "package.json", "symfony.lock" }
  return vim.fs.dirname(vim.fs.find(patterns, { upward = true })[1])
end

return M