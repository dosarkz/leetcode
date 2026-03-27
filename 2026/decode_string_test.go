package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeString(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "3[a]2[bc]",
			want:  "aaabcbc",
		},
		{
			input: "3[a2[c]]",
			want:  "accaccacc",
		},
		{
			input: "2[abc]3[cd]3[ee]ef",
			want:  "abcabccdcdcdeeeeeeef",
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.want, decodeString(test.input))
	}
}
