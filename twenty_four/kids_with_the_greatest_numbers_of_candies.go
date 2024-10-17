package twenty_four

func kidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	m := getMaxValueInArr(candies)

	for i := 0; i < len(candies); i++ {
		result[i] = candies[i]+extraCandies >= m
	}

	return result
}

func getMaxValueInArr(items []int) int {
	m := 0

	for i := 0; i < len(items); i++ {
		if items[i] > m {
			m = items[i]
		}
	}

	return m
}
