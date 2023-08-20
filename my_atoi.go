package main

import (
	"strings"
)

func MyAtoi(s string) int {
	var (
		v byte
		n int64
		m bool
	)

	s = strings.Trim(s, " ")

	for i := 0; i < len(s); i++ {
		d := s[i]

		if s[i] == '-' && i == 0 {
			m = true
			continue
		}

		if s[i] == '+' && i == 0 {
			continue
		}

		if d >= '0' && d <= '9' {
			v = d - '0'
		} else {
			break
		}
		n *= int64(10)
		n += int64(v)

		if m {
			if -n <= int64(-2147483648) {
				return -2147483648
			}
		}

		if n > int64(2147483647) {
			return 2147483647
		}

	}

	if m {
		n = -(n)
	}

	return int(n)
}
