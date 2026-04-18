package tui

import (
	"time"

	tea "charm.land/bubbletea/v2"

	"charm.land/bubbles/v2/key"
	"charm.land/bubbles/v2/spinner"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width, m.height = msg.Width, msg.Height
		m.resizeDetail()
		return m, nil

	case tea.KeyPressMsg:
		return m.handleKey(msg)

	case tickMsg:
		return m, tea.Batch(refreshCmd(m.client), tickCmd())

	case refreshStartedMsg:
		m.pendingSummary = true
		m.pendingIncidents = true
		return m, nil

	case summaryFetchedMsg:
		m.pendingSummary = false
		if msg.err != nil {
			m.lastErr = msg.err
		} else {
			m.summary = msg.s
			m.serviceByID = msg.s.ServicesByID()
			m.flatServices = flattenServices(msg.s.Services)
			m.clampCursors()
			m.lastErr = nil
		}
		m.maybeMarkRefreshed()
		return m, nil

	case incidentsFetchedMsg:
		m.pendingIncidents = false
		if msg.err != nil {
			m.lastErr = msg.err
		} else {
			m.incidents = msg.p.Incidents
			m.clampCursors()
			m.lastErr = nil
		}
		m.maybeMarkRefreshed()
		return m, nil

	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	if m.showDetail {
		var cmd tea.Cmd
		m.detail, cmd = m.detail.Update(msg)
		return m, cmd
	}
	return m, nil
}

func (m Model) handleKey(msg tea.KeyPressMsg) (tea.Model, tea.Cmd) {
	switch {
	case key.Matches(msg, m.keys.Quit):
		return m, tea.Quit

	case key.Matches(msg, m.keys.Help):
		m.showHelp = !m.showHelp
		m.help.ShowAll = m.showHelp
		return m, nil

	case key.Matches(msg, m.keys.Refresh):
		return m, refreshCmd(m.client)

	case key.Matches(msg, m.keys.Esc):
		if m.showDetail {
			m.showDetail = false
		}
		return m, nil

	case key.Matches(msg, m.keys.Enter):
		if m.focused == paneIncidents && len(m.incidents) > 0 && !m.showDetail {
			m.showDetail = true
			m.detail.SetContent(renderDetailBody(m.incidents[m.incidentsCursor], m.serviceByID, m.detail.Width()))
			m.detail.SetYOffset(0)
		}
		return m, nil

	case key.Matches(msg, m.keys.Tab):
		if m.showDetail {
			return m, nil
		}
		if m.focused == paneServices {
			m.focused = paneIncidents
		} else {
			m.focused = paneServices
		}
		return m, nil

	case key.Matches(msg, m.keys.Up):
		if m.showDetail {
			m.detail.ScrollUp(1)
			return m, nil
		}
		m.moveCursor(-1)
		return m, nil

	case key.Matches(msg, m.keys.Down):
		if m.showDetail {
			m.detail.ScrollDown(1)
			return m, nil
		}
		m.moveCursor(+1)
		return m, nil
	}
	return m, nil
}

func (m *Model) moveCursor(delta int) {
	switch m.focused {
	case paneServices:
		n := len(m.flatServices)
		if n == 0 {
			m.servicesCursor = 0
			return
		}
		m.servicesCursor = (m.servicesCursor + delta + n) % n
	case paneIncidents:
		n := len(m.incidents)
		if n == 0 {
			m.incidentsCursor = 0
			return
		}
		m.incidentsCursor = (m.incidentsCursor + delta + n) % n
	}
}

func (m *Model) clampCursors() {
	if m.servicesCursor >= len(m.flatServices) {
		m.servicesCursor = max0(len(m.flatServices) - 1)
	}
	if m.incidentsCursor >= len(m.incidents) {
		m.incidentsCursor = max0(len(m.incidents) - 1)
	}
}

func (m *Model) maybeMarkRefreshed() {
	if !m.pendingSummary && !m.pendingIncidents {
		m.lastRefresh = time.Now()
	}
}

func (m *Model) resizeDetail() {
	if m.width <= 4 || m.height <= 4 {
		return
	}
	m.detail.SetWidth(m.width - 4)
	m.detail.SetHeight(m.height - 4)
}

func max0(n int) int {
	if n < 0 {
		return 0
	}
	return n
}
