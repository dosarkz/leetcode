package twenty_five

func CanPlaceFlowers(flowerbed []int, n int) bool {
	places := 0
	if n == 0 {
		return true
	}

	if len(flowerbed) == 1 {
		if flowerbed[0] == 0 {
			return true
		}
		return false
	}

	// find available places on array
	for i := 0; i < len(flowerbed); i++ {
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				places++
				i++ // skip next place
				continue
			}
		}

		// first places and second places are empty
		if flowerbed[i] == 0 && len(flowerbed)-1 != i && flowerbed[i+1] == 0 && flowerbed[i-1] == 0 {
			places++
			i++ // skip next place
			if places >= n {
				return true
			}
			continue
		}

		// last place
		if i == len(flowerbed)-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				places++
			}
		}

		if places >= n {
			return true
		}
	}

	return places >= n
}
