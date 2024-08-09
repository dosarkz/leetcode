package twenty_four

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i = i + 1 {
		m[nums[i]]++
		if m[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return -1
}
