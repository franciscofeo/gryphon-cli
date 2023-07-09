package btgpactual

import (
	"github.com/gocolly/colly"
	"gryphon/apis/util"
	"log"
)

const url = "https://btgpactualempresas.statuspage.io/"

func CheckStatus() []util.StatusInfo {
	var apiStatusList util.WebScrapingApiStatusList

	c := colly.NewCollector()

	c.OnHTML("div.child-components-container div.component-inner-container", func(e *colly.HTMLElement) {
		name := e.ChildText(".name")
		status := e.ChildText(".component-status")
		apiInfo := util.WebScrapingApiStatus{
			Name:   name,
			Status: status,
		}
		apiStatusList = append(apiStatusList, apiInfo)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong when verifying BTG Pactual Status Page: ", err)
	})

	_ = c.Visit(url)

	return util.ConvertToStatusInfoStruct(apiStatusList)
}
