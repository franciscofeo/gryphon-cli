package atlassian

import (
	"encoding/json"
	"gryphon/apis/atlassian/response"
	"gryphon/apis/util"
	"io"
	"log"
	"net/http"
)

const url = "https://status.atlassian.com/api/v2/status.json"

func CheckStatus() (*util.StatusInfo, error) {
	var atlassianStatusPageResponse response.AtlassianStatusPageResponse

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

	err = json.Unmarshal(data, &atlassianStatusPageResponse)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &util.StatusInfo{
		Name:        atlassianStatusPageResponse.Page.Name,
		UpdatedAt:   atlassianStatusPageResponse.Page.UpdatedAt,
		Description: atlassianStatusPageResponse.Status.Description,
	}, nil
}
