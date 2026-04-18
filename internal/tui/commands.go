package tui

import (
	"context"
	"time"

	tea "charm.land/bubbletea/v2"

	"github.com/parham-alvani/arvanstatus-cli/internal/statuspal"
)

const refreshInterval = 60 * time.Second

func tickCmd() tea.Cmd {
	return tea.Tick(refreshInterval, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func refreshCmd(c *statuspal.Client) tea.Cmd {
	started := func() tea.Msg { return refreshStartedMsg{} }
	return tea.Batch(started, fetchSummaryCmd(c), fetchIncidentsCmd(c))
}

func fetchSummaryCmd(c *statuspal.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		s, err := c.FetchSummary(ctx)
		return summaryFetchedMsg{s: s, err: err}
	}
}

func fetchIncidentsCmd(c *statuspal.Client) tea.Cmd {
	return func() tea.Msg {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		p, err := c.FetchIncidents(ctx)
		return incidentsFetchedMsg{p: p, err: err}
	}
}
