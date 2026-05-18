package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxLevelSumOfABinaryTree(t *testing.T) {
	tests := []struct {
		input *TreeNode
		want  int
	}{
		{
			makeTreeNode([]int{1, 7, 0, 7, -8, -1, -1}),
			2,
		},
		{
			makeTreeNode([]int{989, -1, 10250, 98693, -89388, -1, -1, -1, -32127}),
			2,
		},
		{
			makeTreeNode([]int{-100, -200, -300, -20, -5, -10, -1}),
			3,
		},
	}
	for _, test := range tests {
		assert.Equal(t, test.want, maxLevelSum(test.input))
	}
}
