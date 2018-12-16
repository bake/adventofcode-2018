package pathfinding

type SearchFunc func(w Walker, dst, src XYer) (path XYs, ok bool)

// BreadthFirstSearch returns the shortest path from an XYer to another.
func BreadthFirstSearch(w Walker, dst, src XYer) (path XYs, ok bool) {
	queue := XYs{{dst.X(), dst.Y()}}
	visited := map[XY]bool{}
	previous := map[XY]XYs{}
	paths := XYs{}
	found := false
SearchLoop:
	for len(queue) > 0 {
		for _, xy := range queue {
			queue = queue[1:]
			for _, n := range neighbours(w, xy, src) {
				previous[n] = append(previous[n], xy)
				if visited[n] {
					continue
				}
				queue = append(queue, n)
				visited[n] = true
				if n.X() == src.X() && n.Y() == src.Y() {
					found = true
					paths = append(paths, XY{n.X(), n.Y()})
					break SearchLoop
				}
			}
		}
	}
	if !found {
		return nil, false
	}

	path = XYs{{src.X(), src.Y()}}
	for {
		curr := path[len(path)-1]
		if curr.x == dst.X() && curr.y == dst.Y() {
			break
		}
		// sort.Sort(previous[curr])
		path = append(path, previous[curr][0])
	}
	return path, true
}

func neighbours(w Walker, xy XY, src XYer) XYs {
	// Deltas in reading order.
	ds := XYs{{0, -1}, {-1, 0}, {+1, 0}, {0, +1}}
	xys := XYs{}
	for _, d := range ds {
		xy := XY{xy.x + d.x, xy.y + d.y}
		if src.X() == xy.x && src.Y() == xy.y {
			xys = append(xys, xy)
			continue
		}
		if w.Walkable(xy.x, xy.y) {
			xys = append(xys, xy)
			continue
		}
	}
	return xys
}
