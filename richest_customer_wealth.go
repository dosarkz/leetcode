package main

func maximumWealth(accounts [][]int) int {
	var res int
	var mw int
	for i, v := range accounts {
		for j := 0; j < len(v); j++ {
			res += v[j]
		}

		if mw < res {
			mw = res
		}

		if i == len(accounts)-1 {
			return mw
		}
		res = 0
	}

	return mw
}
