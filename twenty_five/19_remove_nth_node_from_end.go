package twenty_five

/**
 Task:
	Given the head of a linked list,
	remove the nth node from the end of the list and return its head.

	Input: head = [1,2,3,4,5], n = 2
	Output: [1,2,3,5]

	Input: head = [1,2], n = 1
	Output: [1]

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

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head
	left := dummy
	right := head
	for i := 0; i < n; i++ {
		right = right.Next
	}
	for right != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummy.Next
}
