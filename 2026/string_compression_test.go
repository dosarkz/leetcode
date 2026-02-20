package _026

import "testing"

func TestStringCompression(t *testing.T) {
	tests := []struct {
		name  string
		chars []byte
		want  int
	}{
		{"test1", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, 6},
		{"test2", []byte{'a'}, 1},
		{"test3", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, 4},
	}
	for _, tt := range tests {
		if got := StringCompression(tt.chars); got != tt.want {
			t.Errorf("%s: StringCompression(%v) = %v, want %v", tt.name, tt.chars, got, tt.want)
		}
	}
}
