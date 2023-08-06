package main

func RemoveDuplicates(nums []int) int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range nums {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	for i, _ := range nums {
		if len(list) <= i {
			nums = nums[:len(list)]
			break
		}
		nums[i] = list[i]
	}
	return len(nums)
}
