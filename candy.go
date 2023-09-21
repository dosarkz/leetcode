package main

func Candy(ratings []int) int {
	n := len(ratings)

	candies := make([]int, n)
	for i := range candies {
		candies[i] = 1
	}

	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		}
	}

	for j := n - 2; j >= 0; j-- {
		if ratings[j] > ratings[j+1] {
			if candies[j] <= candies[j+1] {
				candies[j] = candies[j+1] + 1
			}
		}
	}

	totalCandies := 0
	for _, candy := range candies {
		totalCandies += candy
	}

	return totalCandies
}
