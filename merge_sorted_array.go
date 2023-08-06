package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = nums1[0:m]
	nums2 = nums2[0:n]

	for _, v := range nums2 {
		nums1 = append(nums1, v)
	}
	sort.Ints(nums1)
}
