package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	assert.Equal(t, LengthOfLongestSubstring("abcabcbb"), 3, "Tests should be equal")
	//assert.Equal(t, LengthOfLongestSubstring("bbbbb"), 1, "Tests should be equal")
	//assert.Equal(t, LengthOfLongestSubstring("pwwkew"), 3, "Tests should be equal")
}
