package _026

import "testing"

func TestFindHighestAltitude(t *testing.T) {
	tests := []struct {
		name string
		gain []int
		want int
	}{
		{
			name: "test1",
			gain: []int{-5, 1, 5, 0, -7},
			want: 1,
		},
		{
			name: "test2",
			gain: []int{-4, -3, -2, -1, 4, 3, 2},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestAltitude(tt.gain); got != tt.want {
				t.Errorf("largestAltitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
