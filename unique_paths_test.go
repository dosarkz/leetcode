package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	//assert.Equal(t, UniquePaths(3, 2), 3)
	assert.Equal(t, UniquePaths(3, 7), 28)
}
