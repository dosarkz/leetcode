package _026

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	numI, numJ := nums[0], int(^uint(0)>>1)
	for i := 1; i < len(nums); i++ {
		if nums[i] <= numI {
			numI = nums[i]
		} else if nums[i] <= numJ {
			numJ = nums[i]
		} else {
			return true
		}
	}

	return false
}
