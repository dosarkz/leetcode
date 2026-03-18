package _026

import "testing"

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		//{[]int{1}, 1, []int{}, 0, []int{1}},
		//{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for _, test := range tests {
		merge(test.nums1, test.m, test.nums2, test.n)
		for i := range test.nums1 {
			if test.nums1[i] != test.want[i] {
				t.Errorf("For input (%v, %d, %v, %d), expected %v but got %v", test.nums1, test.m, test.nums2, test.n, test.want, test.nums1)
				break
			}
		}
	}
}
