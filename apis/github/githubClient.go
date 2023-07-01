package github

import (
	"encoding/json"
	"gryphon/apis/github/response"
	"gryphon/apis/util"
	"io"
	"log"
	"net/http"
)

const url = "https://www.githubstatus.com/api/v2/status.json"

func CheckStatus() (*util.StatusInfo, error) {
	var githubStatusPageResponse response.GithubStatusPageResponse

	resp, err := http.Get(url)
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

	return &util.StatusInfo{
		Name:        githubStatusPageResponse.Page.Name,
		UpdatedAt:   githubStatusPageResponse.Page.UpdatedAt,
		Description: githubStatusPageResponse.Status.Description,
	}, nil
}
