package twenty_five

import "testing"

func TestCanPlaceFlowers(t *testing.T) {
	tests := []struct {
		flowerbed []int
		n         int
		want      bool
	}{
		{
			flowerbed: []int{1},
			n:         0,
			want:      true,
		},
		//{
		//	flowerbed: []int{1, 0, 0, 0, 1},
		//	n:         1,
		//	want:      true,
		//},
		//{
		//	flowerbed: []int{1, 0, 0, 0, 1},
		//	n:         2,
		//	want:      false,
		//},
		//{
		//	flowerbed: []int{0, 0, 1, 0, 0},
		//	n:         1,
		//	want:      true,
		//},
		//{
		//	flowerbed: []int{0, 0, 0, 0, 0},
		//	n:         3,
		//	want:      true,
		//},
		//{
		//	flowerbed: []int{1, 0, 1, 0, 1, 0, 1},
		//	n:         1,
		//	want:      false,
		//},
		//{
		//	flowerbed: []int{0},
		//	n:         1,
		//	want:      true,
		//},
		//{
		//	flowerbed: []int{1},
		//	n:         1,
		//	want:      false,
		//},
		//{
		//	flowerbed: []int{1, 0, 0, 0, 0, 1},
		//	n:         2,
		//	want:      false,
		//},
	}

	for _, tt := range tests {
		got := CanPlaceFlowers(tt.flowerbed, tt.n)
		if got != tt.want {
			t.Errorf("CanPlaceFlowers(%v, %d) = %v, want %v", tt.flowerbed, tt.n, got, tt.want)
		}

	}

}
