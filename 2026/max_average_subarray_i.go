package _026

import (
	"math"
)

func findMaxAverage(nums []int, k int) float64 {
	maxSum := math.MinInt32
	for window, sum, right := 0, 0, 0; right < len(nums); right++ {
		if window < k {
			window++
			sum += nums[right]
			maxSum = sum
		} else {
			sum = sum + nums[right] - nums[right-window]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return float64(maxSum) / float64(k)
}
