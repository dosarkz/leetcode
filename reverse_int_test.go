package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseInt(t *testing.T) {
	assert.Equal(t, ReverseInt(123), 321)
	assert.Equal(t, ReverseInt(-123), -321)
	assert.Equal(t, ReverseInt(1534236469), 0)
}
