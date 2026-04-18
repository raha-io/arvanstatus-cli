package tui

import (
	"time"

	"github.com/parham-alvani/arvanstatus-cli/internal/statuspal"
)

type tickMsg time.Time

type refreshStartedMsg struct{}

type summaryFetchedMsg struct {
	s   *statuspal.Summary
	err error
}

type incidentsFetchedMsg struct {
	p   *statuspal.IncidentsPage
	err error
}
