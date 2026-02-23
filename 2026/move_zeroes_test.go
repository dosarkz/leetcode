package _026

import "testing"

func TestMoveZeres(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "test1",
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.nums)
			for i := range tt.nums {
				if tt.nums[i] != tt.want[i] {
					t.Errorf("got %v, want %v", tt.nums, tt.want)
				}
			}
		})
	}
}
