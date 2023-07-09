package utils

import "strings"

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
