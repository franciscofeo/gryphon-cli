package api

import (
	"fmt"
	"log"
	"startup/api/github"
)

func VerifyAPIs() {
	fmt.Println("***************************************")
	checkGithubStatus()
	fmt.Println("***************************************")
}

func checkGithubStatus() {
	resp, err := github.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Github Status Page informations.")
		return
	}

	fmt.Printf("Checking %s APIs Status \n", resp.Name)
	fmt.Printf("Status: %s \n", resp.Description)
	fmt.Printf("Last Update: %s \n", resp.UpdatedAt)
	fmt.Println("-------------------")
}
