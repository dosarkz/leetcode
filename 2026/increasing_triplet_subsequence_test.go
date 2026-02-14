package _026

import "testing"

func TestIncreasingTripletSubSequence(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		//{"test1", []int{1, 2, 3, 4, 5}, true},
		//{"test2", []int{5, 4, 3, 2, 1}, false},
		{"test3", []int{2, 1, 5, 0, 4, 6}, true},
	}
	for _, tt := range tests {
		if got := increasingTriplet(tt.nums); got != tt.want {
			t.Errorf("%s: increasingTriplet(%v) = %v, want %v", tt.name, tt.nums, got, tt.want)
		}
	}
}
