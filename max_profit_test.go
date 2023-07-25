package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, MaxProfit([]int{7, 1, 5, 3, 6, 4}), 5, "Test should be equal 5")
	assert.Equal(t, MaxProfit([]int{7, 6, 4, 3, 1}), 0, "Test should be equal 5")

}
