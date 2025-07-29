package tui

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// NotificationType represents different types of notifications
type NotificationType int

const (
	NotificationInfo NotificationType = iota
	NotificationSuccess
	NotificationWarning
	NotificationError
)

// Notification represents a single notification
type Notification struct {
	ID       string
	Type     NotificationType
	Title    string
	Message  string
	Duration time.Duration
	Created  time.Time
}

// NotificationManager manages the notification system
type NotificationManager struct {
	notifications   []Notification
	maxVisible      int
	defaultDuration time.Duration
}

// NewNotificationManager creates a new notification manager
func NewNotificationManager() *NotificationManager {
	return &NotificationManager{
		notifications:   make([]Notification, 0),
		maxVisible:      3,
		defaultDuration: 3 * time.Second,
	}
}

// AddNotification adds a new notification
func (nm *NotificationManager) AddNotification(notifType NotificationType, title, message string) tea.Cmd {
	notification := Notification{
		ID:       fmt.Sprintf("notif_%d", time.Now().UnixNano()),
		Type:     notifType,
		Title:    title,
		Message:  message,
		Duration: nm.defaultDuration,
		Created:  time.Now(),
	}

	nm.notifications = append(nm.notifications, notification)

	// Keep only the most recent notifications
	if len(nm.notifications) > nm.maxVisible {
		nm.notifications = nm.notifications[len(nm.notifications)-nm.maxVisible:]
	}

	// Return a command to remove the notification after its duration
	return nm.removeNotificationAfter(notification.ID, notification.Duration)
}

// removeNotificationAfter creates a command to remove a notification after a delay
func (nm *NotificationManager) removeNotificationAfter(id string, duration time.Duration) tea.Cmd {
	return tea.Tick(duration, func(t time.Time) tea.Msg {
		return NotificationExpiredMsg{ID: id}
	})
}

// RemoveNotification removes a notification by ID
func (nm *NotificationManager) RemoveNotification(id string) {
	for i, notif := range nm.notifications {
		if notif.ID == id {
			nm.notifications = append(nm.notifications[:i], nm.notifications[i+1:]...)
			break
		}
	}
}

// GetVisibleNotifications returns currently visible notifications
func (nm *NotificationManager) GetVisibleNotifications() []Notification {
	now := time.Now()
	visible := make([]Notification, 0)

	for _, notif := range nm.notifications {
		if now.Sub(notif.Created) < notif.Duration {
			visible = append(visible, notif)
		}
	}

	return visible
}

// NotificationExpiredMsg is sent when a notification expires
type NotificationExpiredMsg struct {
	ID string
}

// RenderNotifications renders all visible notifications
func (nm *NotificationManager) RenderNotifications() string {
	notifications := nm.GetVisibleNotifications()
	if len(notifications) == 0 {
		return ""
	}

	var rendered []string
	for _, notif := range notifications {
		rendered = append(rendered, nm.renderSingleNotification(notif))
	}

	// Position notifications in top-right corner
	container := lipgloss.NewStyle().
		Align(lipgloss.Right).
		Width(50).
		Render(lipgloss.JoinVertical(lipgloss.Right, rendered...))

	return container
}

// renderSingleNotification renders a single notification
func (nm *NotificationManager) renderSingleNotification(notif Notification) string {
	var style lipgloss.Style
	var icon string

	switch notif.Type {
	case NotificationSuccess:
		style = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorSuccess).
			Padding(0, 1).
			Margin(0, 0, 1, 0)
		icon = "✅"
	case NotificationWarning:
		style = lipgloss.NewStyle().
			Foreground(ColorWarning).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorWarning).
			Padding(0, 1).
			Margin(0, 0, 1, 0)
		icon = "⚠️"
	case NotificationError:
		style = lipgloss.NewStyle().
			Foreground(ColorError).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorError).
			Padding(0, 1).
			Margin(0, 0, 1, 0)
		icon = "❌"
	default: // NotificationInfo
		style = lipgloss.NewStyle().
			Foreground(ColorInfo).
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorInfo).
			Padding(0, 1).
			Margin(0, 0, 1, 0)
		icon = "ℹ️"
	}

	content := fmt.Sprintf("%s %s", icon, notif.Title)
	if notif.Message != "" {
		content += fmt.Sprintf("\n%s", notif.Message)
	}

	return style.Render(content)
}

// NotifiableModel is an interface for models that can show notifications
type NotifiableModel interface {
	tea.Model
	GetNotificationManager() *NotificationManager
	HandleNotificationMsg(msg tea.Msg) (tea.Model, tea.Cmd)
}

// WithNotifications wraps a model with notification capabilities
type WithNotifications struct {
	Model               tea.Model
	NotificationManager *NotificationManager
}

// NewWithNotifications creates a new model with notifications
func NewWithNotifications(model tea.Model) WithNotifications {
	return WithNotifications{
		Model:               model,
		NotificationManager: NewNotificationManager(),
	}
}

// Init initializes the wrapped model
func (wn WithNotifications) Init() tea.Cmd {
	return wn.Model.Init()
}

// Update handles messages for the wrapped model and notifications
func (wn WithNotifications) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	// Handle notification expiration
	if expiredMsg, ok := msg.(NotificationExpiredMsg); ok {
		wn.NotificationManager.RemoveNotification(expiredMsg.ID)
		return wn, nil
	}

	// Update the wrapped model
	wn.Model, cmd = wn.Model.Update(msg)
	return wn, cmd
}

// View renders the wrapped model with notifications overlay
func (wn WithNotifications) View() string {
	baseView := wn.Model.View()
	notifications := wn.NotificationManager.RenderNotifications()

	if notifications == "" {
		return baseView
	}

	// Overlay notifications on top of the base view
	return lipgloss.JoinVertical(lipgloss.Left, notifications, baseView)
}

// ShowNotification is a helper to show notifications
func (wn WithNotifications) ShowNotification(notifType NotificationType, title, message string) tea.Cmd {
	return wn.NotificationManager.AddNotification(notifType, title, message)
}

// Convenience methods for different notification types
func (wn WithNotifications) ShowSuccess(title, message string) tea.Cmd {
	return wn.ShowNotification(NotificationSuccess, title, message)
}

func (wn WithNotifications) ShowWarning(title, message string) tea.Cmd {
	return wn.ShowNotification(NotificationWarning, title, message)
}

func (wn WithNotifications) ShowError(title, message string) tea.Cmd {
	return wn.ShowNotification(NotificationError, title, message)
}

func (wn WithNotifications) ShowInfo(title, message string) tea.Cmd {
	return wn.ShowNotification(NotificationInfo, title, message)
}
