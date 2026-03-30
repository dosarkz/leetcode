package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteMiddleNodeOfALinkedList(t *testing.T) {
	tests := []struct {
		Input  *ListNode
		Output *ListNode
	}{
		//{
		//	Input:  arrayToListNode([]int{1, 2, 3, 4}),
		//	Output: arrayToListNode([]int{1, 2, 4}),
		//},
		//{
		//	Input:  arrayToListNode([]int{1, 3, 4, 7, 1, 2, 6}),
		//	Output: arrayToListNode([]int{1, 3, 4, 1, 2, 6}),
		//},
		{
			Input:  arrayToListNode([]int{1}),
			Output: arrayToListNode([]int{}),
		},
	}

	for _, test := range tests {
		head := deleteMiddle(test.Input)
		assert.Equal(t, head, test.Output)
	}
}

func arrayToListNode(val []int) *ListNode {
	if len(val) == 0 {
		return nil
	}
	head := &ListNode{Val: val[0]}
	curr := head
	for i := 1; i < len(val); i++ {
		curr.Next = &ListNode{Val: val[i]}
		curr = curr.Next
	}
	return head
}
