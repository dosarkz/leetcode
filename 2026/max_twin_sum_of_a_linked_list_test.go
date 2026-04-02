package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxTwinSumOfALinkedList(t *testing.T) {
	tests := []struct {
		Input    *ListNode
		Expected int
	}{
		{
			Input:    arrayToListNode([]int{5, 4, 2, 1}),
			Expected: 6,
		},
		{
			Input:    arrayToListNode([]int{4, 2, 2, 3}),
			Expected: 7,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Expected, pairSum(test.Input))
	}
}
