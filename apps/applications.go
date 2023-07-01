package apps

import (
	"fmt"
	"gryphon/apps/urlFinder"
	"log"
	"os/exec"
	"strings"
	"time"
)

const (
	defaultBrowser = "xdg-open"
	slack          = "slack"
	intelliJ       = "intellij-idea-ultimate"
)

var appHandler = map[string]func(){
	slack:     openSlack,
	"browser": openUrls,
	intelliJ:  openIntelliJ,
}

func OpenApplications() {
	fmt.Println("Opening applications:")
	openUrls()
	openSlack()
	openIntelliJ()
}

func ListApplications() {
	fmt.Println("List of available URLs:")
	urlList := urlFinder.UrlFinder()
	for _, name := range urlList {
		fmt.Println("- " + name)
	}

	fmt.Println("List of available applications:")
	for name := range appHandler {
		fmt.Println("- " + name)
	}
}

func OpenSingleApplication(name string) error {
	nameFormatted := strings.ToLower(name)
	handler, ok := appHandler[nameFormatted]
	if ok {
		handler()
		return nil
	} else {
		return fmt.Errorf("application not found")
	}
}

func openUrls() {
	urls := urlFinder.UrlFinder()

	for url, name := range urls {
		fmt.Printf("-"+" Opening %s in default defaultBrowser! \n", name)
		err := exec.Command(defaultBrowser, url).Start()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func openSlack() {
	fmt.Println("-" + " Opening Slack!")
	err := exec.Command(slack).Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(500 * time.Millisecond)
}

func openIntelliJ() {
	fmt.Println("-" + " Opening IntelliJ IDEA!")
	err := exec.Command(intelliJ).Start()
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(650 * time.Millisecond)
}
