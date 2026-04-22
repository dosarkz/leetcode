package _026

func longestZigZag(root *TreeNode) int {
	count := 0
	var maxZigZag func(node *TreeNode, l, r int)

	maxZigZag = func(node *TreeNode, l, r int) {
		if node != nil {
			count = max(count, l, r)
			maxZigZag(node.Left, r+1, 0)
			maxZigZag(node.Right, 0, l+1)
		}
	}

	maxZigZag(root, 0, 0)
	return count
}
