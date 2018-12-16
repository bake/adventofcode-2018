package game

type Grid [][]byte

// Walkable implements the pathfinding.Walker interface.
func (g Grid) Walkable(x, y int) bool {
	if 0 > y || y >= len(g) {
		return false
	}
	if 0 > x || x >= len(g[0]) {
		return false
	}
	return g[y][x] == '.'
}
