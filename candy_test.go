package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCandy(t *testing.T) {
	//assert.Equal(t, Candy([]int{1, 0, 2}), 5)
	//assert.Equal(t, Candy([]int{1, 2, 2}), 4)
	//assert.Equal(t, Candy([]int{1, 3, 2, 2, 1}), 7)
	assert.Equal(t, Candy([]int{1, 2, 87, 87, 87, 2, 1}), 13)
}
