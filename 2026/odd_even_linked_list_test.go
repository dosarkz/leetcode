package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOddEvenLinkedList(t *testing.T) {
	tests := []struct {
		Input  *ListNode
		Output *ListNode
	}{
		{
			Input:  arrayToListNode([]int{1, 2, 3, 4, 5}),
			Output: arrayToListNode([]int{1, 3, 5, 2, 4}),
		},
		{
			Input:  arrayToListNode([]int{2, 1, 3, 5, 6, 4, 7}),
			Output: arrayToListNode([]int{2, 3, 6, 7, 1, 5, 4}),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Output, addEventList(test.Input))
	}
}
