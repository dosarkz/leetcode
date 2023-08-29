package main

func MiddleNode(head *ListNode) *ListNode {
	node := head
	var l []int
	for node != nil {
		l = append(l, node.Val)
		node = node.Next
	}

	if len(l) == 2 {
		l = l[1:]
	}

	if len(l)%2 == 0 {
		l = l[(((len(l)-1)/2)+(len(l)/2))/2+1:]
	} else {
		l = l[(len(l)-1)/2:]
	}

	h := &ListNode{Val: l[0], Next: nil}
	curr := h
	for i := 1; i < len(l); i++ {
		curr.Next = &ListNode{Val: l[i], Next: nil}
		curr = curr.Next
	}

	return h
}
