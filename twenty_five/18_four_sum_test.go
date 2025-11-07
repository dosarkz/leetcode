package twenty_five

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want: [][]int{
				{-2, -1, 1, 2},
				{-2, 0, 0, 2},
				{-1, 0, 0, 1},
			},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want: [][]int{
				{2, 2, 2, 2},
			},
		},
		{
			nums:   []int{},
			target: 0,
			want:   [][]int{},
		},
		{
			nums:   []int{0, 0, 0, 0},
			target: 0,
			want: [][]int{
				{0, 0, 0, 0},
			},
		},
	}

	for _, tt := range tests {
		got := fourSum(tt.nums, tt.target)
		if !equalUnordered(got, tt.want) {
			t.Errorf("fourSum(%v, %d) = %v, want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

// Helper to compare slices of slices regardless of order
func equalUnordered(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	used := make([]bool, len(b))
	for _, x := range a {
		found := false
		for i, y := range b {
			if !used[i] && reflect.DeepEqual(x, y) {
				used[i] = true
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}
