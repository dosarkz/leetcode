package main

func runningSum(nums []int) []int {
	var res []int
	var curr int
	for i := 0; i < len(nums); i++ {
		curr += nums[i]
		res = append(res, curr)
	}
	return res
}
