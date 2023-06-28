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
