package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsMatch(t *testing.T) {
	//assert.Equal(t, IsMatch("aa", "a"), false)
	//assert.Equal(t, IsMatch("aa", "a*"), true)
	//assert.Equal(t, IsMatch("ab", ".*"), true)
	//assert.Equal(t, IsMatch("aab", "c*a*b"), true)
	//assert.Equal(t, IsMatch("mississippi", "mis*is*ip*."), true)
	//assert.Equal(t, IsMatch("ab", ".*c"), false)
	//assert.Equal(t, IsMatch("aaa", "aaaa"), false)
	assert.Equal(t, IsMatch("aaa", "ab*ac*a"), true)

}
