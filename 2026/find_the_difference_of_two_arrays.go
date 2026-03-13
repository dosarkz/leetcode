package _026

func findDifference(nums1 []int, nums2 []int) [][]int {
	m1 := make(map[int]bool)
	for _, n := range nums1 {
		m1[n] = true
	}

	m2 := make(map[int]bool)
	for _, n := range nums2 {
		m2[n] = true
	}

	res := make([][]int, 2)
	for n := range m1 {
		if !m2[n] {
			res[0] = append(res[0], n)
		}
	}

	for n := range m2 {
		if !m1[n] {
			res[1] = append(res[1], n)
		}
	}

	return res
}
