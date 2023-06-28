package apps

import (
	"fmt"
	"log"
	"os/exec"
	"startup/urlFinder"
	"time"
)

const (
	defaultBrowser = "xdg-open"
	slack          = "slack"
	intelliJ       = "intellij-idea-ultimate"
)

func OpenApplications() {
	fmt.Println("*******************************************************")
	urls := urlFinder.UrlFinder()
	openUrls(urls)
	openSlack()
	openIntelliJ()
	fmt.Println("*******************************************************")
}

func openUrls(urls map[string]string) {
	for url, name := range urls {
		fmt.Printf("Opening %s in default browser! \n", name)
		err := exec.Command(defaultBrowser, url).Start()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func openSlack() {
	fmt.Println("Opening Slack!")
	err := exec.Command(slack).Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(500 * time.Millisecond)
}

func openIntelliJ() {
	fmt.Println("Opening IntelliJ IDEA!")
	err := exec.Command(intelliJ).Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(650 * time.Millisecond)
}
