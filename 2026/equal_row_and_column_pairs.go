package _026

func equalPairs(grid [][]int) int {

	cols := make([][]int, len(grid))
	count := 0

	for i := range cols {
		for j := 0; j < len(grid); j++ {
			cols[i] = append(cols[i], grid[j][i])
		}
	}

	for i := range grid {
		for k := 0; k < len(cols[i]); k++ {
			all := len(grid[k])
			for idx, v := range grid[k] {
				if v != cols[i][idx] {
					all = len(grid[k])
					idx = all
					continue
				}
				all--
				if all == 0 {
					count++
				}
			}
		}
	}

	return count
}
