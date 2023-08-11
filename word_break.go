package main

func WordBreak(s string, wordDict []string) bool {
	var res = false
	var a string

	for i := 0; i < len(wordDict); i++ {
		a = wordDict[i]
		for j := 0; j < len(wordDict); j++ {
			if j+1 > len(wordDict)-1 {
				continue
			}

			b := wordDict[j]
			c := wordDict[j+1]

			if a+b == s ||
				a+c == s ||
				a+b+c == s ||
				a+c+b == s ||
				a+b+a == s ||
				a+c+a == s ||
				b+a == s ||
				b+c == s ||
				b+a+c == s ||
				b+c+a == s ||
				b+c+b == s ||
				b+a+b == s {
				return true
			}
		}

		if a == s {
			res = true
		}
	}

	return res
}
