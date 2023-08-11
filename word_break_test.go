package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWordBreak(t *testing.T) {
	assert.Equal(t, WordBreak("leetcode", []string{"leet", "code"}), true)
	assert.Equal(t, WordBreak("applepenapple", []string{"apple", "pen"}), true)
	assert.Equal(t, WordBreak("abcd", []string{"a", "abc", "b", "cd"}), true)
	assert.Equal(t, WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}), false)
	assert.Equal(t, WordBreak("catsandogcat", []string{"cats", "dog", "sand", "and", "cat", "an"}), true)

}
