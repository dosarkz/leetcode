package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMiddleNode(t *testing.T) {
	//assert.Equal(t, MiddleNode(insertNode([]int{1, 2, 3, 4, 5})), insertNode([]int{3, 4, 5}))
	////assert.Equal(t, MiddleNode(insertNode([]int{1, 2, 3, 4, 5, 6})), insertNode([]int{4, 5, 6}))
	//assert.Equal(t, MiddleNode(insertNode([]int{1})), insertNode([]int{1}))
	assert.Equal(t, MiddleNode(insertNode([]int{1, 2})), insertNode([]int{2}))
}
