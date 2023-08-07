package main

import (
	"fmt"
	"strconv"
)

func ReverseInt(x int) int {
	neg := false
	if x < 0 {
		neg = true
		x = -x
	}
	p := strconv.Itoa(x)
	var a string

	for i := 0; i < len(p); i++ {

		if len(p)-i-1 >= 0 {
			a += string(p[len(p)-i-1])
		}
	}
	fmt.Println(a)

	res, _ := strconv.Atoi(a)
	if neg {
		res = -res
	}

	if res < -2147483648 || res > 2147483648 {
		return 0
	}

	return res
}
