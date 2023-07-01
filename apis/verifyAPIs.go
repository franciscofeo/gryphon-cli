package apis

import (
	"fmt"
	"gryphon/apis/atlassian"
	"gryphon/apis/github"
	"gryphon/apis/slack"
	"gryphon/apis/util"
	"log"
)

func VerifyAPIs() {
	fmt.Println("Verifying APIs:")
	fmt.Println("---------")
	checkGithubStatus()
	checkSlackStatus()
	checkAtlassianStatus()
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
