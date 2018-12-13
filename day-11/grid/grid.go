package grid

import (
	"image"
	"image/color"
	"math"
)

// Grid holds a powers.
type Grid [][]int

// New generates a new grid.
func New(w, h, serial int) Grid {
	g := make(Grid, h)
	for i := range g {
		g[i] = make([]int, w)
		for j := range g[i] {
			g[i][j] = power(j, i, serial)
		}
	}
	return g
}

// power calculates the power of a grid at a specific position.
func power(x, y, serial int) int {
	var power int
	power += x + 10
	power *= y
	power += serial
	power *= x + 10
	power = power / 100 % 10
	power -= 5
	return power
}

// Max returns the position and the maximum value of the grid.
func (g Grid) Max(size int) (x, y, max int) {
	var maxX, maxY int
	for y := 0; y < len(g)-size; y++ {
		for x := 0; x < len(g[y])-size; x++ {
			var power int
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					power += g[y+i][x+j]
				}
			}
			if power > max {
				max, maxX, maxY = power, x, y
			}
		}
	}
	return maxX, maxY, max
}

// Image generates an image with grayscale values from 0 to max.
func (g Grid) Image(w, h, max int) image.Image {
	img := image.NewGray(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			col := color.Gray{math.MaxUint8 / uint8(max) * uint8(g[y][x])}
			img.Set(x, y, col)
		}
	}
	return img
}
