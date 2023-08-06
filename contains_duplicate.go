package main

func ContainsDuplicate(nums []int) bool {
	set := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if set[nums[i]] > 0 {
			return true
		}
		set[nums[i]] = +1
	}
	return false
}
