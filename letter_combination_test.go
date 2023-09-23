package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	//assert.Equal(t, LetterCombinations("23"), []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"})

	assert.Equal(t, LetterCombinations("234"), []string{"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi", "bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi", "cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi"})
	//assert.Equal(t, LetterCombinations("2"), []string{"a", "b", "c"})
	//assert.Equal(t, LetterCombinations(""), []string{})

}
