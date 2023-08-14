package main

import (
	"sort"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := 0.00000
	arr := append(nums1, nums2...)
	sort.Ints(arr)
	n := len(arr)

	if n%2 != 0 {
		// even
		res = float64(arr[n/2])
	} else {
		// odd
		//fmt.Printf("res %f\n", 5.0/2)
		res = float64(arr[(n-1)/2]+arr[n/2]) / 2
	}

	return res
}
