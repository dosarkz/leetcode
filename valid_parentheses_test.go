package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	assert.Equal(t, IsValid("{[]}"), true, "Tests should be equal")
	assert.Equal(t, IsValid("()[]{}"), true, "Tests should be equal")
	assert.Equal(t, IsValid("()"), true, "Tests should be equal")
	assert.Equal(t, IsValid("(]"), false, "Tests should be equal")
	assert.Equal(t, IsValid("["), false, "Tests should be equal")
	assert.Equal(t, IsValid("(("), false, "Tests should be equal")
}
