package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	//assert.Equal(t, MyAtoi(" 42"), 42)
	//assert.Equal(t, MyAtoi("   -42"), -42)
	//assert.Equal(t, MyAtoi("+1"), 1)
	//assert.Equal(t, MyAtoi("9223372036854775808"), 2147483647)
	//assert.Equal(t, MyAtoi("-91283472332"), -2147483648)
	//assert.Equal(t, MyAtoi("  -0012a42"), -12)
	assert.Equal(t, MyAtoi("-2147483648"), -2147483648)
}
