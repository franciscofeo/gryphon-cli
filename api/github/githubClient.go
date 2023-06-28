package github

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"startup/api"
	"startup/api/github/response"
)

func CheckStatus() (*api.StatusInfo, error) {
	var githubStatusPageResponse response.GithubStatusPageResponse

	resp, err := http.Get("https://www.githubstatus.com/api/v2/status.json")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = json.Unmarshal(data, &githubStatusPageResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &api.StatusInfo{
		Name:        githubStatusPageResponse.Page.Name,
		UpdatedAt:   githubStatusPageResponse.Page.UpdatedAt,
		Description: githubStatusPageResponse.Status.Description,
	}, nil
}
