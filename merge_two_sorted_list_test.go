package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	assert.Equal(t, MergeTwoLists(
		insertNode([]int{1, 2, 4}),
		insertNode([]int{1, 3, 4})),
		insertNode([]int{1, 1, 2, 3, 4, 4}))
}
