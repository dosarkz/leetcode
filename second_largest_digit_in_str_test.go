package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSecondHighest(t *testing.T) {
	//assert.Equal(t, SecondHighest("dfa12321afd"), 2, "test should be equal 2")
	//assert.Equal(t, SecondHighest("abc1111"), -1, "test should be equal -1")
	assert.Equal(t, SecondHighest("ck077"), 0, "test should be equal -1")
}
