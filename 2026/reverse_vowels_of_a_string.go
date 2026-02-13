package _026

import "fmt"

func reverseVowels(s string) string {
	l, r := 0, len(s)-1
	// find the first vowel from left and right, then swap them
	for l < r {
		fmt.Printf("l: %d, r: %d, s: %s\n", l, r, s)

		if !isVowel(s[l]) {
			l++
			continue
		}

		if !isVowel(s[r]) {
			r--
			continue
		}

		if isVowel(s[l]) && isVowel(s[r]) {
			s = s[:l] + string(s[r]) + s[l+1:r] + string(s[l]) + s[r+1:]
			l++
			r--
		}

	}
	return s
}

func isVowel(c byte) bool {
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
		c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
}
