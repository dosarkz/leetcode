package _026

import "testing"

func TestDetemineOfTwoStrClose(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  bool
	}{
		//{"abc", "bca", true},
		//{"a", "aa", false},
		{"cabbba", "abbccc", true},
		{"uau", "ssx", false},
		{"aaabbbbccddeeeeefffff", "aaaaabbcccdddeeeeffff", false},
	}

	for _, tt := range tests {
		if got := closeStrings(tt.word1, tt.word2); got != tt.want {
			t.Errorf("closeStrings(%v, %v) = %v, want %v", tt.word1, tt.word2, got, tt.want)
		}
	}
}
