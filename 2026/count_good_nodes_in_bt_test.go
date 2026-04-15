package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		input *TreeNode
		want  int
	}{
		//{
		//	input: makeTreeNode([]int{3, 1, 4, 3, -1, 1, 5}),
		//	want:  4,
		//},
		//{
		//	input: makeTreeNode([]int{3, 3, -1, 4, 2}),
		//	want:  3,
		//},
		//{
		//	input: makeTreeNode([]int{1}),
		//	want:  1,
		//},
		{
			input: makeTreeNode([]int{2, -1, 4, 10, 8, -1, -1, 4}),
			want:  4,
		},
		//{
		//	input: makeTreeNode([]int{2, 9, 1}),
		//	want:  2,
		//},
		//{
		//	input: makeTreeNode([]int{9, -1, 3, 6}),
		//	want:  1,
		//},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, goodNodes(test.input))
	}
}
