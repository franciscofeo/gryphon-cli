package response

type AtlassianStatusPageResponse struct {
	Page   page   `json:"page"`
	Status status `json:"status"`
}

type page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	URL       string `json:"url"`
	TimeZone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
