package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSearchInRotated(t *testing.T) {
	assert.Equal(t, SearchInRotated([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	assert.Equal(t, SearchInRotated([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
	assert.Equal(t, SearchInRotated([]int{1}, 0), -1)
	assert.Equal(t, SearchInRotated([]int{1}, 1), 0)
	assert.Equal(t, SearchInRotated([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
}
