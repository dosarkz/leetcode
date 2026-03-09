package _026

import "testing"

func TestMaxAverageSubArrayI(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"test1", args{[]int{1, 12, -5, -6, 50, 3}, 4}, 12.75000},
		{"test2", args{[]int{5}, 1}, 5.00000},
		{"test3", args{[]int{-1}, 1}, -1.00000},
		{"test4", args{[]int{0, 1, 1, 3, 3}, 4}, 2},
	}

	for _, tt := range tests {
		if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
			t.Errorf("%q. findMaxAverage() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
