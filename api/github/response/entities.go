package response

type GithubStatusPageResponse struct {
	Page   Page   `json:"page"`
	Status Status `json:"status"`
}

type Page struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	Timezone  string `json:"time_zone"`
	UpdatedAt string `json:"updated_at"`
}

type Status struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}
