package statuspal

import "context"

// Incident types and statuses used in the StatusPal API.
const (
	TypeMinor     = "minor"
	TypeMajor     = "major"
	TypeScheduled = "scheduled"

	StatusInvestigating = "investigating"
	StatusIdentified    = "identified"
	StatusMonitoring    = "monitoring"
	StatusResolved      = "resolved"
)

// Incident is one row in the /incidents list.
type Incident struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Type           string    `json:"type"`
	Status         string    `json:"status"`
	StartedAt      LocalTime `json:"started_at"`
	UpdatedAt      LocalTime `json:"updated_at"`
	AffectsUptime  bool      `json:"affects_uptime"`
	Services       []int     `json:"services"`
	LatestUpdate   string    `json:"latest_update"`
}

type PageInfo struct {
	HasNext bool `json:"has_next"`
	HasPrev bool `json:"has_prev"`
}

type IncidentsMetadata struct {
	TotalCount int      `json:"total_count"`
	PageInfo   PageInfo `json:"page_info"`
}

// IncidentsPage is the response of GET /status_pages/{page}/incidents.
type IncidentsPage struct {
	Incidents []Incident        `json:"incidents"`
	Metadata  IncidentsMetadata `json:"metadata"`
}

func (c *Client) FetchIncidents(ctx context.Context) (*IncidentsPage, error) {
	var p IncidentsPage
	if err := c.getJSON(ctx, "incidents", &p); err != nil {
		return nil, err
	}
	return &p, nil
}
