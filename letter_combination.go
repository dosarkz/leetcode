package main

func LetterCombinations(digits string) []string {
	res := []string{}

	if len(digits) == 0 {
		return res
	}
	n := len(digits)

	mapDig := map[byte][]string{
		50: {"a", "b", "c"},
		51: {"d", "e", "f"},
		52: {"g", "h", "i"},
		53: {"j", "k", "l"},
		54: {"m", "n", "o"},
		55: {"p", "q", "r", "s"},
		56: {"t", "u", "v"},
		57: {"w", "x", "y", "z"},
	}

	res = mapDig[digits[0]]
	if n == 1 {
		return res
	}

	for i := 1; i < len(digits); i++ {
		curr := mapDig[digits[i]]
		res = mergeLetters(res, curr)
	}

	return res
}

func mergeLetters(curr, next []string) []string {
	var res []string

	for _, x := range curr {
		for _, y := range next {
			res = append(res, x+y)
		}
	}

	return res
}
