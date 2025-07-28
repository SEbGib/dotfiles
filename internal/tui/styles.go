package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Color palette - Catppuccin Mocha inspired
var (
	// Primary colors
	ColorPrimary   = lipgloss.Color("#cba6f7") // Mauve
	ColorSecondary = lipgloss.Color("#89b4fa") // Blue
	ColorAccent    = lipgloss.Color("#f38ba8") // Pink
	ColorSuccess   = lipgloss.Color("#a6e3a1") // Green
	ColorWarning   = lipgloss.Color("#f9e2af") // Yellow
	ColorError     = lipgloss.Color("#f38ba8") // Red
	ColorInfo      = lipgloss.Color("#74c7ec") // Sky

	// Background colors
	ColorBgPrimary   = lipgloss.Color("#1e1e2e") // Base
	ColorBgSecondary = lipgloss.Color("#313244") // Surface0
	ColorBgTertiary  = lipgloss.Color("#45475a") // Surface1

	// Text colors
	ColorTextPrimary   = lipgloss.Color("#cdd6f4") // Text
	ColorTextSecondary = lipgloss.Color("#bac2de") // Subtext1
	ColorTextMuted     = lipgloss.Color("#a6adc8") // Subtext0
	ColorTextDim       = lipgloss.Color("#6c7086") // Overlay1
)

// Base styles
var (
	// Main container style - minimal padding
	AppStyle = lipgloss.NewStyle().
			Padding(0, 1)

	// Header styles - simple and clean
	HeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			Padding(0, 1).
			Margin(0, 0, 1, 0)

	// Title - clean without overwhelming borders
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			Padding(1, 2).
			Margin(0, 0, 1, 0).
			Align(lipgloss.Center)

	// Subtitle style - improved readability
	SubtitleStyle = lipgloss.NewStyle().
			Foreground(ColorTextSecondary).
			Italic(true).
			Margin(0, 0, 1, 0).
			Padding(0, 2).
			Align(lipgloss.Center)
	// Menu item styles
	MenuItemStyle = lipgloss.NewStyle().
			Padding(0, 2).
			Margin(0, 1).
			Foreground(ColorTextPrimary)

	MenuItemSelectedStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Margin(0, 1).
				Foreground(ColorBgPrimary).
				Background(ColorPrimary).
				Bold(true)

	// Status styles
	StatusSuccessStyle = lipgloss.NewStyle().
				Foreground(ColorSuccess).
				Bold(true).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorSuccess)

	StatusWarningStyle = lipgloss.NewStyle().
				Foreground(ColorWarning).
				Bold(true).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorWarning)

	StatusErrorStyle = lipgloss.NewStyle().
				Foreground(ColorError).
				Bold(true).
				Padding(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorError)

	StatusInfoStyle = lipgloss.NewStyle().
			Foreground(ColorInfo).
			Bold(true).
			Padding(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorInfo)

	// Progress bar styles
	ProgressBarStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorSecondary).
				Padding(0, 1)

	ProgressFillStyle = lipgloss.NewStyle().
				Background(ColorSuccess).
				Foreground(ColorBgPrimary)

	// Card styles - simplified
	CardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorTextDim).
			Padding(1, 2).
			Margin(1, 0)
	CardHeaderStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			Margin(0, 0, 1, 0)

	CardContentStyle = lipgloss.NewStyle().
				Foreground(ColorTextPrimary).
				Margin(0, 0, 0, 2)

	// Log styles
	LogContainerStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorTextDim).
				Padding(1).
				Margin(1, 0).
				Background(ColorBgTertiary).
				Height(8)

	LogEntryStyle = lipgloss.NewStyle().
			Foreground(ColorTextSecondary).
			Margin(0, 0, 0, 1)

	LogTimestampStyle = lipgloss.NewStyle().
				Foreground(ColorTextDim).
				Bold(true)

	// Button styles
	ButtonStyle = lipgloss.NewStyle().
			Foreground(ColorBgPrimary).
			Background(ColorSecondary).
			Padding(0, 3).
			Margin(0, 1).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSecondary).
			Bold(true)

	ButtonActiveStyle = lipgloss.NewStyle().
				Foreground(ColorBgPrimary).
				Background(ColorPrimary).
				Padding(0, 3).
				Margin(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorPrimary).
				Bold(true)

	ButtonDisabledStyle = lipgloss.NewStyle().
				Foreground(ColorTextDim).
				Background(ColorBgTertiary).
				Padding(0, 3).
				Margin(0, 1).
				Border(lipgloss.RoundedBorder()).
				BorderForeground(ColorTextDim)

	// Help styles
	HelpStyle = lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Margin(1, 0, 0, 0).
			Padding(1, 2).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorTextDim).
			Background(ColorBgSecondary)

	// Footer style - minimal
	FooterStyle = lipgloss.NewStyle().
			Foreground(ColorTextMuted).
			Padding(1, 0).
			Margin(1, 0, 0, 0).
			Align(lipgloss.Center)
	// Spinner styles
	SpinnerStyle = lipgloss.NewStyle().
			Foreground(ColorAccent).
			Bold(true)

	// Table styles
	TableHeaderStyle = lipgloss.NewStyle().
				Bold(true).
				Foreground(ColorPrimary).
				Background(ColorBgSecondary).
				Padding(0, 1).
				Border(lipgloss.NormalBorder()).
				BorderForeground(ColorSecondary)

	TableCellStyle = lipgloss.NewStyle().
			Foreground(ColorTextPrimary).
			Padding(0, 1).
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorTextDim)

	TableRowEvenStyle = lipgloss.NewStyle().
				Background(ColorBgSecondary)

	TableRowOddStyle = lipgloss.NewStyle().
				Background(ColorBgPrimary)
)

