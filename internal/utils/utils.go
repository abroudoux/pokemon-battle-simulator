package utils

import (
	"strings"
	"unicode"
)

func CheckParamType(s string) string {
	hasDigit := false
	hasLetter := false

	for _, r := range s {
		if unicode.IsDigit(r) {
			hasDigit = true
		}
		if unicode.IsLetter(r) {
			hasLetter = true
		}
	}

	if hasDigit && hasLetter {
		return "mixed"
	} else if hasDigit {
		return "digit"
	} else if hasLetter {
		return "letter"
	} else {
		return "unknown"
	}
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func CapitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}