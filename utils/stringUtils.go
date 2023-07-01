package utils

import "strings"

func IsStringBlank(s string) bool {
	if len(strings.TrimSpace(s)) == 0 {
		return true
	} else {
		return false
	}
}
