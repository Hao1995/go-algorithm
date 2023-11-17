package findBall

type BallLoc struct {
	Row int
	Col int
}

func findBall(grid [][]int) []int {
	// init ans of each ball
	var ans []int = make([]int, len(grid[0]))
	for i := 0; i < len(ans); i++ {
		ans[i] = -1
	}

	// for each ball
	for i := 0; i < len(grid[0]); i++ {
		ball := BallLoc{
			Col: i,
			Row: 0,
		}

		// keep downstairs
		var isSet bool
		var outputCol int = -1
		for ball.Row < len(grid) {
			direction := grid[ball.Row][ball.Col]

			// right
			if direction == 1 {
				// stuck by right wall || stuck by the opposite direction
				if ball.Col == len(grid[ball.Row])-1 || grid[ball.Row][ball.Col+1] == -1 {
					ans[i] = -1
					isSet = true
					break // next ball
				}

				ball.Col++
				ball.Row++
				outputCol = ball.Col

				continue
			}

			// left
			if direction == -1 {
				// stuck by left wall || stuck by the opposite direction
				if ball.Col == 0 || grid[ball.Row][ball.Col-1] == 1 {
					ans[i] = -1
					isSet = true
					break // next ball
				}

				ball.Col--
				ball.Row++
				outputCol = ball.Col

				continue
			}
		}

		// set ans
		if !isSet {
			ans[i] = outputCol
		}
	}

	return ans
}
