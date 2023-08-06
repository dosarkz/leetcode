package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyPow(t *testing.T) {
	assert.Equal(t, MyPow(2.00000, 10), 1024.00000)
	assert.Equal(t, MyPow(2.10000, 3), 9.26100)
	assert.Equal(t, MyPow(2.00000, -2), 0.25000)
	assert.Equal(t, MyPow(8.88023, 3), 700.28148)
	assert.Equal(t, MyPow(1.00000, -2147483648), 1.00000)
}
