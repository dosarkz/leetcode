package main

func SecondHighest(s string) int {
	mx := -1
	mx2 := -1
	for i := range s {
		next := s[i]
		if next >= '0' && next <= '9' {
			value := int(next - '0')
			if value > mx {
				mx2 = mx
				mx = value
				continue
			}
			if value != mx && value > mx2 {
				mx2 = value
			}
		}
	}
	return mx2
}
