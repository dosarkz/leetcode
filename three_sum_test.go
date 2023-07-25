package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	assert.Equal(t, ThreeSum([]int{-1, 0, 1, 2, -1, -4}), [][]int{{-1, -1, 2}, {-1, 0, 1}}, "Test does not passed")
	assert.Equal(t, ThreeSum([]int{0, 1, 1}), [][]int{}, "Test does not passed")
	assert.Equal(t, ThreeSum([]int{0, 0, 0}), [][]int{{0, 0, 0}}, "Test does not passed")
}
