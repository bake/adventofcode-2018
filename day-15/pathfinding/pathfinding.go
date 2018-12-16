package pathfinding

type XYer interface {
	X() int
	Y() int
}

type Walker interface {
	Walkable(x, y int) bool
}

type XY struct{ x, y int }

func (xy XY) X() int          { return xy.x }
func (xy XY) Y() int          { return xy.y }
func (xy XY) Equal(b XY) bool { return xy.x == b.x && xy.y == b.y }

type XYs []XY

func (xys XYs) At(x, y int) (xy XY, ok bool) {
	for _, xy := range xys {
		if xy.x == x && xy.y == y {
			return xy, true
		}
	}
	return XY{}, false
}

func (xys XYs) Len() int { return len(xys) }
func (xys XYs) Less(i, j int) bool {
	return xys[i].y < xys[j].y || (xys[i].y == xys[j].y && xys[i].x < xys[j].x)
}
func (xys XYs) Swap(i, j int) { xys[i], xys[j] = xys[j], xys[i] }
