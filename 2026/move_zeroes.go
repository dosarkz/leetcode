package _026

func moveZeroes(nums []int) {
	j := 0
	n := len(nums)
	zeroCount := 0

	if n == 0 {
		return
	}

	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			zeroCount++
			continue
		}
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for i := j; i < n; i++ {
		nums[i] = 0
	}
}
