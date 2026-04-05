package _026

import (
	"slices"
)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var node1, node2 []int
	leavesOne := dfs(root1, node1)
	leavesTwo := dfs(root2, node2)

	return slices.Equal(leavesOne, leavesTwo)
}

func dfs(root *TreeNode, node []int) []int {
	if root == nil {
		return node
	}
	if root.Left == nil && root.Right == nil {
		node = append(node, root.Val)
		return node
	}

	node = dfs(root.Left, node)
	node = dfs(root.Right, node)

	return node
}
