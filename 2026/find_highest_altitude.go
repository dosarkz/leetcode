package _026

func largestAltitude(gain []int) int {
	altitude := 0
	maxAltitude := 0

	for _, g := range gain {
		altitude += g
		maxAltitude = max(maxAltitude, altitude)
	}
	return maxAltitude
}
