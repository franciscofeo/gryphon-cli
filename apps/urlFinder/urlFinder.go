package urlFinder

import (
	"fmt"
	"os"
	"strings"
)

func UrlFinder() (map[string]string, error) {
	lines := make(map[string]string)

	fileData, err := os.ReadFile("./urls.txt")
	if err != nil {
		fmt.Println("- Error to read the file.")
		return nil, err
	}

	text := string(fileData)

	for _, line := range strings.Split(text, "\n") {
		urlInfo := strings.Split(line, ", ")
		lines[urlInfo[0]] = urlInfo[1]
	}

	return lines, nil
}
