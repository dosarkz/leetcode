package _026

import (
	"sort"
)

func uniqueOccurrences(arr []int) bool {
	count := make(map[int]int)
	m := make(map[int]bool)
	sort.Ints(arr)

	for i := 0; i < len(arr); i++ {
		count[arr[i]] = count[arr[i]] + 1
	}

	for _, v := range count {
		if m[v] {
			return false
		}
		m[v] = true
	}

	return true
}
