package tui

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"

	"github.com/parham-alvani/arvanstatus-cli/internal/statuspal"
)

func (m Model) renderServicesPane(width, height int) string {
	border := m.paneBorder(m.focused == paneServices)
	innerW := width - 2 // border
	innerH := height - 2
	if innerW < 4 || innerH < 2 {
		return border.Width(width - 2).Height(height - 2).Render("")
	}

	title := titleStyle.Render(m.servicesTitle())
	body := m.renderServicesBody(innerW, innerH-1)

	content := lipgloss.JoinVertical(lipgloss.Left, title, body)
	return border.Width(innerW).Height(innerH).Render(content)
}

func (m Model) servicesTitle() string {
	name := "ArvanCloud"
	if m.summary != nil && m.summary.StatusPage.Name != "" {
		name = m.summary.StatusPage.Name
	}
	badge := "all systems operational"
	badgeStyle := lipgloss.NewStyle().Foreground(colOperational)
	if m.summary != nil {
		if deg := countDegraded(m.summary.Services); deg > 0 {
			badge = fmt.Sprintf("%d service(s) affected", deg)
			badgeStyle = lipgloss.NewStyle().Foreground(colMajor)
		}
	}
	return fmt.Sprintf("%s  %s", name, badgeStyle.Render(badge))
}

func countDegraded(xs []statuspal.Service) int {
	n := 0
	for _, x := range xs {
		if x.CurrentIncidentType != nil {
			n++
		}
		n += countDegraded(x.Children)
	}
	return n
}

func (m Model) renderServicesBody(width, height int) string {
	if len(m.flatServices) == 0 {
		if m.pendingSummary {
			return mutedStyle.Render("  " + m.spinner.View() + " loading services…")
		}
		return mutedStyle.Render("  no data")
	}

	cursor := m.servicesCursor
	start := 0
	if cursor >= height {
		start = cursor - height + 1
	}
	end := start + height
	if end > len(m.flatServices) {
		end = len(m.flatServices)
	}

	var lines []string
	for i := start; i < end; i++ {
		fs := m.flatServices[i]
		indent := strings.Repeat("  ", fs.depth)
		bullet := lipgloss.NewStyle().
			Foreground(severityColor(fs.svc.CurrentIncidentType)).
			Render("●")

		name := fs.svc.Name
		label := fmt.Sprintf("%s%s %s", indent, bullet, name)

		if i == cursor && m.focused == paneServices {
			label = selectedStyle.Render("▸ ") + label
		} else {
			label = "  " + label
		}

		label = lipgloss.NewStyle().MaxWidth(width).Render(label)
		lines = append(lines, label)
	}
	return padLines(strings.Join(lines, "\n"), height)
}
