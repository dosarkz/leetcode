package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root *TreeNode
		p, q *TreeNode
		want int
	}{
		//{
		//	root: makeTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
		//	p:    makeTreeNode([]int{5}),
		//	q:    makeTreeNode([]int{1}),
		//	want: 3,
		//},
		{
			root: makeTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}),
			p:    makeTreeNode([]int{5}),
			q:    makeTreeNode([]int{4}),
			want: 5,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, lowestCommonAncestor(test.root, test.p, test.q).Val)
	}
}
