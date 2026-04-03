package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		Root *TreeNode
		Max  int
	}{
		{
			Root: &TreeNode{
				Val: 1,
				Left: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
					},
				},
			},
			Max: 3,
		},
		{
			Root: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 9,
				},
				Right: &TreeNode{
					Val: 20,
					Left: &TreeNode{
						Val: 15,
					},
					Right: &TreeNode{
						Val: 7,
					},
				},
			},
			Max: 3,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Max, maxDepth(test.Root))
	}
}
