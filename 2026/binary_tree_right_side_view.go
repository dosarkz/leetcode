package _026

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := []*TreeNode{root}
	var result []int

	for len(queue) > 0 {
		levelSize := len(queue)
		var rightmost int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			// The last node in this level is the rightmost
			rightmost = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		result = append(result, rightmost)
	}

	return result
}
