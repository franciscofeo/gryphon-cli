package response

type SlackStatusPageResponse struct {
	Status          string            `json:"status"`
	DateCreated     string            `json:"date_created"`
	DateUpdated     string            `json:"date_updated"`
	ActiveIncidents []activeIncidents `json:"active_incidents"`
}

type activeIncidents struct {
	ID          int32    `json:"id"`
	DateCreated string   `json:"date_created"`
	DateUpdated string   `json:"date_updated"`
	Title       string   `json:"title"`
	Type        string   `json:"type"`
	Status      string   `json:"status"`
	URL         string   `json:"url"`
	Services    []string `json:"services"`
	Notes       []notes  `json:"notes"`
}

type notes struct {
	DateCreated string `json:"date_created"`
	Body        string `json:"body"`
}
