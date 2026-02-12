package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	type args struct {
		s string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{"the sky is blue"}, "blue is sky the"},
		{"test2", args{" hello world "}, "world hello"},
		{"test3", args{"a good   example"}, "example good a"},
		{"test4", args{"EPY2giL"}, "EPY2giL"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseWords(tt.args.s); got != tt.want {
				t.Errorf("reverseWords() = %v, want %v", got, tt.want)
			}

			assert.Equalf(t, tt.want, reverseWords(tt.args.s), "reverseWords(%v)", tt.args.s)
		})
	}
}
