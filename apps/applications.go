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
	defaultBrowser   = "xdg-open"
	slack            = "slack"
	intelliJUltimate = "intellij-idea-ultimate"
	intelliJV1       = "intellij"
	intelliJV2       = "idea"
)

var appHandler = map[string]func(){
	slack:            openSlack,
	"browser":        openUrls,
	intelliJUltimate: openIntelliJ,
}

func OpenApplications() {
	fmt.Println("Opening applications:")
	openUrls()
	openSlack()
	openIntelliJ()
}

func ListAvailableApplications() {
	fmt.Println("List of available applications:")
	for name := range appHandler {
		fmt.Println("- " + name)
	}

	fmt.Println("\nList of available URLs:")
	urlList, err := urlFinder.UrlFinder()
	if err != nil {
		return
	}

	for _, name := range urlList {
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
	urls, err := urlFinder.UrlFinder()
	if err != nil {
		fmt.Println("There is no urls.txt file to open urls in the browser.")
		return
	}

	for url, name := range urls {
		fmt.Printf("-"+" Opening %s in default defaultBrowser! \n", name)
		err := exec.Command(defaultBrowser, url).Start()
		if err != nil {
			log.Println(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func openSlack() {
	fmt.Println("-" + " Opening Slack!")
	err := exec.Command(slack).Start()
	if err != nil {
		log.Println(err)
	}
	time.Sleep(500 * time.Millisecond)
}

func openIntelliJ() {
	fmt.Println("-" + " Opening IntelliJ IDEA!")

	err := exec.Command(intelliJUltimate).Start()
	if err != nil {
		log.Println(err)
	}

	err = exec.Command(intelliJV1).Start()
	if err != nil {
		log.Println(err)
	}

	err = exec.Command(intelliJV2).Start()
	if err != nil {
		log.Println(err)
	}

	time.Sleep(650 * time.Millisecond)
}
