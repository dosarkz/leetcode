package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqualRowAndColumnPairs(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		//{grid: [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, want: 1},
		{grid: [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, want: 3},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, equalPairs(test.grid))
	}
}
