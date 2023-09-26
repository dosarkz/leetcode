package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func testGenerateParenthesis(t *testing.T) {
	assert.Equal(t, GenerateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
}
