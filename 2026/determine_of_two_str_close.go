package _026

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	count1 := make(map[byte]int)
	count2 := make(map[byte]int)

	// count the frequency of each character in both strings
	for i := 0; i < len(word1); i++ {
		count1[word1[i]]++
		count2[word2[i]]++
	}

	freq1 := make(map[int]int)
	freq2 := make(map[int]int)

	// calculate the frequency of each count
	for _, count := range count1 {
		freq1[count]++
	}

	for _, count := range count2 {
		freq2[count]++
	}

	// compare the frequency maps should be the same for both strings
	if len(freq1) != len(freq2) {
		return false
	}

	for key, _ := range freq1 {
		if freq2[key] != freq1[key] {
			return false
		}
	}

	// check if the frequency of counts matches for both strings
	for key, _ := range count1 {
		if _, ok := count2[key]; !ok {
			return false
		}
	}

	return true
}
