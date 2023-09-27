package main

func SwapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	curr := head
	newHead := head.Next

	for curr != nil && curr.Next != nil {
		temp := curr.Next
		curr.Next = curr.Next.Next
		temp.Next = curr
		curr = curr.Next

		if curr != nil && curr.Next != nil {
			temp.Next.Next = curr.Next
		}

	}

	return newHead
}
