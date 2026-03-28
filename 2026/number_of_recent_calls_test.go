package _026

import "testing"

func TestNumberOfRecentCalls(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		wants []int
	}{
		{
			name:  "test1",
			input: []int{1, 100, 3001, 3002},
			wants: []int{1, 2, 3, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := Constructor()
			for _, v := range tt.input {
				if got := obj.Ping(v); got != tt.wants[0] {
					t.Errorf("RecentCounter.Ping() = %v, want %v", got, tt.wants[0])
				}
				tt.wants = tt.wants[1:]
			}
		})
	}
}
