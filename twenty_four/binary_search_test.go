package twenty_four

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}

	type test struct {
		name string
		args args
		want int
	}

	var tests = []test{
		{
			name: "first iteration",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			want: 4,
		},

		{
			name: "sec iteration",
			args: args{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, search(tt.args.nums, tt.args.target), "search(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
