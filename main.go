package main

import (
	"fmt"
	"log"
	"os/exec"
	"startup/urlFinder"
	"time"
)

func main() {
	fmt.Println("Good Morning!")

	urls := urlFinder.UrlFinder()
	openUrls(urls)

}

func openUrls(urls map[string]string) {
	for url, name := range urls {
		fmt.Printf("Opening %s in the default browser \n", name)
		err := exec.Command("xdg-open", url).Start()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
