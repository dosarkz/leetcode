package main

func UniquePaths(m int, n int) int {
	ans := 1
	for i := 1; i <= m-1; i++ {
		ans = ans * (n - 1 + i) / i
	}
	return ans
}
