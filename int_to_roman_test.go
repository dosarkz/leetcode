package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	assert.Equal(t, IntToRoman(1994), "MCMXCIV")
	//assert.Equal(t, IntToRoman(3), "III")
	//assert.Equal(t, IntToRoman(58), "LVIII")
}
