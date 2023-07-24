package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTwoSum(t *testing.T) {
	assert.Equal(t, TwoSum([]int{2, 7, 11, 15}, 9), []int{0, 1}, "Tests should be equal")
	assert.Equal(t, TwoSum([]int{3, 2, 3}, 6), []int{0, 2}, "Tests should be equal")
}
