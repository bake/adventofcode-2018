package canvas

import (
	"image"
	"image/color"
	"image/draw"
	"math"
)

// Point holds coordinates an velocity.
type Point struct{ X, Y, VX, VY int }

// Points is a slice of points.
type Points []Point

// Image plots points to an image.
func (pts Points) Image() image.Image {
	b := pts.Bounds()
	img := image.NewRGBA(b)
	draw.Draw(img, b, &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, p := range pts {
		img.Set(p.X, p.Y, color.RGBA{R: 255, A: 255})
	}
	return img
}

// Bounds returns a recangle containing all points.
func (pts Points) Bounds() image.Rectangle {
	minX, maxX := math.MaxInt32, math.MinInt32
	minY, maxY := math.MaxInt32, math.MinInt32
	for _, p := range pts {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	return image.Rect(minX, minY, maxX+1, maxY+1)
}
