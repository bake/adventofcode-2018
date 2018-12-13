package mine

import (
	"image"
	"image/color"
)

// Track contains a map of rails and their carts.
type Track struct {
	rails [][]byte
	Carts []*Cart
	crash struct {
		x, y    int
		crashed bool
	}
}

// CartAt returns the cart at a given position.
func (t *Track) CartAt(x, y int) (c *Cart, ok bool) {
	for _, c := range t.Carts {
		if c.X == x && c.Y == y {
			return c, true
		}
	}
	return nil, false
}

// Tick lets all carts move one field. Returns true until two carts crash.
// Check their coordinates with Err().
func (t *Track) Tick() bool {
	turns := map[byte]map[direction]direction{
		'/':  {north: east, south: west, west: south, east: north},
		'\\': {north: west, south: east, west: north, east: south},
	}
	intes := map[turn]map[direction]direction{
		left:     {north: west, south: east, west: south, east: north},
		right:    {north: east, south: west, west: north, east: south},
		straight: {north: north, south: south, west: west, east: east},
	}

	var crash bool
	for _, c := range t.Carts {
		if c.Crashed {
			continue
		}

		cx, cy := c.X, c.Y
		switch c.dir {
		case north:
			cy--
		case south:
			cy++
		case west:
			cx--
		case east:
			cx++
		}
		if d, ok := t.CartAt(cx, cy); ok && !d.Crashed {
			c.Crashed, d.Crashed = true, true
			t.crash.x, t.crash.y, t.crash.crashed = cx, cy, true
			crash = true
		}
		c.X, c.Y = cx, cy

		switch t.rails[c.Y][c.X] {
		case '/', '\\':
			c.dir = turns[t.rails[c.Y][c.X]][c.dir]
		case '+':
			c.dir = intes[c.turn][c.dir]
			c.turn = (c.turn + 1) % 3
		}
	}
	return !crash
}

// Alive returns the number of carts that are not crashed into each other.
func (t *Track) Alive() int {
	var alive int
	for _, c := range t.Carts {
		if !c.Crashed {
			alive++
		}
	}
	return alive
}

// Crash returns coordinates of the last crash and false if no crash has
// happend.
func (t *Track) Crash() (int, int, bool) {
	return t.crash.x, t.crash.y, t.crash.crashed
}

func (t *Track) String() string {
	var str string
	for y := range t.rails {
		for x := range t.rails[y] {
			if c, ok := t.CartAt(x, y); ok && !c.Crashed {
				str += c.String()
				continue
			}
			str += string(t.rails[y][x])
		}
		str += "\n"
	}
	return str
}

func (t *Track) Image() image.Image {
	w, h := len(t.rails[0]), len(t.rails)
	img := image.NewRGBA(image.Rect(0, 0, w+2, h+1))
	for y := range t.rails {
		for x := range t.rails[y] {
			if c, ok := t.CartAt(x, y); ok && !c.Crashed {
				img.Set(x, y, color.RGBA{R: 255, A: 255})
				continue
			}
			if t.rails[y][x] != ' ' {
				img.Set(x, y, color.White)
			}
		}
	}
	return img
}

// Parse parses a slice of slice of bytes into a track. The first slice
// contains rows, the second columns.
func Parse(track [][]byte) *Track {
	dirs := map[byte]direction{
		'^': north, 'v': south,
		'<': west, '>': east,
	}

	var carts []*Cart
	for y, row := range track {
		for x, cell := range row {
			switch cell {
			case '^', 'v', '<', '>':
				carts = append(carts, &Cart{
					X:   x,
					Y:   y,
					dir: dirs[cell],
				})
			}

			// Repair the rail.
			switch cell {
			case '^', 'v':
				track[y][x] = '|'
			case '<', '>':
				track[y][x] = '-'
			}
		}
	}

	return &Track{rails: track, Carts: carts}
}
