package main

import (
	"strings"
)

func LengthOfLastWord(s string) int {
	a := strings.Split(strings.TrimSpace(s), " ")
	return len(strings.Join(a[len(a)-1:], ""))
}
