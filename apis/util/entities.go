package util

type StatusInfo struct {
	Name        string
	UpdatedAt   string
	Description string
}

type WebScrapingApiStatus struct {
	Name   string
	Status string
}

type WebScrapingApiStatusList []WebScrapingApiStatus
