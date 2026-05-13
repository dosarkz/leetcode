package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRightSideView(t *testing.T) {
	tests := []struct {
		input *TreeNode
		wants []int
	}{
		{
			makeTreeNode([]int{1, 2, 3, -1, 5, -1, 4}),
			[]int{1, 3, 4},
		},
		{
			makeTreeNode([]int{1, 2, 3, 4, -1, -1, -1, 5}),
			[]int{1, 3, 4, 5},
		},
		{
			makeTreeNode([]int{1, -1}),
			[]int{1},
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.wants, rightSideView(test.input))
	}
}
