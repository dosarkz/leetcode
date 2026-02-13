package _026

import "testing"

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"test1", "hello", "holle"},
		{"test2", "leetcode", "leotcede"},
		{"test3", "aA", "Aa"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseVowels(tt.s); got != tt.want {
				t.Errorf("reverseVowels() = %v, want %v", got, tt.want)
			}
		},
		)
	}
}
