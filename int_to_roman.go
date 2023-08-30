package main

import (
	"strconv"
)

func IntToRoman(num int) string {
	str := strconv.Itoa(num)
	s := ""
	for i := 0; i < len(str); i++ {
		v, err := strconv.Atoi(string(str[i]))
		if err != nil {
			return ""
		}
		s += getRomanValue(v, len(str[i:]))
	}

	return s
}

func getRomanValue(val int, n int) string {
	switch n {
	case 4:
		return genRoman(val * 1000)
	case 3:
		return genRoman(val * 100)
	case 2:
		return genRoman(val * 10)
	case 1:
		return genRoman(val)
	}
	return ""
}

func genRoman(val int) string {
	rms := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		50:   "L",
		40:   "XL",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	res := 0
	r := ""

	for val > res {
		if rms[val] != "" {
			r += rms[val]
			break
		} else {
			if val > 1000 {
				r += "M"
				val -= 1000
			} else if val > 500 {
				r += "D"
				val -= 500
			} else if val > 100 {
				r += "C"
				val -= 100
			} else if val > 50 {
				r += "L"
				val -= 50
			} else if val > 10 {
				r += "X"
				val -= 10
			} else if val > 5 {
				r += "V"
				val -= 5
			} else if val > 1 {
				r += "I"
				val -= 1
			}
		}
	}

	return r
}
