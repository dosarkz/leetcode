package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestZigzagInBinaryTree(t *testing.T) {
	tests := []struct {
		input []int
		want  int
	}{
		{
			input: []int{1, -1, 1, 1, 1, -1, -1, 1, 1, -1, 1, -1, -1, -1, 1},
			want:  3,
		},
		{
			input: []int{1, 1, 1, -1, 1, -1, -1, 1, 1, -1, 1},
			want:  4,
		},
		{
			input: []int{1},
			want:  0,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, longestZigZag(makeTreeNode(test.input)))
	}
}
