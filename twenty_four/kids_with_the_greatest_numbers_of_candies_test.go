package twenty_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_kidsWithCandies(t *testing.T) {
	type args struct {
		candies      []int
		extraCandies int
	}

	tests := []struct {
		name string
		args args
		want []bool
	}{
		{
			name: "first iteration",
			args: args{
				candies:      []int{2, 3, 5, 1, 3},
				extraCandies: 3,
			},
			want: []bool{true, true, true, false, true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kidsWithCandies(tt.args.candies, tt.args.extraCandies), "kidsWithCandies(%v, %v)", tt.args.candies, tt.args.extraCandies)
		})
	}
}
