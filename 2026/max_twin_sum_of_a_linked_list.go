package _026

func pairSum(head *ListNode) int {
	curr := head
	var m []int
	for n := 0; curr != nil; n++ {
		m = append(m, curr.Val)
		curr = curr.Next
	}
	maxSum := 0
	for i := range len(m) / 2 {
		t := len(m) - 1 - i
		maxSum = max(maxSum, m[i]+m[t])
	}
	return maxSum
}
