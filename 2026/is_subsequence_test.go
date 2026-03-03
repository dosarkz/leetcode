package _026

import "testing"

func TestIsSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"test1", args{"abc", "ahbgdc"}, true},
		{"test2", args{"axc", "ahbgdc"}, false},
	}
	for _, tt := range tests {
		if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
			t.Errorf("%q. isSubsequence() = %v, want %v", tt.name, got, tt.want)
		}
	}
}
