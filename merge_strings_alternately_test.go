package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeAlternately(t *testing.T) {
	assert.Equal(t, MergeAlternately("abc", "pqr"), "apbqcr", "Test should be equal apbqcr")
	assert.Equal(t, MergeAlternately("ab", "pqrs"), "apbqrs", "Test should be equal apbqcr")
	assert.Equal(t, MergeAlternately("abcd", "pq"), "apbqcd", "Test should be equal apbqcr")
}
