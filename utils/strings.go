package utils

import (
	"strings"
)

// StringInListCaseInsensitive return true if str is in the list (case insensitive)
func StringInListCaseInsensitive(list []string, str string) bool {
	for _, s := range list {
		if strings.ToLower(s) == strings.ToLower(str) {
			return true
		}
	}
	return false
}
