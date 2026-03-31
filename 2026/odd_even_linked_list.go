package _026

func addEventList(head *ListNode) *ListNode {
	curr := head

	if head == nil {
		return head
	}
	var node = &ListNode{Val: head.Val}
	currNode := node

	var tail []int

	for i := 0; curr != nil; i++ {
		if i%2 == 0 {
			if i == 0 {
				curr = curr.Next
				continue
			}
			currNode.Next = &ListNode{Val: curr.Val}
			currNode = currNode.Next
		} else {
			tail = append(tail, curr.Val)
		}

		curr = curr.Next
	}

	for _, i := range tail {
		currNode.Next = &ListNode{Val: i}
		currNode = currNode.Next
	}

	return node
}
