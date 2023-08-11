package main

func RemoveElement(nums []int, val int) int {
	var k = 0
	var s []int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			k++
			s = append(s, nums[i])
		}
	}

	for j := 0; j < len(nums); j++ {
		if len(s) > j {
			nums[j] = s[j]
		}

	}
	return k
}
