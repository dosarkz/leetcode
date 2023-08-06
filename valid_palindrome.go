package main

import (
	"regexp"
	"strings"
)

func IsPalindromeValid(s string) bool {
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(strings.ToLower(s), "")

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}

		if i == len(s)-1 {
			return true
		}
	}

	return false
}
