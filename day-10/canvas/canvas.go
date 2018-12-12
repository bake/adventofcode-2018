package canvas

import (
	"image"
	"image/color"
	"image/draw"
)

type Point struct{ X, Y, VX, VY int }

type Points []Point

func (pts Points) Image() image.Image {
	b := pts.Bounds()
	img := image.NewRGBA(b)
	draw.Draw(img, b, &image.Uniform{color.Black}, image.ZP, draw.Src)
	for _, p := range pts {
		img.Set(p.X, p.Y, color.RGBA{R: 255, A: 255})
	}
	return img
}

func (pts Points) Bounds() image.Rectangle {
	var r image.Rectangle
	for _, p := range pts {
		if p.X < r.Min.X {
			r.Min.X = p.X
		}
		if p.X > r.Max.X {
			r.Max.X = p.X
		}
		if p.Y < r.Min.Y {
			r.Min.Y = p.Y
		}
		if p.Y > r.Max.Y {
			r.Max.Y = p.Y
		}
	}
	return r
}
