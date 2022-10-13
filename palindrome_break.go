package main

func BreakPalindrome(palindrome string) string {
	if len(palindrome) == 1 {
		return ""
	}

	pals := []rune(palindrome)
	for i, c := range pals[:len(pals)/2] {
		if string(c) != "a" {
			pals[i] = []rune("a")[0]
			return string(pals)
		}
	}

	pals[len(pals)-1] = []rune("b")[0]
	return string(pals)
}
