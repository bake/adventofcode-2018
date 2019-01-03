package grid

import (
	"bytes"
	"image"
	"image/color"
	"strings"
)

const (
	Open       = '.'
	Trees      = '|'
	Lumberyard = '#'
)

// Grid is a "2D" grid.
type Grid []byte

func (g Grid) at(x, y, w int) (byte, bool) {
	i := y*w + y + x
	if i < 0 || i >= len(g) {
		return ' ', false
	}
	return g[i], true
}

func (g Grid) set(x, y, w int, b byte) bool {
	i := y*w + y + x
	if i < 0 || i >= len(g) {
		return false
	}
	g[i] = b
	return true
}

func (g Grid) adjacent(x, y, w int) []byte {
	var adjs []byte
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dy == 0 && dx == 0 {
				continue
			}
			if b, ok := g.at(x+dx, y+dy, w); ok {
				adjs = append(adjs, b)
			}
		}
	}
	return adjs
}

// Tick simulates one minute in the cellular automata.
func (g Grid) Tick() {
	w := bytes.Index(g, []byte("\n"))
	h := len(g) / w

	new := make(Grid, len(g))
	copy(new, g)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			adjs := g.adjacent(x, y, w)
			switch b, _ := g.at(x, y, w); b {
			case Open:
				if bytes.Count(adjs, []byte{Trees}) >= 3 {
					new.set(x, y, w, Trees)
				}
			case Trees:
				if bytes.Count(adjs, []byte{Lumberyard}) >= 3 {
					new.set(x, y, w, Lumberyard)
				}
			case Lumberyard:
				new.set(x, y, w, Open)
				if bytes.Count(adjs, []byte{Trees}) >= 1 && bytes.Count(adjs, []byte{Lumberyard}) >= 1 {
					new.set(x, y, w, Lumberyard)
				}
			}
		}
	}

	copy(g, new)
}

// String returns a multiline string.
func (g Grid) String() string {
	w := bytes.Index(g, []byte("\n"))
	h := len(g) / w

	var b strings.Builder
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			v, _ := g.at(x, y, w)
			b.WriteByte(v)
		}
		b.WriteByte('\n')
	}
	return b.String()
}

// Image generates an image.
func (g Grid) Image() *image.RGBA {
	w := bytes.Index(g, []byte("\n"))
	h := len(g) / w
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			switch b, _ := g.at(x, y, w); b {
			case Open:
				img.Set(x, y, color.RGBA{B: 255, A: 255})
			case Trees:
				img.Set(x, y, color.RGBA{G: 255, A: 255})
			case Lumberyard:
				img.Set(x, y, color.RGBA{R: 255, A: 255})
			}
		}
	}
	return img
}
