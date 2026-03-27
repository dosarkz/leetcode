package _026

import (
	"strconv"
	"unicode"
)

func decodeString(s string) string {
	var stack []int32

	for i, ch := range s {
		if s[i] != ']' {
			stack = append(stack, ch)
		} else {
			subst := ""
			key, value := "", ""

			for stack[len(stack)-1] != '[' {
				subst = string(stack[len(stack)-1]) + subst
				stack = stack[:len(stack)-1]
			}

			stack = stack[:len(stack)-1]

			for len(stack) > 0 && unicode.IsDigit(stack[len(stack)-1]) {
				key = string(stack[len(stack)-1]) + key
				stack = stack[:len(stack)-1]
			}

			keyInt, err := strconv.Atoi(key)
			if err != nil {
				break
			}

			for g := 0; g < keyInt; g++ {
				value += subst
			}

			for k := 0; k < len(value); k++ {
				stack = append(stack, rune(value[k]))
			}
		}
	}

	return string(stack)
}
