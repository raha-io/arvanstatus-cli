package tui

import (
	"image/color"

	"charm.land/lipgloss/v2"
)

var (
	colOperational = lipgloss.Color("#22c55e")
	colMinor       = lipgloss.Color("#eab308")
	colMajor       = lipgloss.Color("#ef4444")
	colScheduled   = lipgloss.Color("#3b82f6")
	colIdentified  = lipgloss.Color("#f97316")
	colMonitoring  = lipgloss.Color("#eab308")
	colMuted       = lipgloss.Color("#71717a")
	colText        = lipgloss.Color("#e4e4e7")
	colAccent      = lipgloss.Color("#d946ef")
	colBorderIdle  = lipgloss.Color("#3f3f46")
)

var (
	titleStyle = lipgloss.NewStyle().
			Foreground(colAccent).
			Bold(true).
			Padding(0, 1)

	mutedStyle = lipgloss.NewStyle().Foreground(colMuted)

	textStyle = lipgloss.NewStyle().Foreground(colText)

	selectedStyle = lipgloss.NewStyle().
			Foreground(colAccent).
			Bold(true)

	errStyle = lipgloss.NewStyle().
			Foreground(colMajor).
			Bold(true)

	paneStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colBorderIdle)

	paneFocusedStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(colAccent)

	detailStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(colAccent).
			Padding(0, 1)
)

// severityColor maps a *current_incident_type to its bullet color.
// Nil (operational) → green.
func severityColor(t *string) color.Color {
	if t == nil {
		return colOperational
	}
	switch *t {
	case "minor":
		return colMinor
	case "major":
		return colMajor
	case "scheduled":
		return colScheduled
	}
	return colMuted
}

// incidentTypeColor returns the accent color for an incident.Type.
func incidentTypeColor(t string) color.Color {
	switch t {
	case "minor":
		return colMinor
	case "major":
		return colMajor
	case "scheduled":
		return colScheduled
	}
	return colMuted
}

// incidentStatusColor returns the accent color for an incident.Status.
func incidentStatusColor(s string) color.Color {
	switch s {
	case "investigating":
		return colMajor
	case "identified":
		return colIdentified
	case "monitoring":
		return colMonitoring
	case "resolved":
		return colOperational
	}
	return colMuted
}
