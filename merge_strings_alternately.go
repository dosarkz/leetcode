package main

func MergeAlternately(word1 string, word2 string) string {
	var word string
	n1, n2 := len(word1), len(word2)
	n1i, n2i, q := 0, 0, 0

	for i := 0; i < n1+n2; i++ {
		if q == 0 && n1i < n1 {
			word += string(word1[n1i])
			n1i++
			if n2i < n2 {
				q = 1
			}
		} else if q == 1 && n2i < n2 {
			word += string(word2[n2i])
			n2i++
			if n1i < n1 {
				q = 0
			}
		}
	}

	return word
}
