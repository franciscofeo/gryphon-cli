package slack

import (
	"encoding/json"
	"gryphon/apis/slack/response"
	"gryphon/apis/util"
	"io"
	"log"
	"net/http"
)

const (
	applicationName        = "Slack"
	incidentDefaultMessage = "There is an incident!"
	okDefaultMessage       = "It's everything ok!"
	url                    = "https://status.slack.com/api/v2.0.0/current"
)

func CheckStatus() (*util.StatusInfo, error) {
	var slackStatusPageResponse response.SlackStatusPageResponse

	resp, err := http.Get(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	err = json.Unmarshal(data, &slackStatusPageResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return buildResponse(slackStatusPageResponse), nil
}

func buildResponse(pageResponse response.SlackStatusPageResponse) *util.StatusInfo {
	var description string
	if pageResponse.Status == "active" {
		description = incidentDefaultMessage
	} else {
		description = okDefaultMessage
	}

	return &util.StatusInfo{
		Name:        applicationName,
		UpdatedAt:   pageResponse.DateUpdated + "Z",
		Description: description,
	}
}
