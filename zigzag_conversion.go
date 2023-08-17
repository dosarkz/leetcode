package main

import (
	"strings"
)

func Convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	z := make([]string, numRows)
	st, diag := 0, 0
	sn := 0
	step := numRows - 2

	for i := 0; i < len(s); i++ {

		if st < numRows {
			z[sn] += string(s[i])
			sn++
			st++
			if st == numRows {
				if numRows > 2 {
					st = numRows
				} else {
					st = 0
				}

				diag = 0
				sn = 0

			}
		} else if diag <= numRows-2 && numRows > 2 {
			z[step] += string(s[i])
			step--
			diag++

			if diag == numRows-2 {
				diag = numRows
				st = 0
				step = numRows - 2
			}
		}
	}

	return strings.Join(z, "")
}
