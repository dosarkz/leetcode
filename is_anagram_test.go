package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, IsAnagram("anagram", "nagaram"), true)
}
