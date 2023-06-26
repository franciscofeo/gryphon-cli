package urlFinder

import (
	"log"
	"os"
	"strings"
)

func UrlFinder() map[string]string {
	lines := make(map[string]string)

	fileData, err := os.ReadFile("./urlFinder/urls.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(fileData)

	for _, line := range strings.Split(text, "\n") {
		urlInfo := strings.Split(line, ", ")
		lines[urlInfo[0]] = urlInfo[1]
	}

	return lines
}
