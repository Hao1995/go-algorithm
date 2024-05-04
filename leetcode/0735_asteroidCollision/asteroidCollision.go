package asteroidcollision

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		if asteroid < 0 && len(stack) > 0 {
			for len(stack) > 0 {
				var prevAsteroid int
				prevAsteroid, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if prevAsteroid < 0 {
					stack = append(stack, prevAsteroid)
					stack = append(stack, asteroid)
					break
				}

				if asteroid*-1 == prevAsteroid {
					// both were collided
					break
				} else if asteroid*-1 > prevAsteroid {
					// keep colliding
					if len(stack) == 0 {
						stack = append(stack, asteroid)
						break
					}
					continue
				} else {
					// curr was collided
					stack = append(stack, prevAsteroid)
					break
				}
			}
			continue
		}
		stack = append(stack, asteroid)
	}
	return stack
}
