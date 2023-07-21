package main

func NumberOfSteps(num int) int {
	var count int = 0

	for num > 0 {
		if num%2 == 0 {
			// even num
			num = num / 2
		} else {
			// odd num
			num = num - 1
		}
		count++
	}
	return count
}
