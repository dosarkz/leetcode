package twenty_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGcdOfStrings(t *testing.T) {
	assert.Equal(t, gcdOfStrings("ABCABC", "ABC"), "ABC")
	assert.Equal(t, gcdOfStrings("ABABAB", "ABAB"), "AB")
	assert.Equal(t, gcdOfStrings("ABCDEF", "ABC"), "")

	assert.Equal(t, gcdOfStrings("LEET", "CODE"), "")
}
