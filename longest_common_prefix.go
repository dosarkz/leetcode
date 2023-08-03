package main

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 || strs == nil {
		return ""
	}

	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		// Iterate through the string one by one from S1 till SN
		for j := 1; j < len(strs); j++ {
			// Start comparing the 0th index, 1st index … ith index simultaneously for each string.
			if i == len(strs[j]) || strs[j][i] != c {
				//In case, any of the ith index characters doesn’t match,
				//terminate the algorithm and return the LPS(1,i)
				return strs[0][0:i]
			}
		}
	}

	return strs[0]
}
