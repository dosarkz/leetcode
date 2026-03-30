package _026

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDota2Senate(t *testing.T) {
	tests := []struct {
		senate string
		want   string
	}{
		{senate: "RD", want: "Radiant"},
		{senate: "RDD", want: "Dire"},
		{senate: "RRDDD", want: "Radiant"},
	}
	for _, tt := range tests {
		t.Run(tt.senate, func(t *testing.T) {
			assert.Equal(t, tt.want, predictPartyVictory(tt.senate))
		})
	}
}
