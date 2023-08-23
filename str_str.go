package main

func StrStr(haystack string, needle string) int {
	h, n := haystack, needle
	start := 0
	for i := 0; i < len(h); i++ {
		if len(h) < len(n) {
			return -1
		}

		if h[i] == n[0] && len(h)-i >= len(n) {
			start = i
			if n == h[start:start+(len(needle))] {
				return i
			}
		}
	}

	return -1
}
