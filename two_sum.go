package main

func TwoSum(nums []int, target int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}

			if nums[i]+nums[j] == target {
				return append(res, i, j)
			}
		}
	}
	return res
}
