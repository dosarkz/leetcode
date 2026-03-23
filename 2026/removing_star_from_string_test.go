package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemovingStarFromString(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{input: "leet**cod*e", output: "lecoe"},
	}
	for _, test := range tests {
		assert.Equal(t, test.output, removeStar(test.input))
	}
}
