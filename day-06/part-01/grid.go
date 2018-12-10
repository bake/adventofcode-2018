package main

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	// the fields default value describes an empty field, therefore ids have to
	// start at 1.
	fieldEmpty = 0
	// equally far from >= 2 points, ignore.
	fieldEqualDist = -1
)

type grid struct {
	points map[xy]int // points are the base coordinates.
	locs   map[xy]int // locs are all fields.
}

func newGrid() grid {
	return grid{map[xy]int{}, map[xy]int{}}
}

func (g grid) at(p xy) int {
	return g.locs[p]
}

// setPoint sets a point and a location, regardless of its prevous values.
func (g grid) setPoint(p xy, id int) {
	g.points[p] = id
	g.locs[p] = id
}

// set saves a points id to a location. If the position is already occupied, it
// sets it to fieldEqualDist. This is only useful, if the grid contains only
// locations that are created within the same iteration.
func (g grid) set(p xy, id int) {
	if g.locs[p] == fieldEmpty {
		g.locs[p] = id
	}
	if g.locs[p] != id {
		g.locs[p] = fieldEqualDist
	}
}

// expend expends the points by a given radius using manhattan distance. This
// way there is no need to iterate over all w*h pixels on every iteration. It
// creates a new grid that contains only the new (expanded) fields.
func (g grid) expand(radius int) {
	h := newGrid()
	for p, id := range g.points {
		for x := 0; x < radius; x++ {
			y := radius - x - 1
			h.set(xy{p.x() - x, p.y() - y}, id)
			h.set(xy{p.x() - x, p.y() + y}, id)
			h.set(xy{p.x() + x, p.y() - y}, id)
			h.set(xy{p.x() + x, p.y() + y}, id)
		}
	}
	for p, id := range h.locs {
		if g.locs[p] == fieldEmpty {
			g.locs[p] = id
		}
	}
}

// print prints the grid to the console. use image() for larger systems.
func (g grid) print(b image.Rectangle) {
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			id := g.locs[xy{x, y}]
			switch id {
			case fieldEmpty:
				fmt.Print("  ")
			case fieldEqualDist:
				fmt.Print("..")
			default:
				i := 97
				if g.points[xy{x, y}] == id {
					i = 65
				}
				fmt.Printf("%c ", i+id-1)
			}
		}
		fmt.Println()
	}
}

// image generates an image drawing the grid.
func (g grid) image(b image.Rectangle) image.Image {
	img := image.NewRGBA(b)
	for p, id := range g.locs {
		switch id {
		case fieldEqualDist:
			img.Set(p.x(), p.y(), color.RGBA{0, 0, 0, 255})
		default:
			img.Set(p.x(), p.y(), palette.WebSafe[id%len(palette.WebSafe)])
		}
	}
	for p, id := range g.points {
		g.label(img, p.x(), p.y(), strconv.Itoa(id))
	}
	return img
}

// label draws a label on an image.
func (g grid) label(img *image.RGBA, x, y int, label string) {
	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(color.RGBA{0, 0, 0, 255}),
		Face: basicfont.Face7x13,
		Dot: fixed.Point26_6{
			X: fixed.Int26_6(x * 64),
			Y: fixed.Int26_6(y * 64),
		},
	}
	d.DrawString(label)
}

func (g grid) copy() grid {
	h := newGrid()
	for p, id := range g.points {
		h.points[p] = id
	}
	for p, id := range g.locs {
		h.locs[p] = id
	}
	return h
}

// max calculates the size of the largest area.
func (g grid) max(edgePoints map[int]bool) int {
	areas := map[int]int{}
	for _, id := range g.locs {
		if !edgePoints[id] {
			areas[id]++
		}
	}
	var max int
	for _, n := range areas {
		if n > max {
			max = n
		}
	}
	return max
}
