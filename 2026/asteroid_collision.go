package _026

func asteroidCollision(asteroids []int) []int {
	var stack []int

	for _, a := range asteroids {
		if a > 0 {
			stack = append(stack, a)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && stack[len(stack)-1] < -a {
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] == -a {
				stack = stack[:len(stack)-1]
			} else if len(stack) == 0 || stack[len(stack)-1] < -a {
				stack = append(stack, a)
			}
		}
	}

	return stack
}
