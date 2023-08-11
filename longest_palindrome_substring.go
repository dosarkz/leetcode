package main

func LongestPalindrome(s string) string {
	l := 0

	for i := 0; i < len(s)-1/2; i++ {
		l = len(s) - i

		for j := 0; j <= i; j++ {
			if isPal(s[j : l+j]) {
				return s[j : l+j]
			}
		}
	}
	return ""
}

func isPal(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}
