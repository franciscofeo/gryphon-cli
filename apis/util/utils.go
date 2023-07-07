package util

import (
	"fmt"
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