// Utility functions for dynamic styling

// CreateGradientTitle creates a title with gradient effect
func CreateGradientTitle(text string) string {
	return TitleStyle.Render(text)
}

// CreateStatusBadge creates a colored status badge
func CreateStatusBadge(status, text string) string {
	switch status {
	case "success", "completed", "passed":
		return StatusSuccessStyle.Render("✅ " + text)
	case "warning", "pending":
		return StatusWarningStyle.Render("⚠️ " + text)
	case "error", "failed":
		return StatusErrorStyle.Render("❌ " + text)
	case "info", "running":
		return StatusInfoStyle.Render("ℹ️ " + text)
	default:
		return MenuItemStyle.Render(text)
	}
}

// CreateCard creates a styled card container
func CreateCard(title, content string) string {
	header := CardHeaderStyle.Render(title)
	body := CardContentStyle.Render(content)
	return CardStyle.Render(lipgloss.JoinVertical(lipgloss.Left, header, body))
}

// CreateProgressBar creates a styled progress bar
func CreateProgressBar(progress float64, width int) string {
	filled := int(progress * float64(width))
	bar := ""

	for i := 0; i < width; i++ {
		if i < filled {
			bar += "█"
		} else {
			bar += "░"
		}
	}

	return ProgressBarStyle.Render(bar)
}

// CreateButton creates a styled button
func CreateButton(text string, active, disabled bool) string {
	if disabled {
		return ButtonDisabledStyle.Render(text)
	}
	if active {
		return ButtonActiveStyle.Render(text)
	}
	return ButtonStyle.Render(text)
}

// CreateLogEntry creates a styled log entry
func CreateLogEntry(timestamp, message string) string {
	ts := LogTimestampStyle.Render("[" + timestamp + "]")
	msg := LogEntryStyle.Render(message)
	return lipgloss.JoinHorizontal(lipgloss.Left, ts, " ", msg)
}

// CreateTable creates a styled table
func CreateTable(headers []string, rows [][]string) string {
	var table strings.Builder

	// Headers
	headerRow := ""
	for _, header := range headers {
		headerRow += TableHeaderStyle.Render(header)
	}
	table.WriteString(headerRow + "\n")

	// Rows
	for i, row := range rows {
		rowStyle := TableRowEvenStyle
		if i%2 == 1 {
			rowStyle = TableRowOddStyle
		}

		rowStr := ""
		for _, cell := range row {
			cellStyle := TableCellStyle.Copy().Background(rowStyle.GetBackground())
			rowStr += cellStyle.Render(cell)
		}
		table.WriteString(rowStr + "\n")
	}

	return table.String()
}

// CreateBanner creates a simple banner
func CreateBanner(text string) string {
	banner := lipgloss.NewStyle().
		Foreground(ColorPrimary).
		Bold(true).
		Align(lipgloss.Center).
		Padding(0, 2).
		Render(text)

	return banner
}

// CreateSeparator creates a styled separator line
func CreateSeparator(width int) string {
	return lipgloss.NewStyle().
		Foreground(ColorTextDim).
		Render(strings.Repeat("─", width))
}
