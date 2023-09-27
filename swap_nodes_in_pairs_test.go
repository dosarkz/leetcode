package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapPairs(t *testing.T) {
	assert.Equal(t, SwapPairs(insertNode([]int{1, 2, 3, 4})), insertNode([]int{2, 1, 4, 3}))
}
