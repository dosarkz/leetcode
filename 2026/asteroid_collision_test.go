package _026

import "testing"

func TestAsteroidCollision(t *testing.T) {
	tests := []struct {
		astr []int
		want []int
	}{
		//{[]int{5, 10, -5}, []int{5, 10}},
		//{[]int{8, -8}, []int{}},
		//{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}

	for _, tt := range tests {
		got := asteroidCollision(tt.astr)
		if len(got) != len(tt.want) {
			t.Errorf("got %v, want %v", got, tt.want)
		}
	}
}
