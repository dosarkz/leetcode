package main

import "strconv"

func IsPalindrome(x int) bool {
	p := strconv.Itoa(x)

	for i := 0; i < len(p); i++ {
		if p[i] != p[len(p)-i-1] {
			return false
		}

		if i == len(p)-1 {
			return true
		}
	}

	return false
}
