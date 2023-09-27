package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOne(t *testing.T) {
	assert.Equal(t, PlusOne([]int{1, 2, 3}), []int{1, 2, 4})
	assert.Equal(t, PlusOne([]int{4, 3, 2, 1}), []int{4, 3, 2, 2})
	assert.Equal(t, PlusOne([]int{9}), []int{1, 0})
	assert.Equal(t, PlusOne([]int{9, 9}), []int{1, 0, 0})
}
