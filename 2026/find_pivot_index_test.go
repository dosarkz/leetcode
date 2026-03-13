package _026

import "testing"

func TestFindPivotIndex(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "test1",
			nums: []int{1, 7, 3, 6, 5, 6},
			want: 3,
		},
		{
			name: "test2",
			nums: []int{1, 2, 3},
			want: -1,
		},
		{
			name: "test3",
			nums: []int{2, 1, -1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pivotIndex(tt.nums); got != tt.want {
				t.Errorf("pivotIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
