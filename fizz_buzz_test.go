package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, FizzBuzz(3), []string{"1", "2", "Fizz"})
}
