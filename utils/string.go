package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

func IsStringBlank(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	} else {
		return false
	}
}

func ToLowerCaseAndRemoveWhiteSpaces(s string) string {
	lowerCaseString := strings.ToLower(s)
	return strings.ReplaceAll(lowerCaseString, " ", "-")
}

func ToTitleAndRemoveHyphen(s string) string {
	wordsList := strings.Split(s, "-")

	for i := 0; i < len(wordsList); i++ {
		wordsList[i] = cases.Title(language.English).String(wordsList[i])
	}

	return strings.Join(wordsList, " ")
}
