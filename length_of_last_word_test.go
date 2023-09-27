package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	//assert.Equal(t, LengthOfLastWord("Hello World"), 5)
	assert.Equal(t, LengthOfLastWord("   fly me   to   the moon  "), 4)
}
