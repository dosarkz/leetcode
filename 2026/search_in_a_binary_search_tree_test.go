package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInABinarySearchTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		val  int
		want *TreeNode
	}{
		{
			root: makeTreeNode([]int{4, 2, 7, 1, 3}),
			val:  2,
			want: makeTreeNode([]int{2, 1, 3}),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, searchBST(test.root, test.val))
	}
}
