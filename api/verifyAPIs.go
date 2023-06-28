package api

import (
	"fmt"
	"log"
	"startup/api/github"
	"startup/api/slack"
	"startup/api/util"
)

func VerifyAPIs() {
	fmt.Println("***************************************")
	checkGithubStatus()
	checkSlackStatus()
	fmt.Println("***************************************")
}

func checkGithubStatus() {
	resp, err := github.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Github Status Page informations.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}

func checkSlackStatus() {
	resp, err := slack.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Slack Status Page informations.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}
