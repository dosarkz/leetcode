package _026

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"test1", args{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}}, 49},
		{"test2", args{[]int{1, 1}}, 1},
	}
	for _, tt := range tests {
		if got := maxArea(tt.args.height); got != tt.want {
			t.Errorf("%q. maxArea() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
