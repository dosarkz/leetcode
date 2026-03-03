package _026

import "testing"

func TestMaxOperations(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 2, 3, 4}, 5}, 2},
		{"test2", args{[]int{3, 1, 3, 4, 3}, 6}, 1},
		{"test1", args{[]int{2, 2, 2, 3, 1, 1, 4, 1}, 4}, 2},
		{"test3", args{[]int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2}, 2},
	}

	for _, tt := range tests {
		if got := maxOperations(tt.args.nums, tt.args.k); got != tt.want {
			t.Errorf("%q. maxOperations() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
