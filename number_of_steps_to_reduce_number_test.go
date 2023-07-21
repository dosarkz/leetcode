package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfSteps(t *testing.T) {
	assert.Equal(t, NumberOfSteps(14), 6, "Tests should be equal")
	assert.Equal(t, NumberOfSteps(8), 4, "Tests should be equal")
	assert.Equal(t, NumberOfSteps(123), 12, "Tests should be equal")
}
