package tui

import (
	"fmt"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

func (m Model) View() tea.View {
	v := tea.NewView("")
	v.AltScreen = true

	if m.width == 0 || m.height == 0 {
		v.SetContent("loading…")
		return v
	}

	if m.showDetail {
		v.SetContent(m.renderDetail())
		return v
	}

	helpBar := m.renderHelpBar()
	helpHeight := lipgloss.Height(helpBar)

	contentH := m.height - helpHeight
	if contentH < 6 {
		contentH = 6
	}
	servicesH := contentH * 40 / 100
	if servicesH < 5 {
		servicesH = 5
	}
	incidentsH := contentH - servicesH

	services := m.renderServicesPane(m.width, servicesH)
	incidents := m.renderIncidentsPane(m.width, incidentsH)

	v.SetContent(lipgloss.JoinVertical(lipgloss.Left, services, incidents, helpBar))
	return v
}

func (m Model) renderHelpBar() string {
	left := m.help.View(m.keys)

	var status string
	switch {
	case m.pendingSummary || m.pendingIncidents:
		status = mutedStyle.Render(m.spinner.View() + " refreshing…")
	case m.lastErr != nil:
		status = errStyle.Render("! " + truncate(m.lastErr.Error(), m.width-len(left)-4))
	case !m.lastRefresh.IsZero():
		status = mutedStyle.Render(fmt.Sprintf("updated %s", m.lastRefresh.Format("15:04:05")))
	default:
		status = mutedStyle.Render("ready")
	}

	gap := m.width - lipgloss.Width(left) - lipgloss.Width(status)
	if gap < 1 {
		gap = 1
	}
	return left + strings.Repeat(" ", gap) + status
}

func (m Model) paneBorder(focused bool) lipgloss.Style {
	if focused {
		return paneFocusedStyle
	}
	return paneStyle
}

func truncate(s string, n int) string {
	if n <= 0 {
		return ""
	}
	if len(s) <= n {
		return s
	}
	if n <= 1 {
		return "…"
	}
	return s[:n-1] + "…"
}

// padLines pads the given text with blank lines up to h lines; if the text has
// more lines than h, it's truncated. Width is clamped.
func padLines(s string, h int) string {
	lines := strings.Split(s, "\n")
	if len(lines) > h {
		lines = lines[:h]
	}
	for len(lines) < h {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}
