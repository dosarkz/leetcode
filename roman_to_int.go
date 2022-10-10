package main

func main() {
	println("test", romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	rms := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var res int
	for i, char := range s {
		l := string(char)
		var next string
		var prev string

		if i > 0 {
			prev = string(s[i-1])
		}

		if i != len(s)-1 {
			next = string(s[i+1])
		}

		if rms[l] == 0 || (i > 0 && isSubNumber(prev, l)) {
			continue
		}

		if isSubNumber(l, next) {
			if i == len(s) {
				continue
			}

			sub := getSubValue(l, next)
			if sub != 0 {
				res += sub
				continue
			}

		}

		res += rms[l]
	}

	return res
}

func isSubNumber(l string, next string) bool {
	for _, r := range []string{"I", "X", "C"} {
		if l == r && getSubValue(l, next) > 0 {
			return true
		}
	}

	return false
}

func getSubValue(l string, next string) int {
	switch l {
	case "I":
		if next == "V" {
			return 4
		}
		if next == "X" {
			return 9
		}
	case "X":
		if next == "L" {
			return 40
		}
		if next == "C" {
			return 90
		}
	case "C":
		if next == "D" {
			return 400
		}
		if next == "M" {
			return 900
		}
	}

	return 0
}
