package _026

func longestOnes(nums []int, k int) int {
	maxLength := 0
	currentCount := 0

	for i := 0; i < len(nums); i++ {
		if k == 0 && nums[i] == 0 {
			if nums[i-currentCount] == 0 {
				continue
			}
			for nums[i-currentCount] == 1 {
				currentCount--
			}
		}

		// flip 0 to 1 if we have flips left
		if nums[i] == 0 && k > 0 {
			currentCount++
			k--
		}

		if nums[i] == 1 {
			currentCount++
		}

		if currentCount > maxLength {
			maxLength = currentCount
		}
	}

	return maxLength

}
