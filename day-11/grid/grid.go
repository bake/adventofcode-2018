package grid

import (
	"image"
	"image/color"
	"math"
)

// Grid holds a grids seiral number.
type Grid int

// Power calculates the power of a grid at a specific postion.
func (g Grid) Power(x, y int) int {
	var power int
	power += x + 10
	power *= y
	power += int(g)
	power *= x + 10
	power = power / 100 % 10
	power -= 5
	return power
}

// Max returns the position and the maximum value of the grid.
func (g Grid) Max(w, h, size int) (x, y, max int) {
	var maxX, maxY int
	for y := 0; y < h-size; y++ {
		for x := 0; x < w-size; x++ {
			var power int
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					power += g.Power(x+j, y+i)
				}
			}
			if power > max {
				max = power
				maxX, maxY = x, y
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
			power := g.Power(x, y)
			img.Set(x, y, color.Gray{math.MaxUint8 / uint8(max) * uint8(power)})
		}
	}
	return img
}
