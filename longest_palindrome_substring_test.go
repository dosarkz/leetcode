package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, LongestPalindrome("babad"), "bab")
	assert.Equal(t, LongestPalindrome("cbbd"), "bb")
}
