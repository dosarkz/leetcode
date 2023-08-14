package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	assert.Equal(t, FindMedianSortedArrays([]int{1, 3}, []int{2}), 2.00000)
	assert.Equal(t, FindMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.50000)
}
