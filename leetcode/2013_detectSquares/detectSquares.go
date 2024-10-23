package detectsquares

type DetectSquares struct {
	// [x][y]=number of points
	Points map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		Points: make(map[int]map[int]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	yPoints, ok := this.Points[point[0]]
	if !ok {
		yPoints = make(map[int]int)
	}
	yPoints[point[1]]++
	this.Points[point[0]] = yPoints
}

func (this *DetectSquares) Count(point []int) int {
	var ans int
	for x, ypoints := range this.Points {
		for y, count := range ypoints {
			// check if it's a diagonal point
			if abs(x-point[0]) != abs(y-point[1]) || x == point[0] || y == point[1] {
				continue
			}

			// count possible squares
			xCount, _ := this.Points[x][point[1]]
			yCount, _ := this.Points[point[0]][y]

			ans += count * xCount * yCount
		}
	}
	return ans
}

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
