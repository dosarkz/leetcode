package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLeafSimilarTree(t *testing.T) {
	tests :=
		[]struct {
			root1 *TreeNode
			root2 *TreeNode
			want  bool
		}{
			{
				root1: makeTreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}),
				root2: makeTreeNode([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8}),
				want:  true,
			},
			{
				root1: makeTreeNode([]int{1, 2, 3}),
				root2: makeTreeNode([]int{1, 3, 2}),
				want:  false,
			},
			{
				root1: makeTreeNode([]int{1}),
				root2: makeTreeNode([]int{2}),
				want:  false,
			},
			{
				root1: makeTreeNode([]int{41, 62, -1, 66, -1, -1, 21, 96, -1, 70, 74}),
				root2: makeTreeNode([]int{55, -1, 84, -1, 29, 116, -1, 7, 74, -1, 70}),
				want:  true,
			},
			{
				root1: makeTreeNode([]int{1}),
				root2: makeTreeNode([]int{2, 1}),
				want:  true,
			},
		}

	for _, test := range tests {
		assert.Equal(t, test.want, leafSimilar(test.root1, test.root2))
	}
}

func makeTreeNode(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}

	if arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(arr) && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(arr) && arr[i] != -1 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(arr) && arr[i] != -1 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root

}
