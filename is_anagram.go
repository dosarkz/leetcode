package main

import (
	"sort"
	"strings"
)

func IsAnagram(s string, t string) bool {
	a := strings.Split(s, "")
	sort.Strings(a)

	b := strings.Split(t, "")
	sort.Strings(b)

	if strings.Join(a, " ") == strings.Join(b, " ") {
		return true
	}
	return false
}
