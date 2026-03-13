package _026

func pivotIndex(nums []int) int {
	totalSum := 0
	leftSum, rightSum := 0, 0

	for _, num := range nums {
		totalSum += num
	}

	for i, n := range nums {
		rightSum = totalSum - leftSum - n

		if leftSum == rightSum {
			return i
		}
		leftSum += n
	}
	return -1
}
