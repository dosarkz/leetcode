package _026

func pathSum(root *TreeNode, targetSum int) int {
	var dfsSum func(node *TreeNode, stack []int, sum int) int

	dfsSum = func(node *TreeNode, stack []int, sum int) int {
		if node == nil {
			return 0
		}

		count, currSum := 0, 0
		stack = append(stack, node.Val)

		for i := len(stack) - 1; i >= 0; i-- {
			currSum += stack[i]
			if currSum == targetSum {
				count++
			}
		}

		count += dfsSum(node.Left, stack, sum)
		count += dfsSum(node.Right, stack, sum)

		return count
	}

	return dfsSum(root, []int{}, 0)
}
