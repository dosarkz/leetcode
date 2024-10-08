package twenty_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_threeSumClosest(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first iteration",
			args: args{
				nums:   []int{-1, 2, 1, -4},
				target: 1,
			},
			want: 2,
		},

		{
			name: "first iteration",
			args: args{
				nums:   []int{0, 0, 0},
				target: 1,
			},
			want: 0,
		},

		{
			name: "first iteration",
			args: args{
				nums:   []int{0, 1, 2},
				target: 3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, threeSumClosest(tt.args.nums, tt.args.target), "threeSumClosest(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
