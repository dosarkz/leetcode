package _026

import "sort"

func maxOperations(nums []int, k int) int {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	count := 0
	for l < r {
		sum := nums[l] + nums[r]

		if sum == k {
			count++
			l++
			r--
		} else if sum > k {
			r--
		} else {
			l++
		}
	}

	return count
}
