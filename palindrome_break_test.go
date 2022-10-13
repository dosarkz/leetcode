package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBreakPalindrome(t *testing.T) {
	assert.Equal(t, BreakPalindrome("abccba"), "aaccba", "Tests should be equal")
	assert.Equal(t, BreakPalindrome("a"), "", "Tests should be equal")
	assert.Equal(t, BreakPalindrome("aaaaaa"), "aaaaab", "Tests should be equal")
}
