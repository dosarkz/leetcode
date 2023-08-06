package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromeValid(t *testing.T) {
	assert.Equal(t, IsPalindromeValid("0P"), false)
}
