// Package tui is the Bubble Tea view layer for arvanstatus-cli.
package tui

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/spinner"
	"charm.land/bubbles/v2/viewport"

	"github.com/parham-alvani/arvanstatus-cli/internal/statuspal"
)

type pane int

const (
	paneServices pane = iota
	paneIncidents
)

// flatService is a DFS-flattened entry of the services tree with a depth hint
// for indented rendering.
type flatService struct {
	depth int
	svc   statuspal.Service
}

// Model holds all TUI state.
type Model struct {
	client *statuspal.Client

	// data
	summary      *statuspal.Summary
	incidents    []statuspal.Incident
	serviceByID  map[int]statuspal.Service
	flatServices []flatService

	// ui
	focused          pane
	servicesCursor   int
	incidentsCursor  int
	detail           viewport.Model
	spinner          spinner.Model
	help             help.Model
	keys             keyMap
	showHelp         bool

	// mode
	showDetail bool

	// chrome
	width, height    int
	lastRefresh      time.Time
	pendingSummary   bool
	pendingIncidents bool
	lastErr          error
}

// New constructs a Model with a default http client pointed at the arvancloud
// status page.
func New(client *statuspal.Client) Model {
	sp := spinner.New()
	sp.Spinner = spinner.Dot

	h := help.New()
	h.Styles = help.DefaultDarkStyles()

	vp := viewport.New()

	return Model{
		client:  client,
		focused: paneIncidents,
		spinner: sp,
		help:    h,
		keys:    defaultKeys(),
		detail:  vp,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(m.spinner.Tick, refreshCmd(m.client), tickCmd())
}

// flattenServices performs a DFS over the services tree and returns a slice of
// (depth, service) entries, good for one-line-per-service rendering.
func flattenServices(xs []statuspal.Service) []flatService {
	var out []flatService
	var walk func(int, []statuspal.Service)
	walk = func(depth int, xs []statuspal.Service) {
		for _, x := range xs {
			out = append(out, flatService{depth: depth, svc: x})
			walk(depth+1, x.Children)
		}
	}
	walk(0, xs)
	return out
}
