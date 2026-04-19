package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPathSumThree(t *testing.T) {
	tests := []struct {
		root   *TreeNode
		target int
		want   int
	}{
		{
			root:   makeTreeNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}),
			target: 8,
			want:   3,
		},
		{
			root:   makeTreeNode([]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, 5, 1}),
			target: 22,
			want:   3,
		},
		{
			root:   makeTreeNode([]int{1}),
			target: 0,
			want:   0,
		},
		{
			root:   makeTreeNode([]int{1, 2}),
			target: 1,
			want:   1,
		},
		{
			root:   makeTreeNode([]int{-2, -1, -3}),
			target: -3,
			want:   1,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, pathSum(test.root, test.target))
	}
}
