package util

import (
	"fmt"
	"time"
)

func GenerateAPIsInformationMessage(resp *StatusInfo) {
	fmt.Printf("Checking %s APIs Status \n", resp.Name)
	fmt.Printf("Status: %s \n", resp.Description)
	fmt.Printf("Last Update: %s \n", resp.UpdatedAt)
	fmt.Println("-------------------")
}

func GenerateBtgAPIsInformationMessage(resp []*StatusInfo) {
	fmt.Println("Checking BTG APIs: ")
	fmt.Println()
	for _, statusInfo := range resp {
		fmt.Printf("Checking %s APIs Status \n", statusInfo.Name)
		fmt.Printf("Status: %s \n", statusInfo.Description)
		fmt.Printf("Last Update: %s \n", statusInfo.UpdatedAt)
		fmt.Println()
	}
	fmt.Println("-------------------")
}

func ConvertToStatusInfoStruct(apiStatusList WebScrapingApiStatusList) []StatusInfo {
	var statusInfoList []StatusInfo
	for _, api := range apiStatusList {
		statusInfoList = append(statusInfoList, StatusInfo{
			Name:        api.Name,
			UpdatedAt:   time.Now().UTC().Format("2006-01-02T15:04:05.000Z"),
			Description: api.Status,
		})
	}

	return statusInfoList
}
