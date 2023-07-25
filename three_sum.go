package main

import (
	"fmt"
	"sort"
)

// ThreeSum use sliding Window Algorithm
func ThreeSum(nums []int) [][]int {
	res := [][]int{}
	visited := map[string]bool{}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	for i := 0; i < len(nums)-1; i++ {
		target := 0 - nums[i]
		start, end := i+1, len(nums)-1

		for start < end {
			key := fmt.Sprintf("%d+%d+%d", nums[i], nums[start], nums[end])
			sum := nums[start] + nums[end]

			if sum == target && !visited[key] {
				visited[key] = true
				res = append(res, []int{nums[i], nums[start], nums[end]})
			} else if sum < target {
				start++
			} else {
				end--
			}
		}
	}

	return res
}
