package _026

import "testing"

func TestLongestSubarrayOfAfterDeletingOne(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		//{
		//	name: "test1",
		//	nums: []int{1, 1, 0, 1},
		//	want: 3,
		//},
		{
			name: "test2",
			nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			want: 5,
		},
		//{
		//	name: "test3",
		//	nums: []int{1, 1, 1},
		//	want: 2,
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
