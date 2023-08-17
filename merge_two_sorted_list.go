package main

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	node := &ListNode{}
	curr := node

	for list1 != nil && list2 != nil {
		n1, n2 := 0, 0

		if list1 != nil {
			n1 = list1.Val
		}
		if list2 != nil {
			n2 = list2.Val
		}

		if n1 < n2 {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}

		curr = curr.Next

		if list1 == nil {
			curr.Next = list2
		} else {
			curr.Next = list1
		}

	}

	return node.Next
}
