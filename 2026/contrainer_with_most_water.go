package _026

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res, minAr := 0, 0
	for l < r {
		if height[l] < height[r] {
			minAr = height[l]
		} else {
			minAr = height[r]
		}

		area := (r - l) * minAr
		if area > res {
			res = area
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
