package twenty_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "first iteration",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: 3,
		},

		{
			name: "second iteration",
			args: args{
				nums: []int{2, 2, 1, 1, 1, 2, 2},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, majorityElement(tt.args.nums), "majorityElement(%v)", tt.args.nums)
		})
	}
}
