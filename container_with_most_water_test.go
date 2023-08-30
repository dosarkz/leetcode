package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, MaxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}), 49)
	assert.Equal(t, MaxArea([]int{1, 1}), 1)
	assert.Equal(t, MaxArea([]int{1, 2, 1}), 2)
	assert.Equal(t, MaxArea([]int{1, 2}), 1)
	assert.Equal(t, MaxArea([]int{2, 1}), 1)
}
