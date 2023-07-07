package btgpactual

import (
	"github.com/gocolly/colly"
	"gryphon/apis/util"
	"log"
	"time"
)

const (
	url = "https://btgpactualempresas.statuspage.io/"
)

type ApiStatus struct {
	Name   string
	Status string
}

type ApiStatusList []ApiStatus

func CheckStatus() []*util.StatusInfo {
	var apiStatusList ApiStatusList

	c := colly.NewCollector()

	c.OnHTML("div.child-components-container div.component-inner-container", func(e *colly.HTMLElement) {
		name := e.ChildText(".name")
		status := e.ChildText(".component-status")
		apiInfo := ApiStatus{
			Name:   name,
			Status: status,
		}
		apiStatusList = append(apiStatusList, apiInfo)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when verifying BTG Pactual Status Page: ", err)
	})

	_ = c.Visit(url)

	return convertToStatusInfoStruct(apiStatusList)
}

func convertToStatusInfoStruct(apiStatusList ApiStatusList) []*util.StatusInfo {
	var statusInfoList []*util.StatusInfo
	for _, api := range apiStatusList {
		statusInfoList = append(statusInfoList, &util.StatusInfo{
			Name:        api.Name,
			UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			Description: api.Status,
		})
	}

	return statusInfoList
}