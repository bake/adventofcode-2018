package distance

// Position describes a point in a 2D space.
type Position struct{ X, Y int }

type item struct{ x, y, dist int }

// All returns a map of all positions and their minimum distance.
func All(bs []byte) map[Position]int {
	var x, y, dist int
	stack := make([]item, len(bs))
	dists := make(map[Position]int)
	for _, b := range bs {
		switch b {
		case '(':
			stack = append(stack, item{x, y, dist})
		case '|':
			head := stack[len(stack)-1]
			x, y, dist = head.x, head.y, head.dist
		case ')':
			head := stack[len(stack)-1]
			x, y, dist = head.x, head.y, head.dist
			stack = stack[:len(stack)-1]
		case 'N', 'S', 'W', 'E':
			switch b {
			case 'N':
				y--
			case 'S':
				y++
			case 'W':
				x--
			case 'E':
				x++
			}
			dist++
			if _, ok := dists[Position{x, y}]; !ok {
				dists[Position{x, y}] = dist
			}
			if dists[Position{x, y}] > dist {
				dists[Position{x, y}] = dist
			}
		}
	}
	return dists
}

// Max returns the length of the longest cycle free path.
func Max(bs []byte) int {
	var max int
	for _, dist := range All(bs) {
		if dist > max {
			max = dist
		}
	}
	return max
}
