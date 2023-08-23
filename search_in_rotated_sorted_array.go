package main

func SearchInRotated(nums []int, target int) int {
	it := map[int]bool{}
	index := -1

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			index = i
		}
		it[nums[i]] = true
	}

	if it[target] {
		return index
	}

	return index
}
