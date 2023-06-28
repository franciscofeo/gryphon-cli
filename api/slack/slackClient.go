package slack

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"startup/api"
	"startup/api/slack/response"
)

const (
	applicationName        = "Slack"
	incidentDefaultMessage = "There is an incident!"
	okDefaultMessage       = "It's everything ok!"
)

func CheckStatus() (*api.StatusInfo, error) {
	var slackStatusPageResponse response.SlackStatusPageResponse

	resp, err := http.Get("https://status.slack.com/api/v2.0.0/current")
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

	err = json.Unmarshal(data, &slackStatusPageResponse)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return buildResponse(slackStatusPageResponse), nil
}

func buildResponse(pageResponse response.SlackStatusPageResponse) *api.StatusInfo {
	var description string
	if pageResponse.Status == "active" {
		description = incidentDefaultMessage
	} else {
		description = okDefaultMessage
	}

	return &api.StatusInfo{
		Name:        applicationName,
		UpdatedAt:   pageResponse.DateUpdated,
		Description: description,
	}
}
