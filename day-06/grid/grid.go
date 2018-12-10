package grid

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

// Grid holds XYs.
type Grid struct {
	points map[XY]int // points are the base coordinates.
	locs   map[XY]int // locs are all fields.
}

// New generates a new grid.
func New() Grid {
	return Grid{map[XY]int{}, map[XY]int{}}
}

// At returns the ID at a given point.
func (g Grid) At(p XY) int {
	return g.locs[p]
}

// SetPoint sets a point and a location, regardless of its prevous values.
func (g Grid) SetPoint(p XY, id int) {
	g.points[p] = id
	g.locs[p] = id
}

// Set saves a points id to a location. If the position is already occupied, it
// sets it to fieldEqualDist. This is only useful, if the grid contains only
// locations that are created within the same iteration.
func (g Grid) Set(p XY, id int) {
	if g.locs[p] == fieldEmpty {
		g.locs[p] = id
	}
	if g.locs[p] != id {
		g.locs[p] = fieldEqualDist
	}
}

// Expand expends the points by a given radius using manhattan distance. This
// way there is no need to iterate over all w*h pixels on every iteration. It
// creates a new grid that contains only the new (expanded) fields.
func (g Grid) Expand(radius int) {
	h := New()
	for p, id := range g.points {
		for x := 0; x < radius; x++ {
			y := radius - x - 1
			h.Set(XY{p.X() - x, p.Y() - y}, id)
			h.Set(XY{p.X() - x, p.Y() + y}, id)
			h.Set(XY{p.X() + x, p.Y() - y}, id)
			h.Set(XY{p.X() + x, p.Y() + y}, id)
		}
	}
	for p, id := range h.locs {
		if g.locs[p] == fieldEmpty {
			g.locs[p] = id
		}
	}
}

// Print prints the grid to the console. use image() for larger systems.
func (g Grid) Print(b image.Rectangle) {
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			id := g.locs[XY{x, y}]
			switch id {
			case fieldEmpty:
				fmt.Print("  ")
			case fieldEqualDist:
				fmt.Print("..")
			default:
				i := 97
				if g.points[XY{x, y}] == id {
					i = 65
				}
				fmt.Printf("%c ", i+id-1)
			}
		}
		fmt.Println()
	}
}

// Image generates an image drawing the grid.
func (g Grid) Image(b image.Rectangle) image.Image {
	img := image.NewRGBA(b)
	for p, id := range g.locs {
		switch id {
		case fieldEqualDist:
			img.Set(p.X(), p.Y(), color.RGBA{0, 0, 0, 255})
		default:
			img.Set(p.X(), p.Y(), palette.WebSafe[id%len(palette.WebSafe)])
		}
	}
	for p, id := range g.points {
		g.label(img, p.X(), p.Y(), strconv.Itoa(id))
	}
	return img
}

// label draws text on an image.
func (g Grid) label(img *image.RGBA, x, y int, label string) {
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

// Max calculates the size of the largest area.
func (g Grid) Max(edgePoints map[int]bool) int {
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
