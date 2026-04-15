package _026

func goodNodes(root *TreeNode) int {
	var count int // Initialize a variable to count good nodes.

	// Define a helper function for the recursive traversal.
	var dfsMax func(node *TreeNode, max int)

	// The dfs function performs a depth-first search traversal of the binary tree.
	dfsMax = func(node *TreeNode, max int) {
		if node == nil {
			return // Base case: If the node is nil, return without processing.
		}

		if node.Val >= max {
			max = node.Val // Update the maximum value encountered so far.
			count++        // Increment the count to indicate a good node.
		}

		// Recursively traverse the left and right subtrees with the updated max value.
		dfsMax(node.Left, max)
		dfsMax(node.Right, max)
	}

	dfsMax(root, root.Val) // Start the traversal with the root node and its value.

	return count
}
