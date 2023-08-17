package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertZigzag(t *testing.T) {
	assert.Equal(t, Convert("PAYPALISHIRING", 3), "PAHNAPLSIIGYIR")
	assert.Equal(t, Convert("PAYPALISHIRING", 4), "PINALSIGYAHRPI")
	assert.Equal(t, Convert("ABC", 2), "ACB")
	assert.Equal(t, Convert("ABCD", 2), "ACBD")

}

// a  c
// b  d
