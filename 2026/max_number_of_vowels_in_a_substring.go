package _026

func maxVowels(s string, k int) int {
	vowels := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}

	maxCount := 0
	currentCount := 0

	for i := 0; i < len(s); i++ {
		if vowels[s[i]] {
			currentCount++
		}

		if i >= k {
			// remove the leftmost character from the count
			if vowels[s[i-k]] {
				currentCount--
			}
		}

		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount
}
