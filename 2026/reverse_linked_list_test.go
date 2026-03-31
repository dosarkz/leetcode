package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseLinkedList(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: arrayToListNode([]int{1, 2, 3, 4, 5}),
			want: arrayToListNode([]int{5, 4, 3, 2, 1}),
		},
	}

	for _, test := range tests {
		actual := reverseList(test.head)
		assert.Equal(t, test.want, actual)
	}
}
