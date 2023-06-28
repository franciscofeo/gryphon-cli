package api

import (
	"fmt"
	"log"
	"startup/api/atlassian"
	"startup/api/github"
	"startup/api/slack"
	"startup/api/util"
)

func VerifyAPIs() {
	checkGithubStatus()
	checkSlackStatus()
	checkAtlassianStatus()
	fmt.Println("*******************************************************")
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
		log.Fatal("Error when checking Slack Status Page information.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}

func checkAtlassianStatus() {
	resp, err := atlassian.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Atlassian Status Page information.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}
