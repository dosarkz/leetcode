package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStrStr(t *testing.T) {
	assert.Equal(t, StrStr("sadbutsad", "sad"), 0)
	assert.Equal(t, StrStr("leetcode", "leeto"), -1)
	assert.Equal(t, StrStr("hello", "ll"), 2)
	assert.Equal(t, StrStr("mississippi", "issip"), 4)
	assert.Equal(t, StrStr("aaa", "aaaa"), -1)
	assert.Equal(t, StrStr("mississippi", "issipi"), -1)
}
