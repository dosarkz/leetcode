package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	assert.Equal(t, RemoveElement([]int{3, 2, 2, 3}, 3), 2)
}
