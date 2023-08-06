package main

func MyPow(x float64, n int) float64 {
	res := 1.0

	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
		x = 1 / x
	}

	for n >= 1 {
		if n%2 == 1 {
			res = res * x
		}
		x = x * x
		n = n / 2
	}

	return res
}
