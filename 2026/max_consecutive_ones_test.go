package _026

import "testing"

func TestLongestOnes(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected int
	}{
		//{[]int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}, 2, 6},
		//{[]int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}, 3, 10},
		{[]int{0, 0, 0}, 2, 2},
	}

	for _, test := range tests {
		result := longestOnes(test.nums, test.k)
		if result != test.expected {
			t.Errorf("For input (%v, %d), expected %d but got %d", test.nums, test.k, test.expected, result)
		}
	}
}
