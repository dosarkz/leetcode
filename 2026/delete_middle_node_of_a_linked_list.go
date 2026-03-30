package _026

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	curr := head
	idx := 1
	middleIndex := 0
	for curr.Next != nil {
		curr = curr.Next
		idx++
		middleIndex = idx / 2
	}

	m := head

	for i := 0; i < middleIndex; i++ {
		if middleIndex == i+1 {
			head.Next = head.Next.Next
		}
		head = head.Next
	}
	return m
}

func printList(head *ListNode) []int {
	var arr []int
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return arr
}
