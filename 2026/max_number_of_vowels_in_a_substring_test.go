package _026

import "testing"

func TestMaxVowels(t *testing.T) {
	tests := []struct {
		s        string
		k        int
		expected int
	}{
		{"abciiidef", 3, 3},
		{"aeiou", 2, 2},
		{"leetcode", 3, 2},
	}

	for _, test := range tests {
		result := maxVowels(test.s, test.k)
		if result != test.expected {
			t.Errorf("For input (%s, %d), expected %d but got %d", test.s, test.k, test.expected, result)
		}
	}
}
