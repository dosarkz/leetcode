package _026

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	total, level, result := 0, 1, 0

	for len(queue) > 0 {
		levelSize := len(queue)
		var sum int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			sum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if total < sum {
			total = sum
			result = level
		}

		// check for a negative sum
		if sum < 0 && total == 0 {
			total = sum
		}

		level++
	}
	return result
}
