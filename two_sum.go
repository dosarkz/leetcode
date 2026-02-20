package main

func TwoSum(nums []int, target int) []int {
	prevMap := make(map[int]int)

	for i := 0; i <= len(nums); i++ {
		diff := target - nums[i]
		if j, ok := prevMap[diff]; ok {
			return []int{j, i}
		}
		prevMap[nums[i]] = i
	}

	return []int{}
}
