package main

func PlusOne(digits []int) []int {
	mod := 0
	for i := len(digits) - 1; i >= 0; i-- {
		if mod == 1 || len(digits)-1 == i {
			if len(digits)-1 == i {
				digits[i] = digits[i] + 1
			} else {
				digits[i] = digits[i] + mod
			}

			if digits[i] == 10 {
				digits[i] = 0
				mod = 1
				if i == 0 {
					digits = append([]int{1}, digits...)
				}

			} else {
				mod = 0
			}
		}
	}

	return digits
}
