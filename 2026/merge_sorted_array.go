package _026

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	j := 0
	for i := 0; i < m+n; i++ {
		if i < m {
			nums1[i] = nums1[i]
		} else {
			nums1[i] = nums2[j]
			j++
		}
	}

	sort.Ints(nums1)
}
