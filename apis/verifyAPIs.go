package apis

import (
	"fmt"
	"gryphon/apis/atlassian"
	"gryphon/apis/github"
	"gryphon/apis/slack"
	"gryphon/apis/util"
	"log"
	"sync"
)

func VerifyAPIs() {
	var wg sync.WaitGroup
	wg.Add(3)

	fmt.Println("Verifying APIs:")
	fmt.Println("---------")
	go checkGithubStatus(&wg)
	go checkSlackStatus(&wg)
	go checkAtlassianStatus(&wg)

	wg.Wait()
}

func checkGithubStatus(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := github.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Github Status Page informations.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}

func checkSlackStatus(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := slack.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Slack Status Page information.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}

func checkAtlassianStatus(wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := atlassian.CheckStatus()
	if err != nil {
		log.Fatal("Error when checking Atlassian Status Page information.")
		return
	}
	util.GenerateAPIsInformationMessage(resp)
}
