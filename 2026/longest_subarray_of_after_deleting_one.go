package _026

func longestSubarray(nums []int) int {
	left := 0
	zeroCount := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		// subtract 1 because we must delete one element
		maxLen = max(maxLen, right-left)
	}

	return maxLen
}
