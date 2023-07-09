package azure

import (
	"fmt"
	"github.com/gocolly/colly"
	"gryphon/apis/util"
	"gryphon/utils"
)

const url = "https://azure.status.microsoft/en-us/status"

func getServices() map[string]bool {
	return map[string]bool{
		"service-bus":                   true,
		"event-hubs":                    true,
		"azure-cosmos-db":               true,
		"azure-database-for-postgresql": true,
	}
}

func CheckStatus() []util.StatusInfo {
	var list util.WebScrapingApiStatusList
	services := getServices()

	c := colly.NewCollector()

	c.OnHTML("table[data-zone-name='americas'] tr", func(e *colly.HTMLElement) {
		htmlServiceName := e.ChildText("td:nth-child(1)")
		serviceName := utils.ToLowerCaseAndRemoveWhiteSpaces(htmlServiceName)

		exists, shouldProcess := services[serviceName]
		if exists && shouldProcess {
			status := e.ChildText("td:nth-child(4)")
			fmt.Println("Existe o serviço: ", serviceName)
			apiInfo := util.WebScrapingApiStatus{
				Name:   serviceName,
				Status: status,
			}
			list = append(list, apiInfo)
		}

	})

	_ = c.Visit(url)

	return util.ConvertToStatusInfoStruct(list)
}
