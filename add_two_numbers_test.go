package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	//assert.Equal(t, AddTwoNumbers(insertNode([]int{2, 4, 3}), insertNode([]int{5, 6, 4})), insertNode([]int{7, 0, 8}), "test should be equal")
	//assert.Equal(t, AddTwoNumbers(&ListNode{Val: 0}, &ListNode{Val: 0}), &ListNode{Val: 0}, "test should be equal")

	assert.Equal(t, AddTwoNumbers(
		insertNode([]int{9, 9, 9, 9, 9, 9, 9}),
		insertNode([]int{9, 9, 9, 9})),
		insertNode([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		"test should be equal")
}

func insertNode(items []int) *ListNode {
	head := &ListNode{Val: items[0], Next: nil}
	curr := head
	for i := 1; i < len(items); i++ {
		curr.Next = &ListNode{Val: items[i], Next: nil}
		curr = curr.Next
	}
	return head
}
