package statuspal

import "context"

// StatusPage is the top-level status page metadata.
type StatusPage struct {
	Name     string `json:"name"`
	URL      string `json:"url"`
	TimeZone string `json:"time_zone"`
}

// Service is one node in the services tree. Children recurse.
type Service struct {
	ID                  int       `json:"id"`
	Name                string    `json:"name"`
	Description         string    `json:"description,omitempty"`
	ParentID            *int      `json:"parent_id"`
	CurrentIncidentType *string   `json:"current_incident_type"`
	Children            []Service `json:"children"`
}

// Summary is the response of GET /status_pages/{page}/summary.
type Summary struct {
	StatusPage           StatusPage `json:"status_page"`
	Services             []Service  `json:"services"`
	OngoingIncidents     []Incident `json:"ongoing_incidents"`
	UpcomingMaintenances []Incident `json:"upcoming_maintenances"`
}

// ServicesByID walks the tree once and returns a map keyed by service ID.
func (s *Summary) ServicesByID() map[int]Service {
	out := map[int]Service{}
	var walk func([]Service)
	walk = func(xs []Service) {
		for _, x := range xs {
			out[x.ID] = x
			walk(x.Children)
		}
	}
	walk(s.Services)
	return out
}

func (c *Client) FetchSummary(ctx context.Context) (*Summary, error) {
	var s Summary
	if err := c.getJSON(ctx, "summary", &s); err != nil {
		return nil, err
	}
	return &s, nil
}
